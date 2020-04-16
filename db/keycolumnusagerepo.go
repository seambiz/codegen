package db

import (
	"database/sql"

	"bitbucket.org/codegen/db/mysql"

	"bitbucket.org/codegen"
)

type KeyColUsageRepo struct {
	conn *sql.DB
}

func NewKeyColUsageRepo(conn *sql.DB) *KeyColUsageRepo {
	return &KeyColUsageRepo{
		conn: conn,
	}
}

func (r KeyColUsageRepo) Create(data *codegen.KeyColumnUsage) error {
	panic("not implemented") // TODO: Implement
}

func (r KeyColUsageRepo) Update(data *codegen.KeyColumnUsage) error {
	panic("not implemented") // TODO: Implement
}

func (r KeyColUsageRepo) Delete(data *codegen.KeyColumnUsage) error {
	panic("not implemented") // TODO: Implement
}

func (r KeyColUsageRepo) QueryBySchemaAndRefSchemaAndTable(schema, refschema, table string) ([]*codegen.KeyColumnUsage, error) {
	return mysql.NewKeyColumnUsageStore(r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND UPPER(referenced_table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?)").
		OrderBy("constraint_name, ordinal_position").
		Query(schema, refschema, table)
}
