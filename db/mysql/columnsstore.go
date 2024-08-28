package mysql

import (
	"database/sql"
	"io"
	"math/big"

	codegen "github.com/seambiz/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// Columns represents a row from 'COLUMNS'.
type Columns struct {
	codegen.Columns
}

// new implements Bindable.new
func (s *Columns) new() Bindable {
	return &Columns{}
}

// helper struct for common query operations.
type ColumnsSlice struct {
	data []*Columns
}

// append implements BindableSlice.append
func (s *ColumnsSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Columns))
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
	ctx *codegen.Context
}

// NewColumnsStore return DAO Store for Columns
func NewColumnsStore(ctx *codegen.Context, conn Execer) *ColumnsStore {
	s := &ColumnsStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	s.ctx = ctx
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *ColumnsStore) WithoutJoins() *ColumnsStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *ColumnsStore) Where(sql string) *ColumnsStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *ColumnsStore) OrderBy(sql string) *ColumnsStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *ColumnsStore) GroupBy(sql string) *ColumnsStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *ColumnsStore) Limit(n int) *ColumnsStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *ColumnsStore) Offset(n int) *ColumnsStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *ColumnsStore) JoinType(jt string) *ColumnsStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *ColumnsStore) Columns(cols ...int) *ColumnsStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *ColumnsStore) SetBits(colSet *big.Int) *ColumnsStore {
	s.colSet = colSet
	return s
}

func (s *Columns) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		s.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_TableSchema) == 1 {
		s.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_TableName) == 1 {
		s.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnName) == 1 {
		if row[*col] == nil {
			s.ColumnName = nil
		} else {
			s.ColumnName = new(string)
			*s.ColumnName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		s.OrdinalPosition = sdb.ToUInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		if row[*col] == nil {
			s.ColumnDefault = nil
		} else {
			s.ColumnDefault = new(string)
			*s.ColumnDefault = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_IsNullable) == 1 {
		s.IsNullable = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_DataType) == 1 {
		if row[*col] == nil {
			s.DataType = nil
		} else {
			s.DataType = new(string)
			*s.DataType = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		if row[*col] == nil {
			s.CharacterMaximumLength = nil
		} else {
			s.CharacterMaximumLength = new(int64)
			*s.CharacterMaximumLength = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		if row[*col] == nil {
			s.CharacterOctetLength = nil
		} else {
			s.CharacterOctetLength = new(int64)
			*s.CharacterOctetLength = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		if row[*col] == nil {
			s.NumericPrecision = nil
		} else {
			s.NumericPrecision = new(uint64)
			*s.NumericPrecision = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_NumericScale) == 1 {
		if row[*col] == nil {
			s.NumericScale = nil
		} else {
			s.NumericScale = new(uint64)
			*s.NumericScale = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		if row[*col] == nil {
			s.DatetimePrecision = nil
		} else {
			s.DatetimePrecision = new(uint)
			*s.DatetimePrecision = sdb.ToUInt(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		if row[*col] == nil {
			s.CharacterSetName = nil
		} else {
			s.CharacterSetName = new(string)
			*s.CharacterSetName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_CollationName) == 1 {
		if row[*col] == nil {
			s.CollationName = nil
		} else {
			s.CollationName = new(string)
			*s.CollationName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnType) == 1 {
		s.ColumnType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		s.ColumnKey = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_Extra) == 1 {
		if row[*col] == nil {
			s.Extra = nil
		} else {
			s.Extra = new(string)
			*s.Extra = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_Privileges) == 1 {
		if row[*col] == nil {
			s.Privileges = nil
		} else {
			s.Privileges = new(string)
			*s.Privileges = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		s.ColumnComment = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		s.GenerationExpression = sdb.ToString(row[*col])
		*col++
	}
}

func (s *ColumnsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", ColumnsQueryFields(s.colSet))
	sql.Append(" FROM information_schema.COLUMNS A ")
	if s.where != "" {
		sql.Append("WHERE", s.where)
	}
	if s.groupBy != "" {
		sql.Append("GROUP BY", s.groupBy)
	}
	if s.orderBy != "" {
		sql.Append("ORDER BY", s.orderBy)
	}
	if s.limit > 0 {
		sql.AppendRaw("LIMIT ", s.limit)
		if s.offset > 0 {
			sql.AppendRaw(",", s.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with 1:1 joined data.
func (s *ColumnsStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Columns, error) {
	dto := &Columns{}
	data := &ColumnsSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Columns, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Columns
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.COLUMNS' as a Columns with 1:1 joined data.
func (s *ColumnsStore) One(args ...interface{}) (*codegen.Columns, error) {
	data := &Columns{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Columns, nil
}

// Query retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with 1:1 joined data.
func (s *ColumnsStore) Query(args ...interface{}) ([]*codegen.Columns, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// columnsUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *ColumnsStore) columnsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnName) == 1 {
		upsert = append(upsert, "column_name = VALUES(column_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		upsert = append(upsert, "ordinal_position = VALUES(ordinal_position)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		upsert = append(upsert, "column_default = VALUES(column_default)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_IsNullable) == 1 {
		upsert = append(upsert, "is_nullable = VALUES(is_nullable)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DataType) == 1 {
		upsert = append(upsert, "data_type = VALUES(data_type)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		upsert = append(upsert, "character_maximum_length = VALUES(character_maximum_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		upsert = append(upsert, "character_octet_length = VALUES(character_octet_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		upsert = append(upsert, "numeric_precision = VALUES(numeric_precision)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericScale) == 1 {
		upsert = append(upsert, "numeric_scale = VALUES(numeric_scale)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		upsert = append(upsert, "datetime_precision = VALUES(datetime_precision)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		upsert = append(upsert, "character_set_name = VALUES(character_set_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CollationName) == 1 {
		upsert = append(upsert, "collation_name = VALUES(collation_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnType) == 1 {
		upsert = append(upsert, "column_type = VALUES(column_type)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		upsert = append(upsert, "column_key = VALUES(column_key)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Extra) == 1 {
		upsert = append(upsert, "extra = VALUES(extra)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Privileges) == 1 {
		upsert = append(upsert, "privileges = VALUES(privileges)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		upsert = append(upsert, "column_comment = VALUES(column_comment)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		upsert = append(upsert, "generation_expression = VALUES(generation_expression)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.COLUMNS")
	sql.Columns("table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "column_default", "is_nullable", "data_type", "character_maximum_length", "character_octet_length", "numeric_precision", "numeric_scale", "datetime_precision", "character_set_name", "collation_name", "column_type", "column_key", "extra", "privileges", "column_comment", "generation_expression")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Columns
func (s *ColumnsStore) Upsert(data ...*codegen.Columns) (int64, error) {
	sql := s.columnsUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "ColumnsUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return -1, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		s.log.Error().Err(err).Msg("rowsaffected")
		return -1, err
	}

	return affected, nil
}

// Insert inserts the Columns to the database.
func (s *ColumnsStore) Insert(data *codegen.Columns) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.COLUMNS (")
	fields := ColumnsQueryFields(s.colSet)
	sql.Fields("", "", fields)
	sql.Append(") VALUES (")
	for i := range fields {
		if i > 0 {
			sql.Append(",")
		}
		sql.Append("?")
	}
	sql.Append(")")

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.COLUMNS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("ColumnName", logString(data.ColumnName)).Uint("OrdinalPosition", data.OrdinalPosition).Str("ColumnDefault", logString(data.ColumnDefault)).Str("IsNullable", data.IsNullable).Str("DataType", logString(data.DataType)).Int64("CharacterMaximumLength", logInt64(data.CharacterMaximumLength)).Int64("CharacterOctetLength", logInt64(data.CharacterOctetLength)).Uint64("NumericPrecision", logUInt64(data.NumericPrecision)).Uint64("NumericScale", logUInt64(data.NumericScale)).Uint("DatetimePrecision", logUInt(data.DatetimePrecision)).Str("CharacterSetName", logString(data.CharacterSetName)).Str("CollationName", logString(data.CollationName)).Str("ColumnType", data.ColumnType).Str("ColumnKey", data.ColumnKey).Str("Extra", logString(data.Extra)).Str("Privileges", logString(data.Privileges)).Str("ColumnComment", data.ColumnComment).Str("GenerationExpression", data.GenerationExpression).Msg("sql")
	}
	_, err = s.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.ColumnName, data.OrdinalPosition, data.ColumnDefault, data.IsNullable, data.DataType, data.CharacterMaximumLength, data.CharacterOctetLength, data.NumericPrecision, data.NumericScale, data.DatetimePrecision, data.CharacterSetName, data.CollationName, data.ColumnType, data.ColumnKey, data.Extra, data.Privileges, data.ColumnComment, data.GenerationExpression)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Columns in the database.
// nolint[gocyclo]
func (s *ColumnsStore) Update(data *codegen.Columns) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.COLUMNS SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		sql.AppendRaw(prepend, "ordinal_position = ?")
		prepend = ","
		args = append(args, data.OrdinalPosition)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		sql.AppendRaw(prepend, "column_default = ?")
		prepend = ","
		args = append(args, data.ColumnDefault)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_IsNullable) == 1 {
		sql.AppendRaw(prepend, "is_nullable = ?")
		prepend = ","
		args = append(args, data.IsNullable)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DataType) == 1 {
		sql.AppendRaw(prepend, "data_type = ?")
		prepend = ","
		args = append(args, data.DataType)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		sql.AppendRaw(prepend, "character_maximum_length = ?")
		prepend = ","
		args = append(args, data.CharacterMaximumLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		sql.AppendRaw(prepend, "character_octet_length = ?")
		prepend = ","
		args = append(args, data.CharacterOctetLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		sql.AppendRaw(prepend, "numeric_precision = ?")
		prepend = ","
		args = append(args, data.NumericPrecision)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericScale) == 1 {
		sql.AppendRaw(prepend, "numeric_scale = ?")
		prepend = ","
		args = append(args, data.NumericScale)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		sql.AppendRaw(prepend, "datetime_precision = ?")
		prepend = ","
		args = append(args, data.DatetimePrecision)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		sql.AppendRaw(prepend, "character_set_name = ?")
		prepend = ","
		args = append(args, data.CharacterSetName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CollationName) == 1 {
		sql.AppendRaw(prepend, "collation_name = ?")
		prepend = ","
		args = append(args, data.CollationName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnType) == 1 {
		sql.AppendRaw(prepend, "column_type = ?")
		prepend = ","
		args = append(args, data.ColumnType)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		sql.AppendRaw(prepend, "column_key = ?")
		prepend = ","
		args = append(args, data.ColumnKey)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Extra) == 1 {
		sql.AppendRaw(prepend, "extra = ?")
		prepend = ","
		args = append(args, data.Extra)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Privileges) == 1 {
		sql.AppendRaw(prepend, "privileges = ?")
		prepend = ","
		args = append(args, data.Privileges)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		sql.AppendRaw(prepend, "column_comment = ?")
		prepend = ","
		args = append(args, data.ColumnComment)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		sql.AppendRaw(prepend, "generation_expression = ?")
		args = append(args, data.GenerationExpression)
	}
	sql.Append(" WHERE ")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.COLUMNS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Columns.
func (s *ColumnsStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.COLUMNS")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.COLUMNS.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *ColumnsStore) ToJSON(t *sdb.JsonBuffer, data *Columns) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnName) == 1 {
		t.JS(prepend, "column_name", *data.ColumnName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_OrdinalPosition) == 1 {
		t.JDu(prepend, "ordinal_position", data.OrdinalPosition)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnDefault) == 1 {
		t.JS(prepend, "column_default", *data.ColumnDefault)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_IsNullable) == 1 {
		t.JS(prepend, "is_nullable", data.IsNullable)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DataType) == 1 {
		t.JS(prepend, "data_type", *data.DataType)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterMaximumLength) == 1 {
		t.JD64(prepend, "character_maximum_length", *data.CharacterMaximumLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterOctetLength) == 1 {
		t.JD64(prepend, "character_octet_length", *data.CharacterOctetLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericPrecision) == 1 {
		t.JD64u(prepend, "numeric_precision", *data.NumericPrecision)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_NumericScale) == 1 {
		t.JD64u(prepend, "numeric_scale", *data.NumericScale)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_DatetimePrecision) == 1 {
		t.JDu(prepend, "datetime_precision", *data.DatetimePrecision)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CharacterSetName) == 1 {
		t.JS(prepend, "character_set_name", *data.CharacterSetName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_CollationName) == 1 {
		t.JS(prepend, "collation_name", *data.CollationName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnType) == 1 {
		t.JS(prepend, "column_type", data.ColumnType)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnKey) == 1 {
		t.JS(prepend, "column_key", data.ColumnKey)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Extra) == 1 {
		t.JS(prepend, "extra", *data.Extra)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_Privileges) == 1 {
		t.JS(prepend, "privileges", *data.Privileges)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_ColumnComment) == 1 {
		t.JS(prepend, "column_comment", data.ColumnComment)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Columns_GenerationExpression) == 1 {
		t.JS(prepend, "generation_expression", data.GenerationExpression)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *ColumnsStore) ToJSONArray(w io.Writer, data []*Columns, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		s.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. ^^
