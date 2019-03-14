package codegen

// THeader template
func THeader(bb *GenBuffer, conf *Config, schema *Schema) {
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
	"math/big"

	"fmt"

	"bitbucket.org/seambiz/seambiz/buffer"
	"bitbucket.org/seambiz/seambiz/sdb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
`)

}
