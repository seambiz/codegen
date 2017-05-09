package codegen

// THeader template
func THeader(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.S("// Package ")
	bb.S(conf.Package)
	bb.S(" contains the types for schema '")
	bb.S(schema.Name)
	bb.S("'.")
	bb.NewLine()
	bb.S("package ")
	bb.S(conf.Package)
	bb.NewLine()
	bb.NewLine()
	bb.S(`import (
        "bitbucket.com/seambiz/logging"
        "bitbucket.com/seambiz/sdb"
        "github.com/jmoiron/sqlx"
		"go.uber.org/zap"
)
`)

}
