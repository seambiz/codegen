package db

import (
	"database/sql"

	"github.com/seambiz/codegen"
	"github.com/seambiz/codegen/db/mysql"
)

// GENERATED BY CODEGEN.

/* COLUMNSRepo implements COLUMNSRepository interface definition. */
type COLUMNSRepo struct {
	conn *sql.DB
}

func NewCOLUMNSRepo(conn *sql.DB) codegen.COLUMNSRepository {
	return &COLUMNSRepo{
		conn: conn,
	}
}
	

	
		
	

	


// ^^ END OF GENERATED BY CODEGEN. ^^

func (r ColumnsRepo) QueryBySchemaAndTable(ctx *codegen.Context, schema, table string) ([]*codegen.Columns, error) {
	return mysql.NewColumnsStore(ctx, r.conn).
		Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?) AND UPPER(extra) not like '%VIRTUAL%'").
		OrderBy("ordinal_position").
		Query(schema, table)
}
