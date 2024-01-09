package tidb

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tidb-parser"
	tidbast "github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
)

type StringsManipulator struct {
	s      string
	l      *parser.TiDBLexer
	stream antlr.TokenStream
	p      *parser.TiDBParser
	tree   antlr.Tree
}

func NewStringsManipulator(s string) *StringsManipulator {
	l := parser.NewTiDBLexer(antlr.NewInputStream(""))
	stream := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)
	return &StringsManipulator{
		s:      s,
		l:      l,
		stream: stream,
		p:      parser.NewTiDBParser(stream),
	}
}

type StringsManipulatorActionType int

const (
	StringsManipulatorActionTypeNone StringsManipulatorActionType = iota
	StringsManipulatorActionTypeDropTable
	StringsManipulatorActionTypeDropColumn
	StringsManipulatorActionTypeModifyColumnType
	StringsManipulatorActionTypeDropColumnOption
	StringsManipulatorActionTypeModifyColumnOption
	StringsManipulatorActionTypeDropTableConstraint
	StringsManipulatorActionTypeModifyTableConstraint
)

type StringsManipulatorAction interface {
	getType() StringsManipulatorActionType
	getTopLevelNaming() string
	getSecondLevelNaming() string
}

type StringsManipulatorActionBase struct {
	Type StringsManipulatorActionType
}

func (s *StringsManipulatorActionBase) getType() StringsManipulatorActionType {
	return s.Type
}

type StringsManipulatorActionDropTable struct {
	StringsManipulatorActionBase
	Table string
}

func (s *StringsManipulatorActionDropTable) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionDropTable) getSecondLevelNaming() string {
	return ""
}

type StringsManipulatorActionDropColumn struct {
	StringsManipulatorActionBase
	Table  string
	Column string
}

func (s *StringsManipulatorActionDropColumn) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionDropColumn) getSecondLevelNaming() string {
	return s.Column
}

type StringsManipulatorActionModifyColumnType struct {
	StringsManipulatorActionBase
	Table  string
	Column string
	Type   string
}

func (s *StringsManipulatorActionModifyColumnType) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionModifyColumnType) getSecondLevelNaming() string {
	return s.Column
}

type StringsManipulatorActionDropColumnOption struct {
	StringsManipulatorActionBase
	Table  string
	Column string
	Option tidbast.ColumnOptionType
}

func (s *StringsManipulatorActionDropColumnOption) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionDropColumnOption) getSecondLevelNaming() string {
	return s.Column
}

type StringsManipulatorActionModifyColumnOption struct {
	StringsManipulatorActionBase
	Table           string
	Column          string
	OldOption       tidbast.ColumnOptionType
	NewOptionDefine string
}

func (s *StringsManipulatorActionModifyColumnOption) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionModifyColumnOption) getSecondLevelNaming() string {
	return s.Column
}

type StringsManipulatorActionDropTableConstraint struct {
	StringsManipulatorActionBase
	Table          string
	Constraint     tidbast.ConstraintType
	ConstraintName string
}

func (s *StringsManipulatorActionDropTableConstraint) getTopLevelNaming() string {
	return s.Table
}

func (s *StringsManipulatorActionDropTableConstraint) getSecondLevelNaming() string {
	return s.ConstraintName
}

type StringsManipulatorActionModifyTableConstraint struct {
	StringsManipulatorActionBase
	Table               string
	OldConstraint       tidbast.ConstraintType
	OldConstraintName   string
	NewConstraintDefine string
}

func (s *StringsManipulatorActionModifyTableConstraint) getTopLevelNaming() string {
	return s.Table
}

var (
	regexpColumn = regexp.MustCompile("^  `([^`]+)`")
)

func (s *StringsManipulator) Manipulate(actions ...StringsManipulatorAction) (string, error) {
	tableActions := make(map[string][]StringsManipulatorAction)

	for _, action := range actions {
		tableName := action.getTopLevelNaming()
		// do copy
		action := action
		tableActions[tableName] = append(tableActions[tableName], action)
	}

	stmts, err := SplitSQL(s.s)
	if err != nil {
		return "", errors.Wrap(err, "failed to split sql")
	}

	var results []string

	for _, stmt := range stmts {
		if stmt.Empty {
			results = append(results, stmt.Text)
			continue
		}
		isCreateTable, tableName := extractTableNameForCreateTable(stmt.Text)
		if !isCreateTable {
			results = append(results, stmt.Text)
			continue
		}
		actions, ok := tableActions[tableName]
		if !ok || len(actions) == 0 {
			results = append(results, stmt.Text)
			continue
		}

		var tableDefinition strings.Builder
		var tableActions []StringsManipulatorAction
		actionsMap := make(map[string][]StringsManipulatorAction)
		for _, action := range actions {
			// do copy
			action := action
			secondName := action.getSecondLevelNaming()
			if secondName == "" {
				tableActions = append(tableActions, action)
			} else {
				actionsMap[secondName] = append(actionsMap[secondName], action)
			}
		}

		hasDropTable := false
		for _, action := range tableActions {
			if _, ok := action.(*StringsManipulatorActionDropTable); ok {
				hasDropTable = true
				continue
			}
		}
		if hasDropTable {
			continue
		}

		scanner := bufio.NewScanner(strings.NewReader(stmt.Text))
		for scanner.Scan() {
			line := scanner.Text()

			columnMatch := regexpColumn.FindStringSubmatch(line)
			if len(columnMatch) > 1 {
				// is column definition
				columnName := columnMatch[1]
				actions, ok := actionsMap[columnName]
				if !ok || len(actions) == 0 {
					if _, err := tableDefinition.WriteString(line); err != nil {
						return "", errors.Wrap(err, "failed to write string")
					}
					continue
				}

				s.Load(line)
				if err := s.ParseColumnDef(); err != nil {
					return "", errors.Wrapf(err, "failed to parse column def: %s", line)
				}

				columnActionMap := make(map[StringsManipulatorActionType][]StringsManipulatorAction)
				for _, action := range actions {
					columnActionMap[action.getType()] = append(columnActionMap[action.getType()], action)
				}
				result, err := s.RewriteColumnDef(columnActionMap)
				if err != nil {
					return "", errors.Wrapf(err, "failed to rewrite column def: %s", line)
				}
				if len(result) > 0 {
					results = append(results, result)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return "", errors.Wrap(err, "failed to scan create table statement")
		}
	}

	return strings.Join(results, "\n"), nil
}

func (s *StringsManipulator) RewriteColumnDef(columnActionMap map[StringsManipulatorActionType][]StringsManipulatorAction) (string, error) {
	if columnActionMap == nil || len(columnActionMap) == 0 {
		return s.stream.GetAllText(), nil
	}
	if dropTable, exists := columnActionMap[StringsManipulatorActionTypeDropColumn]; exists && len(dropTable) > 0 {
		return "", nil
	}
	listener := &rewriter{
		rewriter: antlr.NewTokenStreamRewriter(s.stream),
		actions:  columnActionMap,
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, s.tree)
	if listener.err != nil {
		return "", errors.Wrap(listener.err, "failed to rewrite column def for column")
	}
	return listener.rewriter.GetTextDefault(), nil
}

type rewriter struct {
	*antlr.BaseParseTreeListener

	rewriter *antlr.TokenStreamRewriter
	actions  map[StringsManipulatorActionType][]StringsManipulatorAction
	err      error
}

func (r *rewriter) EnterColumnDef(ctx *parser.ColumnDefContext) {
	if modifyType, exists := r.actions[StringsManipulatorActionTypeModifyColumnType]; exists && len(modifyType) > 0 {
		if len(modifyType) > 1 {
			r.err = errors.New("multiple modify column type actions")
			return
		}
		modifyType := (modifyType[0]).(*StringsManipulatorActionModifyColumnType)
		r.rewriter.ReplaceTokenDefault(ctx.DataType().GetStart(), ctx.DataType().GetStop(), modifyType.Type)
	}

	if modifyColumnOption, exists := r.actions[StringsManipulatorActionTypeModifyColumnOption]; exists && len(modifyColumnOption) > 0 {
		modifyOptionMap := make(map[tidbast.ColumnOptionType]StringsManipulatorAction)
		for _, action := range modifyColumnOption {
			action := action.(*StringsManipulatorActionModifyColumnOption)
			modifyOptionMap[action.OldOption] = action
		}
		dropOptionMap := make(map[tidbast.ColumnOptionType]StringsManipulatorAction)
		for _, action := range r.actions[StringsManipulatorActionTypeDropColumnOption] {
			action := action.(*StringsManipulatorActionDropColumnOption)
			dropOptionMap[action.Option] = action
		}
		if ctx.ColumnOptionList() != nil {
			for _, option := range ctx.ColumnOptionList().AllColumnOption() {
				optionType := convertColumnOptionType(option)
				if _, exists := dropOptionMap[optionType]; exists {
					r.rewriter.DeleteTokenDefault(option.GetStart(), option.GetStop())
					continue
				}
			}
		}
	}

}

func convertColumnOptionType(ctx parser.IColumnOptionContext) tidbast.ColumnOptionType {
	if ctx == nil {
		return tidbast.ColumnOptionNoOption
	}

	switch {
	case ctx.PRIMARY_SYMBOL() != nil:
		return tidbast.ColumnOptionPrimaryKey
	case ctx.NOT_SYMBOL() != nil && ctx.NULL_SYMBOL() != nil:
		return tidbast.ColumnOptionNotNull
	case ctx.AUTO_INCREMENT_SYMBOL() != nil:
		return tidbast.ColumnOptionAutoIncrement
	case ctx.DEFAULT_SYMBOL() != nil && ctx.SERIAL_SYMBOL() == nil:
		return tidbast.ColumnOptionDefaultValue
	case ctx.UNIQUE_SYMBOL() != nil:
		return tidbast.ColumnOptionUniqKey
	case ctx.NOT_SYMBOL() == nil && ctx.NULL_SYMBOL() != nil:
		return tidbast.ColumnOptionNull
	case ctx.ON_SYMBOL() != nil && ctx.UPDATE_SYMBOL() != nil:
		return tidbast.ColumnOptionOnUpdate
	case ctx.COMMENT_SYMBOL() != nil:
		return tidbast.ColumnOptionComment
	case ctx.AS_SYMBOL() != nil:
		return tidbast.ColumnOptionGenerated
	case ctx.References() != nil:
		return tidbast.ColumnOptionReference
	case ctx.COLLATE_SYMBOL() != nil:
		return tidbast.ColumnOptionCollate
	case ctx.CHECK_SYMBOL() != nil:
		return tidbast.ColumnOptionCheck
	case ctx.COLUMN_FORMAT_SYMBOL() != nil:
		return tidbast.ColumnOptionColumnFormat
	case ctx.STORAGE_SYMBOL() != nil:
		return tidbast.ColumnOptionStorage
	case ctx.AUTO_RANDOM_SYMBOL() != nil:
		return tidbast.ColumnOptionAutoRandom
	}

	return tidbast.ColumnOptionNoOption
}

func (s *StringsManipulator) Load(text string) {
	s.l.SetInputStream(antlr.NewInputStream(text))
	s.stream = antlr.NewCommonTokenStream(s.l, antlr.TokenDefaultChannel)
	s.p.SetInputStream(s.stream)
}

func (s *StringsManipulator) ParseColumnDef() error {
	lexerErrorListener := &base.ParseErrorListener{}
	s.p.SetErrorHandler(antlr.NewBailErrorStrategy())
	s.l.RemoveErrorListeners()
	s.l.AddErrorListener(lexerErrorListener)

	parserErrorListener := &base.ParseErrorListener{}
	s.p.RemoveErrorListeners()
	s.p.AddErrorListener(parserErrorListener)

	s.p.BuildParseTrees = true

	s.tree = s.p.SingleColumnDef()

	if lexerErrorListener.Err != nil {
		return lexerErrorListener.Err
	}

	if parserErrorListener.Err != nil {
		return parserErrorListener.Err
	}

	return nil
}

var (
	regexpPattern = regexp.MustCompile("(?m)^-- Table structure for `([^`]+)`")
)

func extractTableNameForCreateTable(s string) (bool, string) {
	matches := regexpPattern.FindStringSubmatch(s)
	if len(matches) > 1 {
		return true, matches[1]
	}
	return false, ""
}
