package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"sort"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
)

var (
	_ advisor.Advisor = (*IndexTotalNumberLimitAdvisor)(nil)
	_ ast.Visitor     = (*indexTotalNumberLimitChecker)(nil)
)

func init() {
	advisor.Register(db.MySQL, advisor.MySQLIndexTotalNumberLimit, &IndexTotalNumberLimitAdvisor{})
	advisor.Register(db.TiDB, advisor.MySQLIndexTotalNumberLimit, &IndexTotalNumberLimitAdvisor{})
	advisor.Register(db.MariaDB, advisor.MySQLIndexTotalNumberLimit, &IndexTotalNumberLimitAdvisor{})
	advisor.Register(db.OceanBase, advisor.MySQLIndexTotalNumberLimit, &IndexTotalNumberLimitAdvisor{})
}

// IndexTotalNumberLimitAdvisor is the advisor checking for index total number limit.
type IndexTotalNumberLimitAdvisor struct {
}

// Check checks for index total number limit.
func (*IndexTotalNumberLimitAdvisor) Check(ctx advisor.Context, statement string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &indexTotalNumberLimitChecker{
		level:        level,
		title:        string(ctx.Rule.Type),
		max:          payload.Number,
		lineForTable: make(map[string]int),
		catalog:      ctx.Catalog,
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.generateAdvice(), nil
}

type indexTotalNumberLimitChecker struct {
	adviceList   []advisor.Advice
	level        advisor.Status
	title        string
	text         string
	line         int
	max          int
	lineForTable map[string]int
	catalog      *catalog.Finder
}

func (checker *indexTotalNumberLimitChecker) generateAdvice() []advisor.Advice {
	type tableName struct {
		name string
		line int
	}
	var tableList []tableName

	for k, v := range checker.lineForTable {
		tableList = append(tableList, tableName{
			name: k,
			line: v,
		})
	}
	sort.Slice(tableList, func(i, j int) bool {
		return tableList[i].line < tableList[j].line
	})

	for _, table := range tableList {
		tableInfo := checker.catalog.Final.FindTable(&catalog.TableFind{TableName: table.name})
		if tableInfo != nil && tableInfo.CountIndex() > checker.max {
			checker.adviceList = append(checker.adviceList, advisor.Advice{
				Status:  checker.level,
				Code:    advisor.IndexCountExceedsLimit,
				Title:   checker.title,
				Content: fmt.Sprintf("The count of index in table `%s` should be no more than %d, but found %d", table.name, checker.max, tableInfo.CountIndex()),
				Line:    table.line,
			})
		}
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList
}

// Enter implements the ast.Visitor interface.
func (checker *indexTotalNumberLimitChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		checker.lineForTable[node.Table.Name.O] = node.OriginTextPosition()
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					if createIndex(column) {
						checker.lineForTable[node.Table.Name.O] = node.OriginTextPosition()
						break
					}
				}
			case ast.AlterTableAddConstraint:
				if createIndex(spec.Constraint) {
					checker.lineForTable[node.Table.Name.O] = node.OriginTextPosition()
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				if createIndex(spec.NewColumns[0]) {
					checker.lineForTable[node.Table.Name.O] = node.OriginTextPosition()
				}
			}
		}
	case *ast.CreateIndexStmt:
		checker.lineForTable[node.Table.Name.O] = node.OriginTextPosition()
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*indexTotalNumberLimitChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func createIndex(in ast.Node) bool {
	switch node := in.(type) {
	case *ast.ColumnDef:
		for _, option := range node.Options {
			switch option.Tp {
			case ast.ColumnOptionPrimaryKey, ast.ColumnOptionUniqKey:
				return true
			}
		}
	case *ast.Constraint:
		switch node.Tp {
		case ast.ConstraintPrimaryKey,
			ast.ConstraintUniq,
			ast.ConstraintUniqKey,
			ast.ConstraintUniqIndex,
			ast.ConstraintKey,
			ast.ConstraintIndex,
			ast.ConstraintFulltext:
			return true
		}
	}
	return false
}
