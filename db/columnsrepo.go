package db

import (
	"database/sql"

	"bitbucket.org/codegen/db/mysql"

	"bitbucket.org/codegen"
)

type ColumnsRepo struct {
	conn *sql.DB
}

func NewColumnsRepo(conn *sql.DB) *ColumnsRepo {
	return &ColumnsRepo{
		conn: conn,
	}
}

func (r ColumnsRepo) Create(data *codegen.Columns) error {
	panic("not implemented") // TODO: Implement
}

func (r ColumnsRepo) Update(data *codegen.Columns) error {
	panic("not implemented") // TODO: Implement
}

func (r ColumnsRepo) Delete(data *codegen.Columns) error {
	panic("not implemented") // TODO: Implement
}

func (r ColumnsRepo) QueryBySchemaAndTable(schema, table string) ([]*codegen.Columns, error) {
	return mysql.NewColumnsStore(r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?) AND UPPER(extra) not like '%VIRTUAL%'").
		OrderBy("ordinal_position").
		Query(schema, table)
}
