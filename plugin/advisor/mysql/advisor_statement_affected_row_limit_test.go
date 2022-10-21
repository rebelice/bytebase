package mysql

// Framework code is generated by the generator.

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestStatementAffectedRowLimit(t *testing.T) {
	tests := []advisor.TestCase{
		{
			Statement: `DELETE FROM tech_book WHERE a > 1`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			Statement: `UPDATE tech_book SET id = 1`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.StatementAffectedRowExceedsLimit,
					Title:   "statement.affected-row-limit",
					Content: "\"UPDATE tech_book SET id = 1\" affected 1000 rows. The count exceeds 5.",
					Line:    1,
				},
			},
		},
	}

	payload, err := json.Marshal(advisor.NumberTypeRulePayload{
		Number: 5,
	})
	require.NoError(t, err)
	advisor.RunSQLReviewRuleTests(t, tests, &StatementAffectedRowLimitAdvisor{}, &advisor.SQLReviewRule{
		Type:    advisor.SchemaRuleStatementAffectedRowLimit,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: string(payload),
	}, advisor.MockMySQLDatabase)
}
