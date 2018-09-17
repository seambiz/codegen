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
	"database/sql"
	"errors"
	"io"

	"fmt"

	"bitbucket.com/seambiz/seambiz/buffer"
	"bitbucket.com/seambiz/seambiz/sdb"
	"github.com/jmoiron/sqlx"
	"github.com/willf/bitset"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
`)

}
