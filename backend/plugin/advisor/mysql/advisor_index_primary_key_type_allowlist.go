package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
	bbparser "github.com/bytebase/bytebase/backend/plugin/parser/sql"
)

var (
	_ advisor.Advisor = (*IndexPrimaryKeyTypeAllowlistAdvisor)(nil)
	_ ast.Visitor     = (*indexPrimaryKeyTypeAllowlistChecker)(nil)
)

func init() {
	advisor.Register(db.MySQL, advisor.MySQLPrimaryKeyTypeAllowlist, &IndexPrimaryKeyTypeAllowlistAdvisor{})
	advisor.Register(db.TiDB, advisor.MySQLPrimaryKeyTypeAllowlist, &IndexPrimaryKeyTypeAllowlistAdvisor{})
	advisor.Register(db.OceanBase, advisor.MySQLPrimaryKeyTypeAllowlist, &IndexPrimaryKeyTypeAllowlistAdvisor{})
}

// IndexPrimaryKeyTypeAllowlistAdvisor is the advisor checking for primary key type allowlist.
type IndexPrimaryKeyTypeAllowlistAdvisor struct {
}

// Check checks for primary key type allowlist.
func (*IndexPrimaryKeyTypeAllowlistAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalStringArrayTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	allowlist := make(map[string]bool)
	for _, tp := range payload.List {
		allowlist[strings.ToLower(tp)] = true
	}
	checker := &indexPrimaryKeyTypeAllowlistChecker{
		level:            level,
		title:            string(ctx.Rule.Type),
		allowlist:        allowlist,
		catalog:          ctx.Catalog,
		tablesNewColumns: make(map[string]columnNameToColumnDef),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type indexPrimaryKeyTypeAllowlistChecker struct {
	adviceList       []advisor.Advice
	level            advisor.Status
	title            string
	text             string
	line             int
	allowlist        map[string]bool
	catalog          *catalog.Finder
	tablesNewColumns tableNewColumn
}

// Enter implements the ast.Visitor interface.
func (v *indexPrimaryKeyTypeAllowlistChecker) Enter(in ast.Node) (ast.Node, bool) {
	var pkDataList []pkData
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		tableName := node.Table.Name.String()
		for _, column := range node.Cols {
			pds := v.addNewColumn(tableName, column.OriginTextPosition(), column)
			pkDataList = append(pkDataList, pds...)
		}
		for _, constraint := range node.Constraints {
			pds := v.addConstraint(tableName, constraint.OriginTextPosition(), constraint)
			pkDataList = append(pkDataList, pds...)
		}
	case *ast.AlterTableStmt:
		tableName := node.Table.Name.String()
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					pds := v.addNewColumn(tableName, node.OriginTextPosition(), column)
					pkDataList = append(pkDataList, pds...)
				}
			case ast.AlterTableAddConstraint:
				pds := v.addConstraint(tableName, node.OriginTextPosition(), spec.Constraint)
				pkDataList = append(pkDataList, pds...)
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				newColumnDef := spec.NewColumns[0]
				oldColumnName := newColumnDef.Name.Name.String()
				if spec.OldColumnName != nil {
					oldColumnName = spec.OldColumnName.Name.String()
				}
				pds := v.changeColumn(tableName, oldColumnName, node.OriginTextPosition(), newColumnDef)
				pkDataList = append(pkDataList, pds...)
			}
		}
	}
	for _, pd := range pkDataList {
		v.adviceList = append(v.adviceList, advisor.Advice{
			Status:  v.level,
			Code:    advisor.IndexPKType,
			Title:   v.title,
			Content: fmt.Sprintf("The column `%s` in table `%s` is one of the primary key, but its type \"%s\" is not in allowlist", pd.column, pd.table, pd.columnType),
			Line:    pd.line,
		})
	}
	return in, false
}

// Leave implements the ast.Visitor interface.
func (*indexPrimaryKeyTypeAllowlistChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (v *indexPrimaryKeyTypeAllowlistChecker) addNewColumn(tableName string, line int, colDef *ast.ColumnDef) []pkData {
	var pkDataList []pkData
	for _, option := range colDef.Options {
		if option.Tp == ast.ColumnOptionPrimaryKey {
			tp := bbparser.TypeString(colDef.Tp.GetType())
			if _, exists := v.allowlist[strings.ToLower(tp)]; !exists {
				pkDataList = append(pkDataList, pkData{
					table:      tableName,
					column:     colDef.Name.String(),
					columnType: tp,
					line:       line,
				})
			}
		}
	}
	v.tablesNewColumns.set(tableName, colDef.Name.String(), colDef)
	return pkDataList
}

func (v *indexPrimaryKeyTypeAllowlistChecker) changeColumn(tableName, oldColumnName string, line int, newColumnDef *ast.ColumnDef) []pkData {
	var pkDataList []pkData
	v.tablesNewColumns.delete(tableName, oldColumnName)
	for _, option := range newColumnDef.Options {
		if option.Tp == ast.ColumnOptionPrimaryKey {
			tp := bbparser.TypeString(newColumnDef.Tp.GetType())
			if _, exists := v.allowlist[strings.ToLower(tp)]; !exists {
				pkDataList = append(pkDataList, pkData{
					table:      tableName,
					column:     newColumnDef.Name.String(),
					columnType: tp,
					line:       line,
				})
			}
		}
	}
	v.tablesNewColumns.set(tableName, newColumnDef.Name.String(), newColumnDef)
	return pkDataList
}

func (v *indexPrimaryKeyTypeAllowlistChecker) addConstraint(tableName string, line int, constraint *ast.Constraint) []pkData {
	var pkDataList []pkData
	if constraint.Tp == ast.ConstraintPrimaryKey {
		for _, key := range constraint.Keys {
			columnName := key.Column.Name.String()
			columnType, err := v.getPKColumnType(tableName, columnName)
			if err != nil {
				continue
			}
			if _, exists := v.allowlist[strings.ToLower(columnType)]; !exists {
				pkDataList = append(pkDataList, pkData{
					table:      tableName,
					column:     columnName,
					columnType: columnType,
					line:       line,
				})
			}
		}
	}
	return pkDataList
}

// getPKColumnType gets the column type string from v.tablesNewColumns or catalog, returns empty string and non-nil error if cannot find the column in given table.
func (v *indexPrimaryKeyTypeAllowlistChecker) getPKColumnType(tableName string, columnName string) (string, error) {
	if colDef, ok := v.tablesNewColumns.get(tableName, columnName); ok {
		return bbparser.TypeString(colDef.Tp.GetType()), nil
	}
	column := v.catalog.Origin.FindColumn(&catalog.ColumnFind{
		TableName:  tableName,
		ColumnName: columnName,
	})
	if column != nil {
		return column.Type(), nil
	}
	return "", errors.Errorf("cannot find the type of `%s`.`%s`", tableName, columnName)
}
