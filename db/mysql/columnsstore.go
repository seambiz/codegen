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

// Columns represents a row from 'COLUMNS'.
type Columns struct {
	codegen.Columns
}

// new implements Bindable.new
func (co *Columns) new() Bindable {
	return &Columns{}
}

// helper struct for common query operations.
type ColumnsSlice struct {
	data []*Columns
}

// append implements BindableSlice.append
func (co *ColumnsSlice) append(d Bindable) {
	co.data = append(co.data, d.(*Columns))
}

// constant slice for all fields of the table "Columns".
// nolint[gochecknoglobals]
var columnsQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "column_default", "is_nullable", "data_type", "character_maximum_length", "character_octet_length", "numeric_precision", "numeric_scale", "datetime_precision", "character_set_name", "collation_name", "column_type", "column_key", "extra", "privileges", "column_comment", "generation_expression"}

// returns fields, that should be used.
// nolint[gocyclo]
func ColumnsQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return columnsQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(codegen.Columns_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(codegen.Columns_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(codegen.Columns_ColumnName) == 1 {
		fields = append(fields, "column_name")
	}

	if colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		fields = append(fields, "ordinal_position")
	}

	if colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		fields = append(fields, "column_default")
	}

	if colSet.Bit(codegen.Columns_IsNullable) == 1 {
		fields = append(fields, "is_nullable")
	}

	if colSet.Bit(codegen.Columns_DataType) == 1 {
		fields = append(fields, "data_type")
	}

	if colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		fields = append(fields, "character_maximum_length")
	}

	if colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		fields = append(fields, "character_octet_length")
	}

	if colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		fields = append(fields, "numeric_precision")
	}

	if colSet.Bit(codegen.Columns_NumericScale) == 1 {
		fields = append(fields, "numeric_scale")
	}

	if colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		fields = append(fields, "datetime_precision")
	}

	if colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		fields = append(fields, "character_set_name")
	}

	if colSet.Bit(codegen.Columns_CollationName) == 1 {
		fields = append(fields, "collation_name")
	}

	if colSet.Bit(codegen.Columns_ColumnType) == 1 {
		fields = append(fields, "column_type")
	}

	if colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		fields = append(fields, "column_key")
	}

	if colSet.Bit(codegen.Columns_Extra) == 1 {
		fields = append(fields, "extra")
	}

	if colSet.Bit(codegen.Columns_Privileges) == 1 {
		fields = append(fields, "privileges")
	}

	if colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		fields = append(fields, "column_comment")
	}

	if colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		fields = append(fields, "generation_expression")
	}
	return fields
}

// ColumnsStore is used to query for 'Columns' records.
type ColumnsStore struct {
	Store
}

// NewColumnsStore return DAO Store for Columns
func NewColumnsStore(conn Execer) *ColumnsStore {
	co := &ColumnsStore{}
	co.db = conn
	co.withJoin = true
	co.joinType = sdb.LEFT
	co.batch = 1000
	return co
}

// WithoutJoins won't execute JOIN when querying for records.
func (co *ColumnsStore) WithoutJoins() *ColumnsStore {
	co.withJoin = false
	return co
}

// Where sets local sql, that will be appended to SELECT.
func (co *ColumnsStore) Where(sql string) *ColumnsStore {
	co.where = sql
	return co
}

// OrderBy sets local sql, that will be appended to SELECT.
func (co *ColumnsStore) OrderBy(sql string) *ColumnsStore {
	co.orderBy = sql
	return co
}

// GroupBy sets local sql, that will be appended to SELECT.
func (co *ColumnsStore) GroupBy(sql string) *ColumnsStore {
	co.groupBy = sql
	return co
}

// Limit result set size
func (co *ColumnsStore) Limit(n int) *ColumnsStore {
	co.limit = n
	return co
}

// Offset used, if a limit is provided
func (co *ColumnsStore) Offset(n int) *ColumnsStore {
	co.offset = n
	return co
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (co *ColumnsStore) JoinType(jt string) *ColumnsStore {
	co.joinType = jt
	return co
}

// Columns sets bits for specific columns.
func (co *ColumnsStore) Columns(cols ...int) *ColumnsStore {
	co.Store.Columns(cols...)
	return co
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (co *ColumnsStore) SetBits(colSet *big.Int) *ColumnsStore {
	co.colSet = colSet
	return co
}

// nolint[gocyclo]
func (co *Columns) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaColumns(&co.Columns, row, withJoin, colSet, col)
}

func BindInformationSchemaColumns(co *codegen.Columns, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		co.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_TableSchema) == 1 {
		co.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_TableName) == 1 {
		co.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnName) == 1 {
		co.ColumnName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		co.OrdinalPosition = sdb.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		if row[*col] == nil {
			co.ColumnDefault = nil
		} else {
			co.ColumnDefault = new(string)
			*co.ColumnDefault = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_IsNullable) == 1 {
		co.IsNullable = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_DataType) == 1 {
		co.DataType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		if row[*col] == nil {
			co.CharacterMaximumLength = nil
		} else {
			co.CharacterMaximumLength = new(uint64)
			*co.CharacterMaximumLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		if row[*col] == nil {
			co.CharacterOctetLength = nil
		} else {
			co.CharacterOctetLength = new(uint64)
			*co.CharacterOctetLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		if row[*col] == nil {
			co.NumericPrecision = nil
		} else {
			co.NumericPrecision = new(uint64)
			*co.NumericPrecision = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_NumericScale) == 1 {
		if row[*col] == nil {
			co.NumericScale = nil
		} else {
			co.NumericScale = new(uint64)
			*co.NumericScale = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		if row[*col] == nil {
			co.DatetimePrecision = nil
		} else {
			co.DatetimePrecision = new(uint64)
			*co.DatetimePrecision = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		if row[*col] == nil {
			co.CharacterSetName = nil
		} else {
			co.CharacterSetName = new(string)
			*co.CharacterSetName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CollationName) == 1 {
		if row[*col] == nil {
			co.CollationName = nil
		} else {
			co.CollationName = new(string)
			*co.CollationName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnType) == 1 {
		co.ColumnType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		co.ColumnKey = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_Extra) == 1 {
		co.Extra = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_Privileges) == 1 {
		co.Privileges = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		co.ColumnComment = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		co.GenerationExpression = sdb.ToString(row[*col])
		*col++
	}
}

func (co *ColumnsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", ColumnsQueryFields(co.colSet))
	sql.Append(" FROM information_schema.COLUMNS A ")
	if co.where != "" {
		sql.Append("WHERE", co.where)
	}
	if co.groupBy != "" {
		sql.Append("GROUP BY", co.groupBy)
	}
	if co.orderBy != "" {
		sql.Append("ORDER BY", co.orderBy)
	}
	if co.limit > 0 {
		sql.AppendRaw("LIMIT ", co.limit)
		if co.offset > 0 {
			sql.AppendRaw(",", co.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with 1:1 joined data.
func (co *ColumnsStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Columns, error) {
	dto := &Columns{}
	data := &ColumnsSlice{}
	err := co.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Columns, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Columns
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.COLUMNS' as a Columns with 1:1 joined data.
func (co *ColumnsStore) One(args ...interface{}) (*codegen.Columns, error) {
	data := &Columns{}

	err := co.one(data, co.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Columns, nil
}

// Query retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with 1:1 joined data.
func (co *ColumnsStore) Query(args ...interface{}) ([]*codegen.Columns, error) {
	stmt := co.selectStatement()
	return co.QueryCustom(stmt.Query(), args...)
}

// Insert inserts the Columns to the database.
func (co *ColumnsStore) Insert(data *codegen.Columns) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.COLUMNS (")
	fields := ColumnsQueryFields(co.colSet)
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
		log.Debug().Str("fn", "information_schema.COLUMNS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("ColumnName", data.ColumnName).Uint64("OrdinalPosition", data.OrdinalPosition).Str("ColumnDefault", logString(data.ColumnDefault)).Str("IsNullable", data.IsNullable).Str("DataType", data.DataType).Uint64("CharacterMaximumLength", logUInt64(data.CharacterMaximumLength)).Uint64("CharacterOctetLength", logUInt64(data.CharacterOctetLength)).Uint64("NumericPrecision", logUInt64(data.NumericPrecision)).Uint64("NumericScale", logUInt64(data.NumericScale)).Uint64("DatetimePrecision", logUInt64(data.DatetimePrecision)).Str("CharacterSetName", logString(data.CharacterSetName)).Str("CollationName", logString(data.CollationName)).Str("ColumnType", data.ColumnType).Str("ColumnKey", data.ColumnKey).Str("Extra", data.Extra).Str("Privileges", data.Privileges).Str("ColumnComment", data.ColumnComment).Str("GenerationExpression", data.GenerationExpression).Msg("sql")
	}
	_, err = co.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.ColumnName, data.OrdinalPosition, data.ColumnDefault, data.IsNullable, data.DataType, data.CharacterMaximumLength, data.CharacterOctetLength, data.NumericPrecision, data.NumericScale, data.DatetimePrecision, data.CharacterSetName, data.CollationName, data.ColumnType, data.ColumnKey, data.Extra, data.Privileges, data.ColumnComment, data.GenerationExpression)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Columns in the database.
// nolint[gocyclo]
func (co *ColumnsStore) Update(data *codegen.Columns) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.COLUMNS SET")
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		sql.AppendRaw(prepend, "ordinal_position = ?")
		prepend = ","
		args = append(args, data.OrdinalPosition)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		sql.AppendRaw(prepend, "column_default = ?")
		prepend = ","
		args = append(args, data.ColumnDefault)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_IsNullable) == 1 {
		sql.AppendRaw(prepend, "is_nullable = ?")
		prepend = ","
		args = append(args, data.IsNullable)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_DataType) == 1 {
		sql.AppendRaw(prepend, "data_type = ?")
		prepend = ","
		args = append(args, data.DataType)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		sql.AppendRaw(prepend, "character_maximum_length = ?")
		prepend = ","
		args = append(args, data.CharacterMaximumLength)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		sql.AppendRaw(prepend, "character_octet_length = ?")
		prepend = ","
		args = append(args, data.CharacterOctetLength)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		sql.AppendRaw(prepend, "numeric_precision = ?")
		prepend = ","
		args = append(args, data.NumericPrecision)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_NumericScale) == 1 {
		sql.AppendRaw(prepend, "numeric_scale = ?")
		prepend = ","
		args = append(args, data.NumericScale)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		sql.AppendRaw(prepend, "datetime_precision = ?")
		prepend = ","
		args = append(args, data.DatetimePrecision)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		sql.AppendRaw(prepend, "character_set_name = ?")
		prepend = ","
		args = append(args, data.CharacterSetName)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CollationName) == 1 {
		sql.AppendRaw(prepend, "collation_name = ?")
		prepend = ","
		args = append(args, data.CollationName)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnType) == 1 {
		sql.AppendRaw(prepend, "column_type = ?")
		prepend = ","
		args = append(args, data.ColumnType)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		sql.AppendRaw(prepend, "column_key = ?")
		prepend = ","
		args = append(args, data.ColumnKey)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_Extra) == 1 {
		sql.AppendRaw(prepend, "extra = ?")
		prepend = ","
		args = append(args, data.Extra)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_Privileges) == 1 {
		sql.AppendRaw(prepend, "privileges = ?")
		prepend = ","
		args = append(args, data.Privileges)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		sql.AppendRaw(prepend, "column_comment = ?")
		prepend = ","
		args = append(args, data.ColumnComment)
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		sql.AppendRaw(prepend, "generation_expression = ?")
		args = append(args, data.GenerationExpression)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.COLUMNS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := co.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Columns.
func (co *ColumnsStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.COLUMNS")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.COLUMNS.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := co.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (co *ColumnsStore) ToJSON(t *sdb.JsonBuffer, data *Columns) {
	prepend := "{"
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnName) == 1 {
		t.JS(prepend, "column_name", data.ColumnName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		t.JD64u(prepend, "ordinal_position", data.OrdinalPosition)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		t.JS(prepend, "column_default", *data.ColumnDefault)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_IsNullable) == 1 {
		t.JS(prepend, "is_nullable", data.IsNullable)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_DataType) == 1 {
		t.JS(prepend, "data_type", data.DataType)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		t.JD64u(prepend, "character_maximum_length", *data.CharacterMaximumLength)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		t.JD64u(prepend, "character_octet_length", *data.CharacterOctetLength)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		t.JD64u(prepend, "numeric_precision", *data.NumericPrecision)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_NumericScale) == 1 {
		t.JD64u(prepend, "numeric_scale", *data.NumericScale)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		t.JD64u(prepend, "datetime_precision", *data.DatetimePrecision)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		t.JS(prepend, "character_set_name", *data.CharacterSetName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_CollationName) == 1 {
		t.JS(prepend, "collation_name", *data.CollationName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnType) == 1 {
		t.JS(prepend, "column_type", data.ColumnType)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		t.JS(prepend, "column_key", data.ColumnKey)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_Extra) == 1 {
		t.JS(prepend, "extra", data.Extra)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_Privileges) == 1 {
		t.JS(prepend, "privileges", data.Privileges)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		t.JS(prepend, "column_comment", data.ColumnComment)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		t.JS(prepend, "generation_expression", data.GenerationExpression)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (co *ColumnsStore) ToJSONArray(w io.Writer, data []*Columns, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		co.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
