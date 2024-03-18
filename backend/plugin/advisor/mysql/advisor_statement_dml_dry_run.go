package mysql

// Framework code is generated by the generator.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*StatementDmlDryRunAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLStatementDMLDryRun, &StatementDmlDryRunAdvisor{})
}

// StatementDmlDryRunAdvisor is the advisor checking for DML dry run.
type StatementDmlDryRunAdvisor struct {
}

// Check checks for DML dry run.
func (*StatementDmlDryRunAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementDmlDryRunChecker{
		level:  level,
		title:  string(ctx.Rule.Type),
		driver: ctx.Driver,
		ctx:    ctx.Context,
	}

	if checker.driver != nil {
		for _, stmt := range stmtList {
			checker.baseLine = stmt.BaseLine
			antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
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
	return checker.adviceList, nil
}

type statementDmlDryRunChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	line       int
	driver     *sql.DB
	ctx        context.Context
}

// EnterUpdateStatement is called when production updateStatement is entered.
func (checker *statementDmlDryRunChecker) EnterUpdateStatement(ctx *mysql.UpdateStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	checker.handleStmt(ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx), ctx.GetStart().GetLine())
}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (checker *statementDmlDryRunChecker) EnterDeleteStatement(ctx *mysql.DeleteStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	checker.handleStmt(ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx), ctx.GetStart().GetLine())
}

// EnterInsertStatement is called when production insertStatement is entered.
func (checker *statementDmlDryRunChecker) EnterInsertStatement(ctx *mysql.InsertStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	checker.handleStmt(ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx), ctx.GetStart().GetLine())
}

// Enter implements the ast.Visitor interface.
func (checker *statementDmlDryRunChecker) handleStmt(text string, lineNumber int) {
	if _, err := advisor.Query(checker.ctx, checker.driver, fmt.Sprintf("EXPLAIN %s", text)); err != nil {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.StatementDMLDryRunFailed,
			Title:   checker.title,
			Content: fmt.Sprintf("\"%s\" dry runs failed: %s", text, err.Error()),
			Line:    checker.line + lineNumber,
		})
	}
}
