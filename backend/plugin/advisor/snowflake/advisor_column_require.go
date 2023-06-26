// Package snowflake is the advisor for snowflake database.
package snowflake

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/snowsql-parser"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
)

var (
	_ advisor.Advisor = (*ColumnRequireAdvisor)(nil)
)

func init() {
	advisor.Register(db.Snowflake, advisor.SnowflakeColumnRequirement, &ColumnRequireAdvisor{})
}

// ColumnRequireAdvisor is the advisor checking for column requirement.
type ColumnRequireAdvisor struct {
}

// Check checks for column requirement.
func (*ColumnRequireAdvisor) Check(ctx advisor.Context, statement string) ([]advisor.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	columnList, err := advisor.UnmarshalRequiredColumnList(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}

	listener := &columnRequireChecker{
		level:          level,
		title:          string(ctx.Rule.Type),
		requireColumns: make(map[string]any),
	}

	for _, column := range columnList {
		listener.requireColumns[column] = true
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// columnRequireChecker is the listener for column requirement.
type columnRequireChecker struct {
	*parser.BaseSnowflakeParserListener

	level advisor.Status
	title string

	adviceList []advisor.Advice

	// requireColumns is the required columns, the key is the normalized column name.
	requireColumns map[string]any

	// The following variables should be clean up when ENTER some statement.
	//
	// currentMissingColumn is the missing column, the key is the normalized column name.
	currentMissingColumn map[string]any
	// currentOriginalTableName is the original table name, should be reset when QUIT some statement.
	currentOriginalTableName string
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *columnRequireChecker) generateAdvice() ([]advisor.Advice, error) {
	if len(l.adviceList) == 0 {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return l.adviceList, nil
}

// EnterCreate_table is called when production create_table is entered.
func (l *columnRequireChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	l.currentOriginalTableName = ctx.Object_name().GetText()
	l.currentMissingColumn = make(map[string]any)
	for column := range l.requireColumns {
		l.currentMissingColumn[column] = true
	}
}

// EnterColumn_decl_item_list is called when production column_decl_item_list is entered.
func (l *columnRequireChecker) EnterColumn_decl_item_list(ctx *parser.Column_decl_item_listContext) {
	if l.currentOriginalTableName == "" {
		return
	}
	allColumnDeclItems := ctx.AllColumn_decl_item()
	for _, columnDeclItem := range allColumnDeclItems {
		if fullColDecl := columnDeclItem.Full_col_decl(); fullColDecl != nil {
			normalizedColumnName := normalizeObjectNamePart(fullColDecl.Col_decl().Column_name().Id_())
			delete(l.currentMissingColumn, normalizedColumnName)
		}
	}
}

// ExitCreate_table is called when production create_table is exited.
func (l *columnRequireChecker) ExitCreate_table(ctx *parser.Create_tableContext) {
	columnNames := make([]string, 0, len(l.currentMissingColumn))
	for column := range l.currentMissingColumn {
		columnNames = append(columnNames, column)
	}
	if len(columnNames) == 0 {
		return
	}

	slices.SortFunc[string](columnNames, func(i, j string) bool {
		return i < j
	})
	for _, column := range columnNames {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NoRequiredColumn,
			Title:   l.title,
			Content: fmt.Sprintf("Table %s missing required column %q", l.currentOriginalTableName, column),
			Line:    ctx.Column_decl_item_list().GetStop().GetLine(),
		})
	}
	l.currentOriginalTableName = ""
	l.currentMissingColumn = nil
}

// EnterAlter_table is called when production alter_table is entered.
func (l *columnRequireChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	l.currentOriginalTableName = ctx.Object_name(0).GetText()
	l.currentMissingColumn = make(map[string]any)
}

func (l *columnRequireChecker) EnterTable_column_action(ctx *parser.Table_column_actionContext) {
	if l.currentOriginalTableName == "" || len(ctx.AllDROP()) != 1 || ctx.Alter_modify() != nil {
		return
	}

	for _, columnName := range ctx.Column_list().AllColumn_name() {
		originalColumName := columnName.GetText()
		normalizedColumnName := extractOrdinaryIdentifier(originalColumName)
		if _, ok := l.requireColumns[normalizedColumnName]; ok {
			l.currentMissingColumn[normalizedColumnName] = true
		}
	}
}

// ExitAlter_table is called when production alter_table is exited.
func (l *columnRequireChecker) ExitAlter_table(ctx *parser.Alter_tableContext) {
	columnNames := make([]string, 0, len(l.currentMissingColumn))
	for column := range l.currentMissingColumn {
		columnNames = append(columnNames, column)
	}
	if len(columnNames) == 0 {
		return
	}

	slices.SortFunc[string](columnNames, func(i, j string) bool {
		return i < j
	})
	for _, column := range columnNames {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NoRequiredColumn,
			Title:   l.title,
			Content: fmt.Sprintf("Table %s missing required column %q", l.currentOriginalTableName, column),
			Line:    ctx.Table_column_action().GetStart().GetLine(),
		})
	}
	l.currentOriginalTableName = ""
	l.currentMissingColumn = nil
}
