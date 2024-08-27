package mysql

import (
	"database/sql"
	"math/big"
	"time"
)

// GENERATED BY CODEGEN.

// Bindable defines the interface to make the query results bindable to the structs
type Bindable interface {
	bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int)
	new() Bindable
}

type BindableSlice interface {
	append(Bindable)
}

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func logString(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func logFloat32(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

func logFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func logInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func logUInt(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}

func logInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func logUInt64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

func logTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func logBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func logBytes(b *[]byte) []byte {
	if b == nil {
		return []byte("<nil>")
	}
	return *b
}


// ^^ END OF GENERATED BY CODEGEN. ^^
