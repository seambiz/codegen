package db

import (
	"database/sql"

	"bitbucket.org/codegen"
	"bitbucket.org/codegen/db/mysql"
)

// GENERATED BY CODEGEN. DO NOT EDIT.
// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^

/* StatisticsRepo implemente StatisticsRepository interface definition. */
type StatisticsRepo struct {
	conn *sql.DB
	ctx  *codegen.BaseContext
}

func NewStatisticsRepo(ctx *codegen.BaseContext, conn *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{
		conn: conn,
		ctx:  ctx,
	}
}

func (r StatisticsRepo) Create(data *codegen.Statistics) error {
	panic("not implemented")
}

func (r StatisticsRepo) Update(data *codegen.Statistics) error {
	panic("not implemented")
}

func (r StatisticsRepo) UpdatePartial(data *codegen.StatisticsPartial) error {
	panic("not implemented")
}

func (r StatisticsRepo) Delete(data *codegen.Statistics) error {
	panic("not implemented")
}

func (r StatisticsRepo) Upsert(data []*codegen.Statistics) error {
	panic("not implemented")
}

func (r StatisticsRepo) IndexNameBySchemaAndTable(schema, table string) ([]*codegen.Statistics, error) {
	store := mysql.NewStatisticsStore(r.ctx, r.conn)
	store.Columns(codegen.Statistics_IndexName)
	return store.Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?)").
		GroupBy("index_name").
		Query(schema, table)
}

func (r StatisticsRepo) QueryBySchemaAndTableAndIndex(schema, table, index string) ([]*codegen.Statistics, error) {
	return mysql.NewStatisticsStore(r.ctx, r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?) AND UPPER(index_name) = UPPER(?)").
		OrderBy("seq_in_index").
		Query(schema, table, index)
}
