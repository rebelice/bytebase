package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"sort"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
)

var (
	_ advisor.Advisor = (*StatementMergeAlterTableAdvisor)(nil)
	_ ast.Visitor     = (*statementMergeAlterTableChecker)(nil)
)

func init() {
	advisor.Register(db.MySQL, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
	advisor.Register(db.TiDB, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
	advisor.Register(db.MariaDB, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
	advisor.Register(db.OceanBase, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
}

// StatementMergeAlterTableAdvisor is the advisor checking for merging ALTER TABLE statements.
type StatementMergeAlterTableAdvisor struct {
}

// Check checks for merging ALTER TABLE statements.
func (*StatementMergeAlterTableAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementMergeAlterTableChecker{
		level:    level,
		title:    string(ctx.Rule.Type),
		tableMap: make(map[string]tableStatement),
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	return checker.generateAdvice(), nil
}

type statementMergeAlterTableChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	text       string
	line       int
	tableMap   map[string]tableStatement
}

type tableStatement struct {
	name     string
	count    int
	lastLine int
}

// Enter implements the ast.Visitor interface.
func (checker *statementMergeAlterTableChecker) Enter(in ast.Node) (ast.Node, bool) {
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		data := tableStatement{
			name:     node.Table.Name.O,
			count:    1,
			lastLine: checker.line,
		}
		checker.tableMap[node.Table.Name.O] = data
	case *ast.AlterTableStmt:
		data, ok := checker.tableMap[node.Table.Name.O]
		if !ok {
			data = tableStatement{
				name:  node.Table.Name.O,
				count: 0,
			}
		}
		data.count++
		data.lastLine = checker.line
		checker.tableMap[node.Table.Name.O] = data
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*statementMergeAlterTableChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (checker *statementMergeAlterTableChecker) generateAdvice() []advisor.Advice {
	var tableList []tableStatement
	for _, table := range checker.tableMap {
		tableList = append(tableList, table)
	}
	sort.Slice(tableList, func(i, j int) bool {
		return tableList[i].lastLine < tableList[j].lastLine
	})

	for _, table := range tableList {
		if table.count > 1 {
			checker.adviceList = append(checker.adviceList, advisor.Advice{
				Status:  checker.level,
				Code:    advisor.StatementRedundantAlterTable,
				Title:   checker.title,
				Content: fmt.Sprintf("There are %d statements to modify table `%s`", table.count, table.name),
				Line:    table.lastLine,
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
