package db

import (
	"database/sql"

	"bitbucket.org/codegen/db/mysql"

	"bitbucket.org/codegen"
)

type StatisticsRepo struct {
	conn *sql.DB
}

func NewStatisticsRepo(conn *sql.DB) *StatisticsRepo {
	return &StatisticsRepo{
		conn: conn,
	}
}

func (r StatisticsRepo) Create(data *codegen.Statistics) error {
	panic("not implemented") // TODO: Implement
}

func (r StatisticsRepo) Update(data *codegen.Statistics) error {
	panic("not implemented") // TODO: Implement
}

func (r StatisticsRepo) Delete(data *codegen.Statistics) error {
	panic("not implemented") // TODO: Implement
}

func (r StatisticsRepo) IndexNameBySchemaAndTable(schema, table string) ([]*codegen.Statistics, error) {
	store := mysql.NewStatisticsStore(r.conn)
	store.Columns(codegen.Statistics_IndexName)
	return store.Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?)").
		GroupBy("index_name").
		Query(schema, table)
}

func (r StatisticsRepo) QueryBySchemaAndTableAndIndex(schema, table, index string) ([]*codegen.Statistics, error) {
	return mysql.NewStatisticsStore(r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?) AND UPPER(index_name) = UPPER(?)").
		OrderBy("seq_in_index").
		Query(schema, table, index)
}
