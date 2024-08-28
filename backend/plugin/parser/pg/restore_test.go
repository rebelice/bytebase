package pg

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	"github.com/bytebase/bytebase/backend/store/model"
	"github.com/bytebase/bytebase/proto/generated-go/store"
)

type restoreCase struct {
	Input            string
	BackupDatabase   string
	BackupTable      string
	OriginalDatabase string
	OriginalTable    string
	Result           string
}

func TestRestore(t *testing.T) {
	tests := []restoreCase{}

	const (
		record = true
	)
	var (
		filepath = "test-data/test_restore.yaml"
	)

	a := require.New(t)
	yamlFile, err := os.Open(filepath)
	a.NoError(err)

	byteValue, err := io.ReadAll(yamlFile)
	a.NoError(yamlFile.Close())
	a.NoError(err)
	a.NoError(yaml.Unmarshal(byteValue, &tests))

	for i, t := range tests {
		result, err := GenerateRestoreSQL(context.Background(), base.RestoreContext{
			GetDatabaseMetadataFunc: fixedMockDatabaseMetadataGetter,
		}, t.Input, &store.PriorBackupDetail_Item{
			SourceTable: &store.PriorBackupDetail_Item_Table{
				Database: "instances/i1/databases/" + t.OriginalDatabase,
				Schema:   "public",
				Table:    t.OriginalTable,
			},
			TargetTable: &store.PriorBackupDetail_Item_Table{
				Database: "instances/i1/databases/" + t.BackupDatabase,
				Schema:   t.BackupDatabase,
				Table:    t.BackupTable,
			},
			StartPosition: &store.Position{
				Line:   1,
				Column: 1,
			},
			EndPosition: &store.Position{
				Line:   1000000000,
				Column: 1,
			},
		})
		a.NoError(err)

		if record {
			tests[i].Result = result
		} else {
			a.Equal(t.Result, result, t.Input)
		}
	}
	if record {
		byteValue, err := yaml.Marshal(tests)
		a.NoError(err)
		err = os.WriteFile(filepath, byteValue, 0644)
		a.NoError(err)
	}
}

func fixedMockDatabaseMetadataGetter(_ context.Context, _ string, database string) (string, *model.DatabaseMetadata, error) {
	return database, model.NewDatabaseMetadata(&store.DatabaseSchemaMetadata{
		Name: database,
		Schemas: []*store.SchemaMetadata{
			{
				Name: "public",
				Tables: []*store.TableMetadata{
					{
						Name: "t_generated",
						Columns: []*store.ColumnMetadata{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
						},
						Indexes: []*store.IndexMetadata{
							{
								Name:        "t_generated_pk",
								Expressions: []string{"b"},
							},
						},
					},
					{
						Name: "t1",
						Columns: []*store.ColumnMetadata{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
						},
					},
					{
						Name: "t2",
						Columns: []*store.ColumnMetadata{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
						},
					},
					{
						Name: "test",
						Columns: []*store.ColumnMetadata{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
						},
						Indexes: []*store.IndexMetadata{
							{
								Name:        "test_pk",
								Expressions: []string{"a"},
							},
						},
					},
					{
						Name: "test2",
						Columns: []*store.ColumnMetadata{
							{
								Name: "a",
							},
							{
								Name: "b",
							},
							{
								Name: "c",
							},
						},
					},
				},
			},
		},
	}), nil
}
