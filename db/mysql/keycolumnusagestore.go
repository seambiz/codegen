package mysql

import (
	"database/sql"
	"io"
	"math/big"

	codegen "bitbucket.org/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// KeyColumnUsage represents a row from 'KEY_COLUMN_USAGE'.
type KeyColumnUsage struct {
	codegen.KeyColumnUsage
}

// new implements Bindable.new
func (s *KeyColumnUsage) new() Bindable {
	return &KeyColumnUsage{}
}

// helper struct for common query operations.
type KeyColumnUsageSlice struct {
	data []*KeyColumnUsage
}

// append implements BindableSlice.append
func (s *KeyColumnUsageSlice) append(d Bindable) {
	s.data = append(s.data, d.(*KeyColumnUsage))
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
func NewKeyColumnUsageStore(ctx *codegen.BaseContext, conn Execer) *KeyColumnUsageStore {
	s := &KeyColumnUsageStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *KeyColumnUsageStore) WithoutJoins() *KeyColumnUsageStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *KeyColumnUsageStore) Where(sql string) *KeyColumnUsageStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *KeyColumnUsageStore) OrderBy(sql string) *KeyColumnUsageStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *KeyColumnUsageStore) GroupBy(sql string) *KeyColumnUsageStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *KeyColumnUsageStore) Limit(n int) *KeyColumnUsageStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *KeyColumnUsageStore) Offset(n int) *KeyColumnUsageStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *KeyColumnUsageStore) JoinType(jt string) *KeyColumnUsageStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *KeyColumnUsageStore) Columns(cols ...int) *KeyColumnUsageStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *KeyColumnUsageStore) SetBits(colSet *big.Int) *KeyColumnUsageStore {
	s.colSet = colSet
	return s
}

func (s *KeyColumnUsage) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaKeyColumnUsage(&s.KeyColumnUsage, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindInformationSchemaKeyColumnUsage(s *codegen.KeyColumnUsage, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		s.ConstraintCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		s.ConstraintSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		if row[*col] == nil {
			s.ConstraintName = nil
		} else {
			s.ConstraintName = new(string)
			*s.ConstraintName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		s.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		s.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		s.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		if row[*col] == nil {
			s.ColumnName = nil
		} else {
			s.ColumnName = new(string)
			*s.ColumnName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		s.OrdinalPosition = sdb.ToUInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		if row[*col] == nil {
			s.PositionInUniqueConstraint = nil
		} else {
			s.PositionInUniqueConstraint = new(uint)
			*s.PositionInUniqueConstraint = sdb.ToUInt(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		if row[*col] == nil {
			s.ReferencedTableSchema = nil
		} else {
			s.ReferencedTableSchema = new(string)
			*s.ReferencedTableSchema = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		if row[*col] == nil {
			s.ReferencedTableName = nil
		} else {
			s.ReferencedTableName = new(string)
			*s.ReferencedTableName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		if row[*col] == nil {
			s.ReferencedColumnName = nil
		} else {
			s.ReferencedColumnName = new(string)
			*s.ReferencedColumnName = sdb.ToString(row[*col])
		}
		*col++
	}
}

func (s *KeyColumnUsageStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", KeyColumnUsageQueryFields(s.colSet))
	sql.Append(" FROM information_schema.KEY_COLUMN_USAGE A ")
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

// QueryCustom retrieves many rows from 'information_schema.KEY_COLUMN_USAGE' as a slice of KeyColumnUsage with 1:1 joined data.
func (s *KeyColumnUsageStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.KeyColumnUsage, error) {
	dto := &KeyColumnUsage{}
	data := &KeyColumnUsageSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.KeyColumnUsage, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].KeyColumnUsage
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.KEY_COLUMN_USAGE' as a KeyColumnUsage with 1:1 joined data.
func (s *KeyColumnUsageStore) One(args ...interface{}) (*codegen.KeyColumnUsage, error) {
	data := &KeyColumnUsage{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.KeyColumnUsage, nil
}

// Query retrieves many rows from 'information_schema.KEY_COLUMN_USAGE' as a slice of KeyColumnUsage with 1:1 joined data.
func (s *KeyColumnUsageStore) Query(args ...interface{}) ([]*codegen.KeyColumnUsage, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// keyColumnUsageUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *KeyColumnUsageStore) keyColumnUsageUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		upsert = append(upsert, "constraint_catalog = VALUES(constraint_catalog)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		upsert = append(upsert, "constraint_schema = VALUES(constraint_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		upsert = append(upsert, "constraint_name = VALUES(constraint_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		upsert = append(upsert, "column_name = VALUES(column_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		upsert = append(upsert, "ordinal_position = VALUES(ordinal_position)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		upsert = append(upsert, "position_in_unique_constraint = VALUES(position_in_unique_constraint)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		upsert = append(upsert, "referenced_table_schema = VALUES(referenced_table_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		upsert = append(upsert, "referenced_table_name = VALUES(referenced_table_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		upsert = append(upsert, "referenced_column_name = VALUES(referenced_column_name)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.KEY_COLUMN_USAGE")
	sql.Columns("constraint_catalog", "constraint_schema", "constraint_name", "table_catalog", "table_schema", "table_name", "column_name", "ordinal_position", "position_in_unique_constraint", "referenced_table_schema", "referenced_table_name", "referenced_column_name")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of KeyColumnUsage
func (s *KeyColumnUsageStore) Upsert(data ...*codegen.KeyColumnUsage) (int64, error) {
	sql := s.keyColumnUsageUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "KeyColumnUsageUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the KeyColumnUsage to the database.
func (s *KeyColumnUsageStore) Insert(data *codegen.KeyColumnUsage) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.KEY_COLUMN_USAGE (")
	fields := KeyColumnUsageQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "information_schema.KEY_COLUMN_USAGE.Insert").Str("stmt", sql.String()).Str("ConstraintCatalog", data.ConstraintCatalog).Str("ConstraintSchema", data.ConstraintSchema).Str("ConstraintName", logString(data.ConstraintName)).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("ColumnName", logString(data.ColumnName)).Uint("OrdinalPosition", data.OrdinalPosition).Uint("PositionInUniqueConstraint", logUInt(data.PositionInUniqueConstraint)).Str("ReferencedTableSchema", logString(data.ReferencedTableSchema)).Str("ReferencedTableName", logString(data.ReferencedTableName)).Str("ReferencedColumnName", logString(data.ReferencedColumnName)).Msg("sql")
	}
	_, err = s.db.Exec(sql.Query(), data.ConstraintCatalog, data.ConstraintSchema, data.ConstraintName, data.TableCatalog, data.TableSchema, data.TableName, data.ColumnName, data.OrdinalPosition, data.PositionInUniqueConstraint, data.ReferencedTableSchema, data.ReferencedTableName, data.ReferencedColumnName)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the KeyColumnUsage in the database.
// nolint[gocyclo]
func (s *KeyColumnUsageStore) Update(data *codegen.KeyColumnUsage) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.KEY_COLUMN_USAGE SET")
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		sql.AppendRaw(prepend, "constraint_catalog = ?")
		prepend = ","
		args = append(args, data.ConstraintCatalog)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		sql.AppendRaw(prepend, "constraint_schema = ?")
		prepend = ","
		args = append(args, data.ConstraintSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		sql.AppendRaw(prepend, "constraint_name = ?")
		prepend = ","
		args = append(args, data.ConstraintName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		sql.AppendRaw(prepend, "ordinal_position = ?")
		prepend = ","
		args = append(args, data.OrdinalPosition)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		sql.AppendRaw(prepend, "position_in_unique_constraint = ?")
		prepend = ","
		args = append(args, data.PositionInUniqueConstraint)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		sql.AppendRaw(prepend, "referenced_table_schema = ?")
		prepend = ","
		args = append(args, data.ReferencedTableSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		sql.AppendRaw(prepend, "referenced_table_name = ?")
		prepend = ","
		args = append(args, data.ReferencedTableName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		sql.AppendRaw(prepend, "referenced_column_name = ?")
		args = append(args, data.ReferencedColumnName)
	}
	sql.Append(" WHERE ")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.KEY_COLUMN_USAGE.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from KeyColumnUsage.
func (s *KeyColumnUsageStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.KEY_COLUMN_USAGE")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.KEY_COLUMN_USAGE.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *KeyColumnUsageStore) ToJSON(t *sdb.JsonBuffer, data *KeyColumnUsage) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintCatalog) == 1 {
		t.JS(prepend, "constraint_catalog", data.ConstraintCatalog)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintSchema) == 1 {
		t.JS(prepend, "constraint_schema", data.ConstraintSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ConstraintName) == 1 {
		t.JS(prepend, "constraint_name", *data.ConstraintName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ColumnName) == 1 {
		t.JS(prepend, "column_name", *data.ColumnName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_OrdinalPosition) == 1 {
		t.JDu(prepend, "ordinal_position", data.OrdinalPosition)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_PositionInUniqueConstraint) == 1 {
		t.JDu(prepend, "position_in_unique_constraint", *data.PositionInUniqueConstraint)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableSchema) == 1 {
		t.JS(prepend, "referenced_table_schema", *data.ReferencedTableSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedTableName) == 1 {
		t.JS(prepend, "referenced_table_name", *data.ReferencedTableName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.KeyColumnUsage_ReferencedColumnName) == 1 {
		t.JS(prepend, "referenced_column_name", *data.ReferencedColumnName)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *KeyColumnUsageStore) ToJSONArray(w io.Writer, data []*KeyColumnUsage, name string) {
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
