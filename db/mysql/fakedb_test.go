package mysql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"strconv"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

var (
	dsns    = map[string]QueryResult{}
	counter = 0
)

type QueryResult struct {
	*ResultRows
	*Query
	NumInput int
}

type ResultRows struct {
	NumRows int
}

type Query struct {
	Cols []string
	Vals [][]driver.Value
}

type mimic struct{}

func (m *mimic) Open(dsn string) (driver.Conn, error) {
	if len(dsn) == 0 {
		dsn = strconv.Itoa(counter)
		counter++
	}
	return &mimicConn{dsns[dsn]}, nil
}

type mimicConn struct {
	Q QueryResult
}

func (m *mimicConn) Prepare(query string) (driver.Stmt, error) {
	return &mimicStmt{m.Q}, nil
}

func (m *mimicConn) Close() error              { return nil }
func (m *mimicConn) Begin() (driver.Tx, error) { return nil, errors.New("tx not supported") }

type mimicStmt struct {
	Q QueryResult
}

func (m *mimicStmt) Close() error  { return nil }
func (m *mimicStmt) NumInput() int { return m.Q.NumInput }
func (m *mimicStmt) Exec(args []driver.Value) (driver.Result, error) {
	if m.Q.ResultRows == nil {
		return nil, errors.New("statement was not a result type")
	}

	return &mimicResult{m.Q.ResultRows.NumRows}, nil
}

func (m *mimicStmt) Query(args []driver.Value) (driver.Rows, error) {
	if m.Q.Query == nil {
		return nil, errors.New("statement was not a query type")
	}

	return &mimicRows{columns: m.Q.Query.Cols, values: m.Q.Query.Vals}, nil
}

type mimicResult struct {
	rowsAffected int
}

func (m *mimicResult) LastInsertId() (int64, error) {
	return 0, errors.New("not supported")
}

func (m *mimicResult) RowsAffected() (int64, error) {
	return int64(m.rowsAffected), nil
}

type mimicRows struct {
	cursor  int
	columns []string
	values  [][]driver.Value
}

func (m *mimicRows) Columns() []string { return m.columns }
func (m *mimicRows) Close() error      { return nil }
func (m *mimicRows) Next(dest []driver.Value) error {
	if m.cursor == len(m.values) {
		return io.EOF
	}

	for i, val := range m.values[m.cursor] {
		dest[i] = val
	}
	m.cursor++

	return nil
}

func init() {
	sql.Register("mimic", &mimic{})
}

func newResult(q QueryResult) {
	dsns[strconv.Itoa(counter)] = q
}

func newQuery(q QueryResult) {
	dsns[strconv.Itoa(counter)] = q
}

func newResultDSN(dsn string, q QueryResult) {
	dsns[dsn] = q
}

func addResultRowDSN(dsn string, row []driver.Value) {
	dsns[dsn].Vals = append(dsns[dsn].Vals, row)
}

func newQueryDSN(dsn string, q QueryResult) {
	dsns[dsn] = q
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
