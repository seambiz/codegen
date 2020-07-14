package mysql

import (
	"database/sql"
	"io"
	"math/big"

	codegen "bitbucket.org/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

// KeyColumnUsage represents a row from 'KEY_COLUMN_USAGE'.
type KeyColumnUsage struct {
	codegen.KeyColumnUsage
}

// new implements Bindable.new
func (ke *KeyColumnUsage) new() Bindable {
	return &KeyColumnUsage{}
}

// helper struct for common query operations.
type KeyColumnUsageSlice struct {
	data []*KeyColumnUsage
}

// append implements BindableSlice.append
func (ke *KeyColumnUsageSlice) append(d Bindable) {
	ke.data = append(ke.data, d.(*KeyColumnUsage))
}

// constant slice for all fields of the table "KeyColumnUsage".
// nolint[gochecknoglobals]
var keycolumnusageQueryFieldsAll = []string{"constraint_catalog", "constraint_schema", "constraint_name", "table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "position_in_unique_constraint", "referenced_table_schema", "referenced_table_name", "referenced_column_name"}

// returns fields, that should be used.
// nolint[gocyclo]
func KeyColumnUsageQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return keycolumnusageQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		fields = append(fields, "constraint_catalog")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		fields = append(fields, "constraint_schema")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		fields = append(fields, "constraint_name")
	}

	if colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		fields = append(fields, "column_name")
	}

	if colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		fields = append(fields, "ordinal_position")
	}

	if colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		fields = append(fields, "position_in_unique_constraint")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		fields = append(fields, "referenced_table_schema")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		fields = append(fields, "referenced_table_name")
	}

	if colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		fields = append(fields, "referenced_column_name")
	}
	return fields
}

// KeyColumnUsageStore is used to query for 'KeyColumnUsage' records.
type KeyColumnUsageStore struct {
	Store
}

// NewKeyColumnUsageStore return DAO Store for KeyColumnUsage
func NewKeyColumnUsageStore(conn Execer) *KeyColumnUsageStore {
	ke := &KeyColumnUsageStore{}
	ke.db = conn
	ke.withJoin = true
	ke.joinType = sdb.LEFT
	ke.batch = 1000
	return ke
}

// WithoutJoins won't execute JOIN when querying for records.
func (ke *KeyColumnUsageStore) WithoutJoins() *KeyColumnUsageStore {
	ke.withJoin = false
	return ke
}

// Where sets local sql, that will be appended to SELECT.
func (ke *KeyColumnUsageStore) Where(sql string) *KeyColumnUsageStore {
	ke.where = sql
	return ke
}

// OrderBy sets local sql, that will be appended to SELECT.
func (ke *KeyColumnUsageStore) OrderBy(sql string) *KeyColumnUsageStore {
	ke.orderBy = sql
	return ke
}

// GroupBy sets local sql, that will be appended to SELECT.
func (ke *KeyColumnUsageStore) GroupBy(sql string) *KeyColumnUsageStore {
	ke.groupBy = sql
	return ke
}

// Limit result set size
func (ke *KeyColumnUsageStore) Limit(n int) *KeyColumnUsageStore {
	ke.limit = n
	return ke
}

// Offset used, if a limit is provided
func (ke *KeyColumnUsageStore) Offset(n int) *KeyColumnUsageStore {
	ke.offset = n
	return ke
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (ke *KeyColumnUsageStore) JoinType(jt string) *KeyColumnUsageStore {
	ke.joinType = jt
	return ke
}

// Columns sets bits for specific columns.
func (ke *KeyColumnUsageStore) Columns(cols ...int) *KeyColumnUsageStore {
	ke.Store.Columns(cols...)
	return ke
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (ke *KeyColumnUsageStore) SetBits(colSet *big.Int) *KeyColumnUsageStore {
	ke.colSet = colSet
	return ke
}

func (ke *KeyColumnUsage) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaKeyColumnUsage(&ke.KeyColumnUsage, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindInformationSchemaKeyColumnUsage(ke *codegen.KeyColumnUsage, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		ke.ConstraintCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		ke.ConstraintSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		ke.ConstraintName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		ke.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		ke.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		ke.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		ke.ColumnName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		ke.OrdinalPosition = sdb.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		if row[*col] == nil {
			ke.PositionInUniqueConstraint = nil
		} else {
			ke.PositionInUniqueConstraint = new(int64)
			*ke.PositionInUniqueConstraint = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		if row[*col] == nil {
			ke.ReferencedTableSchema = nil
		} else {
			ke.ReferencedTableSchema = new(string)
			*ke.ReferencedTableSchema = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		if row[*col] == nil {
			ke.ReferencedTableName = nil
		} else {
			ke.ReferencedTableName = new(string)
			*ke.ReferencedTableName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		if row[*col] == nil {
			ke.ReferencedColumnName = nil
		} else {
			ke.ReferencedColumnName = new(string)
			*ke.ReferencedColumnName = sdb.ToString(row[*col])
		}
		*col++
	}
}

func (ke *KeyColumnUsageStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", KeyColumnUsageQueryFields(ke.colSet))
	sql.Append(" FROM information_schema.KEY_COLUMN_USAGE A ")
	if ke.where != "" {
		sql.Append("WHERE", ke.where)
	}
	if ke.groupBy != "" {
		sql.Append("GROUP BY", ke.groupBy)
	}
	if ke.orderBy != "" {
		sql.Append("ORDER BY", ke.orderBy)
	}
	if ke.limit > 0 {
		sql.AppendRaw("LIMIT ", ke.limit)
		if ke.offset > 0 {
			sql.AppendRaw(",", ke.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'information_schema.KEY_COLUMN_USAGE' as a slice of KeyColumnUsage with 1:1 joined data.
func (ke *KeyColumnUsageStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.KeyColumnUsage, error) {
	dto := &KeyColumnUsage{}
	data := &KeyColumnUsageSlice{}
	err := ke.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.KeyColumnUsage, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].KeyColumnUsage
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.KEY_COLUMN_USAGE' as a KeyColumnUsage with 1:1 joined data.
func (ke *KeyColumnUsageStore) One(args ...interface{}) (*codegen.KeyColumnUsage, error) {
	data := &KeyColumnUsage{}

	err := ke.one(data, ke.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.KeyColumnUsage, nil
}

// Query retrieves many rows from 'information_schema.KEY_COLUMN_USAGE' as a slice of KeyColumnUsage with 1:1 joined data.
func (ke *KeyColumnUsageStore) Query(args ...interface{}) ([]*codegen.KeyColumnUsage, error) {
	stmt := ke.selectStatement()
	return ke.QueryCustom(stmt.Query(), args...)
}

// keyColumnUsageUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (ke *KeyColumnUsageStore) keyColumnUsageUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		upsert = append(upsert, "constraint_catalog = VALUES(constraint_catalog)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		upsert = append(upsert, "constraint_schema = VALUES(constraint_schema)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		upsert = append(upsert, "constraint_name = VALUES(constraint_name)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		upsert = append(upsert, "column_name = VALUES(column_name)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		upsert = append(upsert, "ordinal_position = VALUES(ordinal_position)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		upsert = append(upsert, "position_in_unique_constraint = VALUES(position_in_unique_constraint)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		upsert = append(upsert, "referenced_table_schema = VALUES(referenced_table_schema)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		upsert = append(upsert, "referenced_table_name = VALUES(referenced_table_name)")
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		upsert = append(upsert, "referenced_column_name = VALUES(referenced_column_name)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.KEY_COLUMN_USAGE")
	sql.Columns("constraint_catalog", "constraint_schema", "constraint_name", "table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "position_in_unique_constraint", "referenced_table_schema", "referenced_table_name", "referenced_column_name")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of KeyColumnUsage
func (ke *KeyColumnUsageStore) Upsert(data ...*codegen.KeyColumnUsage) (int64, error) {
	sql := ke.keyColumnUsageUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "KeyColumnUsageUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := ke.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return -1, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("rowsaffected")
		return -1, err
	}

	return affected, nil
}

// Insert inserts the KeyColumnUsage to the database.
func (ke *KeyColumnUsageStore) Insert(data *codegen.KeyColumnUsage) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.KEY_COLUMN_USAGE (")
	fields := KeyColumnUsageQueryFields(ke.colSet)
	sql.Fields("", "", fields)
	sql.Append(") VALUES (")
	for i := range fields {
		if i > 0 {
			sql.Append(",")
		}
		sql.Append("?")
	}
	sql.Append(")")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.KEY_COLUMN_USAGE.Insert").Str("stmt", sql.String()).Str("ConstraintCatalog", data.ConstraintCatalog).Str("ConstraintSchema", data.ConstraintSchema).Str("ConstraintName", data.ConstraintName).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("ColumnName", data.ColumnName).Int64("OrdinalPosition", data.OrdinalPosition).Int64("PositionInUniqueConstraint", logInt64(data.PositionInUniqueConstraint)).Str("ReferencedTableSchema", logString(data.ReferencedTableSchema)).Str("ReferencedTableName", logString(data.ReferencedTableName)).Str("ReferencedColumnName", logString(data.ReferencedColumnName)).Msg("sql")
	}
	_, err = ke.db.Exec(sql.Query(), data.ConstraintCatalog, data.ConstraintSchema, data.ConstraintName, data.TableCatalog, data.TableSchema, data.TableName, data.ColumnName, data.OrdinalPosition, data.PositionInUniqueConstraint, data.ReferencedTableSchema, data.ReferencedTableName, data.ReferencedColumnName)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the KeyColumnUsage in the database.
// nolint[gocyclo]
func (ke *KeyColumnUsageStore) Update(data *codegen.KeyColumnUsage) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.KEY_COLUMN_USAGE SET")
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		sql.AppendRaw(prepend, "constraint_catalog = ?")
		prepend = ","
		args = append(args, data.ConstraintCatalog)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		sql.AppendRaw(prepend, "constraint_schema = ?")
		prepend = ","
		args = append(args, data.ConstraintSchema)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		sql.AppendRaw(prepend, "constraint_name = ?")
		prepend = ","
		args = append(args, data.ConstraintName)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		sql.AppendRaw(prepend, "ordinal_position = ?")
		prepend = ","
		args = append(args, data.OrdinalPosition)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		sql.AppendRaw(prepend, "position_in_unique_constraint = ?")
		prepend = ","
		args = append(args, data.PositionInUniqueConstraint)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		sql.AppendRaw(prepend, "referenced_table_schema = ?")
		prepend = ","
		args = append(args, data.ReferencedTableSchema)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		sql.AppendRaw(prepend, "referenced_table_name = ?")
		prepend = ","
		args = append(args, data.ReferencedTableName)
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		sql.AppendRaw(prepend, "referenced_column_name = ?")
		args = append(args, data.ReferencedColumnName)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.KEY_COLUMN_USAGE.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := ke.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from KeyColumnUsage.
func (ke *KeyColumnUsageStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.KEY_COLUMN_USAGE")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.KEY_COLUMN_USAGE.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := ke.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (ke *KeyColumnUsageStore) ToJSON(t *sdb.JsonBuffer, data *KeyColumnUsage) {
	prepend := "{"
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		t.JS(prepend, "constraint_catalog", data.ConstraintCatalog)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		t.JS(prepend, "constraint_schema", data.ConstraintSchema)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		t.JS(prepend, "constraint_name", data.ConstraintName)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		t.JS(prepend, "column_name", data.ColumnName)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		t.JD64(prepend, "ordinal_position", data.OrdinalPosition)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		t.JD64(prepend, "position_in_unique_constraint", *data.PositionInUniqueConstraint)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		t.JS(prepend, "referenced_table_schema", *data.ReferencedTableSchema)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		t.JS(prepend, "referenced_table_name", *data.ReferencedTableName)
		prepend = ","
	}
	if ke.colSet == nil || ke.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		t.JS(prepend, "referenced_column_name", *data.ReferencedColumnName)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (ke *KeyColumnUsageStore) ToJSONArray(w io.Writer, data []*KeyColumnUsage, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		ke.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
