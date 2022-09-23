package templates

import (
	"bytes"
	"go/format"
	"os"
	"strings"
	"testing"

	"bitbucket.org/codegen"
	"bitbucket.org/codegen/gen"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func compareToFile(t *testing.T, filename string, code []byte) (string, bool) {
	golden, err := os.ReadFile(filename)
	if err != nil {
		t.Error("could not read file", filename)
	}

	return string(golden), bytes.Equal(golden, code)
}

func generatorTestCase(t *testing.T, gen gen.Codegen, conf codegen.Config, goldenFile string) {
	t.Run(goldenFile, func(t *testing.T) {
		uglyCode := gen.Generate(conf)

		prettyCode, err := format.Source([]byte(uglyCode))
		if err != nil {
			prettyCode = []byte(uglyCode)
		}

		newFilename := strings.Replace(goldenFile, ".go", "_new.txt", 1)

		err = os.WriteFile(newFilename, prettyCode, 0o644)
		if err != nil {
			t.Fatal(err)
		}

		if golden, ok := compareToFile(t, goldenFile, prettyCode); !ok {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(golden), string(prettyCode), false)
			t.Error(dmp.DiffPrettyText(diffs))
		}
	})
}

func TestTemplates_Generate(t *testing.T) {
	conf := codegen.Config{
		RootPackage: "testing",
	}

	t1 := &codegen.Table{
		Generate: true,
		Title:    "TestTable1",
		Name:     "test_table_1",
		Fields: []*codegen.Field{
			{Title: "First", Name: "first", GoType: "string"},
			{Title: "Second", Name: "second", GoType: "int"},
			{Title: "PtrSecond", Name: "ptr_second", GoType: "int", IsNullable: true},
		},
	}

	t2 := &codegen.Table{
		Generate: true,
		Title:    "TestTable2",
		Name:     "test_table_2",
		Fields: []*codegen.Field{
			{Title: "Third", Name: "third", GoType: "string"},
			{Title: "Fourth", Name: "fourth", GoType: "string"},
		},
	}

	schema := &codegen.Schema{
		Name: "my_schema",
	}
	schema.Tables = append(schema.Tables, t1, t2)
	conf.Schemas = append(conf.Schemas, schema)

	generatorTestCase(t, Constants{}, conf, "test/constants.go")
	generatorTestCase(t, DBTypes{}, conf, "test/dbtypes.go")
	generatorTestCase(t, Router{}, conf, "test/router.go")
}
