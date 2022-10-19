package mysql

// Framework code is generated by the generator.

import (
	"testing"

	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestIndexPkType(t *testing.T) {
	tests := []advisor.TestCase{
		// CREATE TABLE COLUMN OPTION PRIMARY KEY
		{
			Statement: `CREATE TABLE t(id VARCHAR(5) PRIMARY KEY)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    1,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(id INT PRIMARY KEY)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// CREATE TABLE TABLE OPTION PRIAMRY KEY
		{
			Statement: `CREATE TABLE t(id VARCHAR(5), PRIMARY KEY(id))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    1,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(id INT, id2 VARCHAR(5), id3 VARCHAR(5), PRIMARY KEY(id, id2, id3))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id2` is varchar(5)",
					Line:    1,
				},
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id3` is varchar(5)",
					Line:    1,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(id INT, id2 BIGINT, PRIMARY KEY(id, id2))`,
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
			Statement: `CREATE TABLE t(id INT, PRIMARY KEY(id))`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// ALTER TABLE ADD COLUMN OPTION PRIMARY KEY
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id VARCHAR(5) PRIMARY KEY`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id INT PRIMARY KEY`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// ALTER TABLE ADD COLUMN ADD PRIAMRY KEY CONSTRAINT
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id VARCHAR(5), ADD PRIMARY KEY (id)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id INT, ADD COLUMN id2 VARCHAR(5), ADD COLUMN id3 VARCHAR(5), ADD PRIMARY KEY (id, id2, id3)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id2` is varchar(5)",
					Line:    3,
				},
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id3` is varchar(5)",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id INT, ADD COLUMN id2 BIGINT, ADD PRIMARY KEY (id, id2)`,
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
			Statement: `
				CREATE TABLE t(a int);
				ALTER TABLE t ADD COLUMN id BIGINT, ADD PRIMARY KEY (id)`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// CRAETE TABLE; ALTER TABLE ADD PRIMARY KEY CONSTRAINT
		{
			Statement: `CREATE TABLE t(id VARCHAR(5));
				ALTER TABLE t ADD PRIMARY KEY(id);
			`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    2,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(id INT, id2 VARCHAR(5), id3 VARCHAR(5));
				ALTER TABLE t ADD PRIMARY KEY(id, id2, id3);
			`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id2` is varchar(5)",
					Line:    2,
				},
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id3` is varchar(5)",
					Line:    2,
				},
			},
		},
		{
			Statement: `CREATE TABLE t(id INT, id2 BIGINT);
				ALTER TABLE t ADD PRIMARY KEY(id, id2);
			`,
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
			Statement: `CREATE TABLE t(id BIGINT);
			ALTER TABLE t ADD PRIMARY KEY(id);
			`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// ALTER TABLE MODIFY COLUMN
		{
			Statement: `
				CREATE TABLE t(id int);
				ALTER TABLE t MODIFY COLUMN id VARCHAR(5) PRIMARY KEY;`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id` is varchar(5)",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(id int);
				ALTER TABLE t MODIFY COLUMN id INT(5) PRIMARY KEY;`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		// ALTER TABLE CHANGE COLUMN
		{
			Statement: `
				CREATE TABLE t(id int);
				ALTER TABLE t CHANGE COLUMN id id2 VARCHAR(5) PRIMARY KEY`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    advisor.IndexPKType,
					Title:   "index.pk-type-limit",
					Content: "Columns in primary key must be INT/BIGINT but `t`.`id2` is varchar(5)",
					Line:    3,
				},
			},
		},
		{
			Statement: `
				CREATE TABLE t(id int);
				ALTER TABLE t CHANGE COLUMN id id2 INT(5) PRIMARY KEY`,
			Want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    advisor.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
	}

	advisor.RunSQLReviewRuleTests(t, tests, &IndexPkTypeAdvisor{}, &advisor.SQLReviewRule{
		Type:    advisor.SchemaRuleIndexPKTypeLimit,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: "",
	}, advisor.MockMySQLDatabase)
}
