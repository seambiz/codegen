package db

import (
	"database/sql"

	"bitbucket.org/codegen"
	"bitbucket.org/codegen/db/mysql"
)

// GENERATED BY CODEGEN. DO NOT EDIT.
// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^

/* TablesRepo implemente TablesRepository interface definition. */
type TablesRepo struct {
	conn *sql.DB
}

func NewTablesRepo(conn *sql.DB) *TablesRepo {
	return &TablesRepo{
		conn: conn,
	}
}

func (r TablesRepo) Create(ctx *codegen.BaseContext, data *codegen.Tables) error {
	panic("not implemented")
}

func (r TablesRepo) Update(ctx *codegen.BaseContext, data *codegen.Tables) error {
	panic("not implemented")
}

func (r TablesRepo) UpdatePartial(ctx *codegen.BaseContext, data *codegen.TablesPartial) error {
	panic("not implemented")
}

func (r TablesRepo) Delete(ctx *codegen.BaseContext, data *codegen.Tables) error {
	panic("not implemented")
}

func (r TablesRepo) Upsert(ctx *codegen.BaseContext, data []*codegen.Tables) error {
	panic("not implemented")
}

func (r TablesRepo) QueryBySchema(ctx *codegen.BaseContext, schema string) ([]*codegen.Tables, error) {
	return mysql.NewTablesStore(ctx, r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND table_type IN (UPPER('base table'), UPPER('system view'))").
		Query(schema)
}
