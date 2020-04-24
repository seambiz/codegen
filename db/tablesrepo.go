package db

import (
	"database/sql"

	"bitbucket.org/codegen/db/mysql"

	"bitbucket.org/codegen"
)

type TablesRepo struct {
	conn *sql.DB
}

func NewTablesRepo(conn *sql.DB) *TablesRepo {
	return &TablesRepo{
		conn: conn,
	}
}

func (r TablesRepo) Create(data *codegen.Tables) error {
	panic("not implemented") // TODO: Implement
}

func (r TablesRepo) Update(data *codegen.Tables) error {
	panic("not implemented") // TODO: Implement
}

func (r TablesRepo) UpdatePartial(data *codegen.TablesPartial) error {
	panic("not implemented") // TODO: Implement
}

func (r TablesRepo) Delete(data *codegen.Tables) error {
	panic("not implemented") // TODO: Implement
}

func (r TablesRepo) QueryBySchema(schema string) ([]*codegen.Tables, error) {
	return mysql.NewTablesStore(r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND table_type IN (UPPER('base table'), UPPER('system view'))").
		Query(schema)
}
