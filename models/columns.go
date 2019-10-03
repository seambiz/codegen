// Package db contains the types for schema 'information_schema'.
package models

import (
	"database/sql"
	"io"
	"math/big"

	"bitbucket.org/codegen/convert"

	"bitbucket.org/seambiz/buffer"
	"bitbucket.org/seambiz/sdb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GENERATED BY CODEGEN. DO NOT EDIT.
// constant slice for all fields of the table.
// nolint[gochecknoglobals]
var columnsQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "column_default", "is_nullable", "data_type", "character_maximum_length", "character_octet_length", "numeric_precision", "numeric_scale", "datetime_precision", "character_set_name", "collation_name", "column_type", "column_key", "extra", "privileges", "column_comment", "generation_expression"}

// returns fields, that should be used.
// nolint[gocyclo]
func ColumnsQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return columnsQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(ColumnsTableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}
	if colSet.Bit(ColumnsTableSchema) == 1 {
		fields = append(fields, "table_schema")
	}
	if colSet.Bit(ColumnsTableName) == 1 {
		fields = append(fields, "table_name")
	}
	if colSet.Bit(ColumnsColumnName) == 1 {
		fields = append(fields, "column_name")
	}
	if colSet.Bit(ColumnsOrdinalPosition) == 1 {
		fields = append(fields, "ordinal_position")
	}
	if colSet.Bit(ColumnsColumnDefault) == 1 {
		fields = append(fields, "column_default")
	}
	if colSet.Bit(ColumnsIsNullable) == 1 {
		fields = append(fields, "is_nullable")
	}
	if colSet.Bit(ColumnsDataType) == 1 {
		fields = append(fields, "data_type")
	}
	if colSet.Bit(ColumnsCharacterMaximumLength) == 1 {
		fields = append(fields, "character_maximum_length")
	}
	if colSet.Bit(ColumnsCharacterOctetLength) == 1 {
		fields = append(fields, "character_octet_length")
	}
	if colSet.Bit(ColumnsNumericPrecision) == 1 {
		fields = append(fields, "numeric_precision")
	}
	if colSet.Bit(ColumnsNumericScale) == 1 {
		fields = append(fields, "numeric_scale")
	}
	if colSet.Bit(ColumnsDatetimePrecision) == 1 {
		fields = append(fields, "datetime_precision")
	}
	if colSet.Bit(ColumnsCharacterSetName) == 1 {
		fields = append(fields, "character_set_name")
	}
	if colSet.Bit(ColumnsCollationName) == 1 {
		fields = append(fields, "collation_name")
	}
	if colSet.Bit(ColumnsColumnType) == 1 {
		fields = append(fields, "column_type")
	}
	if colSet.Bit(ColumnsColumnKey) == 1 {
		fields = append(fields, "column_key")
	}
	if colSet.Bit(ColumnsExtra) == 1 {
		fields = append(fields, "extra")
	}
	if colSet.Bit(ColumnsPrivileges) == 1 {
		fields = append(fields, "privileges")
	}
	if colSet.Bit(ColumnsColumnComment) == 1 {
		fields = append(fields, "column_comment")
	}
	if colSet.Bit(ColumnsGenerationExpression) == 1 {
		fields = append(fields, "generation_expression")
	}
	return fields
}

// Columns represents a row from 'information_schema.COLUMNS'.
type Columns struct {
	TableCatalog           string `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema            string `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName              string `json:"TABLE_NAME" db:"table_name"`
	ColumnName             string `json:"COLUMN_NAME" db:"column_name"`
	OrdinalPosition        uint64 `json:"ORDINAL_POSITION" db:"ordinal_position"`
	ColumnDefault          string `json:"COLUMN_DEFAULT" db:"column_default"`
	IsNullable             string `json:"IS_NULLABLE" db:"is_nullable"`
	DataType               string `json:"DATA_TYPE" db:"data_type"`
	CharacterMaximumLength uint64 `json:"CHARACTER_MAXIMUM_LENGTH" db:"character_maximum_length"`
	CharacterOctetLength   uint64 `json:"CHARACTER_OCTET_LENGTH" db:"character_octet_length"`
	NumericPrecision       uint64 `json:"NUMERIC_PRECISION" db:"numeric_precision"`
	NumericScale           uint64 `json:"NUMERIC_SCALE" db:"numeric_scale"`
	DatetimePrecision      uint64 `json:"DATETIME_PRECISION" db:"datetime_precision"`
	CharacterSetName       string `json:"CHARACTER_SET_NAME" db:"character_set_name"`
	CollationName          string `json:"COLLATION_NAME" db:"collation_name"`
	ColumnType             string `json:"COLUMN_TYPE" db:"column_type"`
	ColumnKey              string `json:"COLUMN_KEY" db:"column_key"`
	Extra                  string `json:"EXTRA" db:"extra"`
	Privileges             string `json:"PRIVILEGES" db:"privileges"`
	ColumnComment          string `json:"COLUMN_COMMENT" db:"column_comment"`
	GenerationExpression   string `json:"GENERATION_EXPRESSION" db:"generation_expression"`
}

// new implements DTO.new
func (co *Columns) new() DTO {
	return &Columns{}
}

// helper struct for common query operations.
type ColumnsSlice struct {
	data []*Columns
}

// append implements DTOSlice.append
func (co *ColumnsSlice) append(d DTO) {
	co.data = append(co.data, d.(*Columns))
}

// Columns to be used for various statements.
func (co *ColumnsStore) Columns(cols ...int) *ColumnsStore {
	co.colSet = big.NewInt(0)
	for _, col := range cols {
		co.colSet.SetBit(co.colSet, col, 1)
	}
	return co
}

// IsEmpty checks if primary key fields are zero.
func (co *Columns) IsEmpty() bool {
	return true
}

// ColumnsStore is used to query for 'Columns' records.
type ColumnsStore struct {
	Store
}

// NewColumnsStore return DAO Store for Columns
func NewColumnsStore(conn *sql.DB) *ColumnsStore {
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

// nolint[gocyclo]
func (co *Columns) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(ColumnsTableCatalog) == 1 {
		co.TableCatalog = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsTableSchema) == 1 {
		co.TableSchema = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsTableName) == 1 {
		co.TableName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsColumnName) == 1 {
		co.ColumnName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsOrdinalPosition) == 1 {
		co.OrdinalPosition = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsColumnDefault) == 1 {
		co.ColumnDefault = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsIsNullable) == 1 {
		co.IsNullable = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsDataType) == 1 {
		co.DataType = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsCharacterMaximumLength) == 1 {
		co.CharacterMaximumLength = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsCharacterOctetLength) == 1 {
		co.CharacterOctetLength = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsNumericPrecision) == 1 {
		co.NumericPrecision = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsNumericScale) == 1 {
		co.NumericScale = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsDatetimePrecision) == 1 {
		co.DatetimePrecision = convert.ToUInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsCharacterSetName) == 1 {
		co.CharacterSetName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsCollationName) == 1 {
		co.CollationName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsColumnType) == 1 {
		co.ColumnType = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsColumnKey) == 1 {
		co.ColumnKey = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsExtra) == 1 {
		co.Extra = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsPrivileges) == 1 {
		co.Privileges = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsColumnComment) == 1 {
		co.ColumnComment = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(ColumnsGenerationExpression) == 1 {
		co.GenerationExpression = convert.ToString(row[*col])
		*col++
	}
}
func (co *ColumnsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", ColumnsQueryFields(co.colSet))
	sql.Append("FROM information_schema.COLUMNS A")
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

// One retrieves a row from 'information_schema.COLUMNS' as a Columns with possible joined data.
func (co *ColumnsStore) One(args ...interface{}) (*Columns, error) {
	data := &Columns{}

	err := co.one(data, co.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return data, nil
}

// Query retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with possible joined data.
func (co *ColumnsStore) Query(args ...interface{}) ([]*Columns, error) {
	stmt := co.selectStatement()
	return co.QueryCustom(stmt.Query(), args...)
}

// QueryCustom retrieves many rows from 'information_schema.COLUMNS' as a slice of Columns with possible joined data.
func (co *ColumnsStore) QueryCustom(stmt string, args ...interface{}) ([]*Columns, error) {
	dto := &Columns{}
	data := &ColumnsSlice{}
	err := co.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	return data.data, nil
}

// columnsUpsertStmt helper for generating Upserts general statement
// nolint[gocyclo]
func (co *ColumnsStore) columnsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableCatalog) == 1 {
		upsert = append(upsert, "TABLE_CATALOG = VALUES(TABLE_CATALOG)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableSchema) == 1 {
		upsert = append(upsert, "TABLE_SCHEMA = VALUES(TABLE_SCHEMA)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableName) == 1 {
		upsert = append(upsert, "TABLE_NAME = VALUES(TABLE_NAME)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnName) == 1 {
		upsert = append(upsert, "COLUMN_NAME = VALUES(COLUMN_NAME)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsOrdinalPosition) == 1 {
		upsert = append(upsert, "ORDINAL_POSITION = VALUES(ORDINAL_POSITION)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnDefault) == 1 {
		upsert = append(upsert, "COLUMN_DEFAULT = VALUES(COLUMN_DEFAULT)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsIsNullable) == 1 {
		upsert = append(upsert, "IS_NULLABLE = VALUES(IS_NULLABLE)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDataType) == 1 {
		upsert = append(upsert, "DATA_TYPE = VALUES(DATA_TYPE)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterMaximumLength) == 1 {
		upsert = append(upsert, "CHARACTER_MAXIMUM_LENGTH = VALUES(CHARACTER_MAXIMUM_LENGTH)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterOctetLength) == 1 {
		upsert = append(upsert, "CHARACTER_OCTET_LENGTH = VALUES(CHARACTER_OCTET_LENGTH)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericPrecision) == 1 {
		upsert = append(upsert, "NUMERIC_PRECISION = VALUES(NUMERIC_PRECISION)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericScale) == 1 {
		upsert = append(upsert, "NUMERIC_SCALE = VALUES(NUMERIC_SCALE)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDatetimePrecision) == 1 {
		upsert = append(upsert, "DATETIME_PRECISION = VALUES(DATETIME_PRECISION)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterSetName) == 1 {
		upsert = append(upsert, "CHARACTER_SET_NAME = VALUES(CHARACTER_SET_NAME)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCollationName) == 1 {
		upsert = append(upsert, "COLLATION_NAME = VALUES(COLLATION_NAME)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnType) == 1 {
		upsert = append(upsert, "COLUMN_TYPE = VALUES(COLUMN_TYPE)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnKey) == 1 {
		upsert = append(upsert, "COLUMN_KEY = VALUES(COLUMN_KEY)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsExtra) == 1 {
		upsert = append(upsert, "EXTRA = VALUES(EXTRA)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsPrivileges) == 1 {
		upsert = append(upsert, "PRIVILEGES = VALUES(PRIVILEGES)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnComment) == 1 {
		upsert = append(upsert, "COLUMN_COMMENT = VALUES(COLUMN_COMMENT)")
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsGenerationExpression) == 1 {
		upsert = append(upsert, "GENERATION_EXPRESSION = VALUES(GENERATION_EXPRESSION)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.COLUMNS")
	sql.Columns("table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "column_default", "is_nullable", "data_type", "character_maximum_length", "character_octet_length", "numeric_precision", "numeric_scale", "datetime_precision", "character_set_name", "collation_name", "column_type", "column_key", "extra", "privileges", "column_comment", "generation_expression")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// UpsertOne inserts the Columns to the database.
func (co *ColumnsStore) UpsertOne(data *Columns) (int64, error) {
	return co.Upsert([]*Columns{data})
}

// Upsert executes upsert for array of Columns
func (co *ColumnsStore) Upsert(data []*Columns) (int64, error) {
	sql := co.columnsUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "ColumnsUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := co.db.Exec(sql.Query())
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

// Insert inserts the Columns to the database.
func (co *ColumnsStore) Insert(data *Columns) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO information_schema.COLUMNS (")
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
		log.Debug().Str("fn", "information_schema.COLUMNS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("ColumnName", data.ColumnName).Uint64("OrdinalPosition", data.OrdinalPosition).Str("ColumnDefault", data.ColumnDefault).Str("IsNullable", data.IsNullable).Str("DataType", data.DataType).Uint64("CharacterMaximumLength", data.CharacterMaximumLength).Uint64("CharacterOctetLength", data.CharacterOctetLength).Uint64("NumericPrecision", data.NumericPrecision).Uint64("NumericScale", data.NumericScale).Uint64("DatetimePrecision", data.DatetimePrecision).Str("CharacterSetName", data.CharacterSetName).Str("CollationName", data.CollationName).Str("ColumnType", data.ColumnType).Str("ColumnKey", data.ColumnKey).Str("Extra", data.Extra).Str("Privileges", data.Privileges).Str("ColumnComment", data.ColumnComment).Str("GenerationExpression", data.GenerationExpression).Msg("sql")
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
func (co *ColumnsStore) Update(data *Columns) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.COLUMNS SET")
	if co.colSet == nil || co.colSet.Bit(ColumnsTableCatalog) == 1 {
		sql.AppendRaw(prepend, "TABLE_CATALOG = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableSchema) == 1 {
		sql.AppendRaw(prepend, "TABLE_SCHEMA = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableName) == 1 {
		sql.AppendRaw(prepend, "TABLE_NAME = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnName) == 1 {
		sql.AppendRaw(prepend, "COLUMN_NAME = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsOrdinalPosition) == 1 {
		sql.AppendRaw(prepend, "ORDINAL_POSITION = ?")
		prepend = ","
		args = append(args, data.OrdinalPosition)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnDefault) == 1 {
		sql.AppendRaw(prepend, "COLUMN_DEFAULT = ?")
		prepend = ","
		args = append(args, data.ColumnDefault)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsIsNullable) == 1 {
		sql.AppendRaw(prepend, "IS_NULLABLE = ?")
		prepend = ","
		args = append(args, data.IsNullable)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDataType) == 1 {
		sql.AppendRaw(prepend, "DATA_TYPE = ?")
		prepend = ","
		args = append(args, data.DataType)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterMaximumLength) == 1 {
		sql.AppendRaw(prepend, "CHARACTER_MAXIMUM_LENGTH = ?")
		prepend = ","
		args = append(args, data.CharacterMaximumLength)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterOctetLength) == 1 {
		sql.AppendRaw(prepend, "CHARACTER_OCTET_LENGTH = ?")
		prepend = ","
		args = append(args, data.CharacterOctetLength)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericPrecision) == 1 {
		sql.AppendRaw(prepend, "NUMERIC_PRECISION = ?")
		prepend = ","
		args = append(args, data.NumericPrecision)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericScale) == 1 {
		sql.AppendRaw(prepend, "NUMERIC_SCALE = ?")
		prepend = ","
		args = append(args, data.NumericScale)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDatetimePrecision) == 1 {
		sql.AppendRaw(prepend, "DATETIME_PRECISION = ?")
		prepend = ","
		args = append(args, data.DatetimePrecision)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterSetName) == 1 {
		sql.AppendRaw(prepend, "CHARACTER_SET_NAME = ?")
		prepend = ","
		args = append(args, data.CharacterSetName)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCollationName) == 1 {
		sql.AppendRaw(prepend, "COLLATION_NAME = ?")
		prepend = ","
		args = append(args, data.CollationName)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnType) == 1 {
		sql.AppendRaw(prepend, "COLUMN_TYPE = ?")
		prepend = ","
		args = append(args, data.ColumnType)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnKey) == 1 {
		sql.AppendRaw(prepend, "COLUMN_KEY = ?")
		prepend = ","
		args = append(args, data.ColumnKey)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsExtra) == 1 {
		sql.AppendRaw(prepend, "EXTRA = ?")
		prepend = ","
		args = append(args, data.Extra)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsPrivileges) == 1 {
		sql.AppendRaw(prepend, "PRIVILEGES = ?")
		prepend = ","
		args = append(args, data.Privileges)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnComment) == 1 {
		sql.AppendRaw(prepend, "COLUMN_COMMENT = ?")
		prepend = ","
		args = append(args, data.ColumnComment)
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsGenerationExpression) == 1 {
		sql.AppendRaw(prepend, "GENERATION_EXPRESSION = ?")
		args = append(args, data.GenerationExpression)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.COLUMNS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err :=
		co.db.Exec(sql.Query(), args...)
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
func (co *ColumnsStore) ToJSON(t *buffer.TemplateBuffer, data *Columns) {
	prepend := "{"
	if co.colSet == nil || co.colSet.Bit(ColumnsTableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsTableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnName) == 1 {
		t.JS(prepend, "column_name", data.ColumnName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsOrdinalPosition) == 1 {
		t.JD64u(prepend, "ordinal_position", data.OrdinalPosition)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnDefault) == 1 {
		t.JS(prepend, "column_default", data.ColumnDefault)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsIsNullable) == 1 {
		t.JS(prepend, "is_nullable", data.IsNullable)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDataType) == 1 {
		t.JS(prepend, "data_type", data.DataType)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterMaximumLength) == 1 {
		t.JD64u(prepend, "character_maximum_length", data.CharacterMaximumLength)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterOctetLength) == 1 {
		t.JD64u(prepend, "character_octet_length", data.CharacterOctetLength)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericPrecision) == 1 {
		t.JD64u(prepend, "numeric_precision", data.NumericPrecision)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsNumericScale) == 1 {
		t.JD64u(prepend, "numeric_scale", data.NumericScale)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsDatetimePrecision) == 1 {
		t.JD64u(prepend, "datetime_precision", data.DatetimePrecision)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCharacterSetName) == 1 {
		t.JS(prepend, "character_set_name", data.CharacterSetName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsCollationName) == 1 {
		t.JS(prepend, "collation_name", data.CollationName)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnType) == 1 {
		t.JS(prepend, "column_type", data.ColumnType)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnKey) == 1 {
		t.JS(prepend, "column_key", data.ColumnKey)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsExtra) == 1 {
		t.JS(prepend, "extra", data.Extra)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsPrivileges) == 1 {
		t.JS(prepend, "privileges", data.Privileges)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsColumnComment) == 1 {
		t.JS(prepend, "column_comment", data.ColumnComment)
		prepend = ","
	}
	if co.colSet == nil || co.colSet.Bit(ColumnsGenerationExpression) == 1 {
		t.JS(prepend, "generation_expression", data.GenerationExpression)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (co *ColumnsStore) ToJSONArray(w io.Writer, data []*Columns, name string) {
	t := buffer.NewTemplateBuffer()
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
