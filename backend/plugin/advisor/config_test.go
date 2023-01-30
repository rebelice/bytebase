package advisor

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

var mockConfigOverrideYAMLStr = `
template: bb.sql-review.prod # Provide the template id, then we can extend rules from the specific template.
ruleList:
  - type: mysql.statement.select.no-select-all
    level: DISABLED
  - type: mysql.table.drop-naming-convention
    level: WARNING
  - type: mysql.table.require-pk
    level: TEST
  - type: mysql.naming.table
    payload:
      format: "^table_[a-z]+(_[a-z]+)*$"
  - type: mysql.naming.column
    payload:
      maxLength: 24
  - type: mysql.column.required
    level: ERROR
    payload:
      list:
        - name
`

func TestConfigOverride(t *testing.T) {
	override := &SQLReviewConfigOverride{}
	err := yaml.Unmarshal([]byte(mockConfigOverrideYAMLStr), override)
	require.NoError(t, err)

	ruleList, err := MergeSQLReviewRules(override)
	require.NoError(t, err)

	for _, rule := range ruleList {
		switch rule.Type {
		case "mysql.statement.select.no-select-all":
			assert.Equal(t, SchemaRuleLevelDisabled, rule.Level)
		case "mysql.table.drop-naming-convention":
			assert.Equal(t, SchemaRuleLevelWarning, rule.Level)
		case "mysql.table.require-pk":
			assert.Equal(t, SchemaRuleLevelError, rule.Level)
		case "mysql.naming.table":
			assert.Equal(t, SchemaRuleLevelWarning, rule.Level)

			var nr NamingRulePayload
			err := json.Unmarshal([]byte(rule.Payload), &nr)
			require.NoError(t, err)

			assert.Equal(t, "^table_[a-z]+(_[a-z]+)*$", nr.Format)
			assert.Equal(t, 63, nr.MaxLength)
		case "mysql.naming.column":
			assert.Equal(t, SchemaRuleLevelWarning, rule.Level)

			var nr NamingRulePayload
			err := json.Unmarshal([]byte(rule.Payload), &nr)
			require.NoError(t, err)

			assert.Equal(t, "^[a-z]+(_[a-z]+)*$", nr.Format)
			assert.Equal(t, 24, nr.MaxLength)
		case "mysql.column.required":
			assert.Equal(t, SchemaRuleLevelError, rule.Level)

			var payload StringArrayTypeRulePayload
			err := json.Unmarshal([]byte(rule.Payload), &payload)
			require.NoError(t, err)

			assert.Equal(t, 1, len(payload.List))
			assert.Equal(t, "name", payload.List[0])
		}
	}
}
