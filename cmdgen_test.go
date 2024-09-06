package codegen

import (
	"reflect"
	"testing"

	"github.com/seambiz/codegen/config"
	"github.com/stretchr/testify/assert"
)

func Test_generateTemplatesConfig(t *testing.T) {
	expectedSchemaTemplatefiles := map[string][]string{}
	expectedTableTemplatefiles := map[string][]string{}

	expectedSchemaTemplatefiles["once.root.dbconstants"] = []string{"once.root.dbconstants.tmpl"}
	expectedSchemaTemplatefiles["once.root.audittype"] = []string{"once.root.audittype.tmpl"}
	expectedSchemaTemplatefiles["once.package.fakedb_test"] = []string{"once.package.fakedb_test.tmpl"}
	expectedSchemaTemplatefiles["once.package.shared"] = []string{"once.package.shared.tmpl"}
	expectedSchemaTemplatefiles["once.package.store"] = []string{"once.package.store.tmpl"}
	expectedSchemaTemplatefiles["once.package.store_timing_test"] = []string{"once.package.store_timing_test.tmpl"}
	expectedSchemaTemplatefiles["once.root.dbtypes"] = []string{"once.root.dbtypes.tmpl"}

	expectedTableTemplatefiles["table.package.store"] = []string{"table.package.store.01-header.tmpl", "table.package.store.02-type.tmpl", "table.package.store.03-queryfields.tmpl", "table.package.store.04-store.tmpl", "table.package.store.05-bind.tmpl", "table.package.store.06-select.tmpl", "table.package.store.07-queries-custom.tmpl", "table.package.store.07-queries-one.tmpl", "table.package.store.07-queries.tmpl", "table.package.store.08-foreigndata.tmpl", "table.package.store.09-upsert.tmpl", "table.package.store.10-insert.tmpl", "table.package.store.11-update.tmpl", "table.package.store.12-delete.tmpl", "table.package.store.13-truncate.tmpl", "table.package.store.14-indexqueries.tmpl", "table.package.store.15-json.tmpl", "table.package.store.16-footer.tmpl"}
	expectedTableTemplatefiles["table.package.store_test"] = []string{"table.package.store_test.tmpl"}
	expectedTableTemplatefiles["table.package.store_timing_test"] = []string{"table.package.store_timing_test.tmpl"}
	expectedTableTemplatefiles["table.root.repository"] = []string{"table.root.repository.tmpl"}
	expectedTableTemplatefiles["table.subpackage.repo"] = []string{"table.subpackage.repo.tmpl"}

	type args struct {
		conf *config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "templates folder",
			args: args{
				conf: &config.Config{
					Schemas: []*config.Schema{
						{
							Name:           "testschema",
							TemplateFolder: "templates",
							Tables: []*config.Table{
								{Name: "testtable", Generate: true},
							},
						},
					},
				},
			},
		},
		{
			name: "static templates",
			args: args{
				conf: &config.Config{
					Schemas: []*config.Schema{
						{
							Name:           "static schema",
							TemplateFolder: "",
							Tables: []*config.Table{
								{Name: "static table", Generate: true},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateTemplatesConfig(tt.args.conf)
			for _, schema := range tt.args.conf.Schemas {
				assert.EqualValuesf(t, expectedSchemaTemplatefiles, schema.PreparedTemplatefiles, "%v failed")

				for _, table := range schema.Tables {
					if !reflect.DeepEqual(table.PreparedTemplatefiles, expectedTableTemplatefiles) {
						assert.EqualValuesf(t, expectedTableTemplatefiles, table.PreparedTemplatefiles, "%v failed")
					}
				}
			}
		})
	}
}
