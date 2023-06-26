package pg

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
	"github.com/bytebase/bytebase/backend/plugin/parser/sql/ast"
)

var (
	_ advisor.Advisor = (*CommentConventionAdvisor)(nil)
	_ ast.Visitor     = (*commentConventionChecker)(nil)
)

func init() {
	advisor.Register(db.Postgres, advisor.PostgreSQLCommentConvention, &CommentConventionAdvisor{})
}

// CommentConventionAdvisor is the advisor checking for comment convention.
type CommentConventionAdvisor struct {
}

// Check checks for comment convention.
func (*CommentConventionAdvisor) Check(ctx advisor.Context, statement string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.Node)
	if !ok {
		return nil, fmt.Errorf("failed to convert to Node")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &commentConventionChecker{
		level:     level,
		title:     string(ctx.Rule.Type),
		maxLength: payload.Number,
	}

	for _, stmt := range stmtList {
		ast.Walk(checker, stmt)
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

type commentConventionChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	maxLength  int
}

// Visit implements the ast.Visitor interface.
func (checker *commentConventionChecker) Visit(node ast.Node) ast.Visitor {
	type commentData struct {
		comment string
		line    int
	}
	var comment commentData

	if n, ok := node.(*ast.CommentStmt); ok {
		comment = commentData{
			comment: n.Comment,
			line:    n.LastLine(),
		}
	}

	if checker.maxLength > 0 && len(comment.comment) > checker.maxLength {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.CommentTooLong,
			Title:   checker.title,
			Content: fmt.Sprintf("The length of comment should be within %d characters", checker.maxLength),
			Line:    comment.line,
		})
	}

	return checker
}
