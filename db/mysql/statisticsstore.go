package mysql

import (
	"database/sql"
	"io"
	"math/big"

	codegen "bitbucket.org/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// Statistics represents a row from 'STATISTICS'.
type Statistics struct {
	codegen.Statistics
}

// new implements Bindable.new
func (s *Statistics) new() Bindable {
	return &Statistics{}
}

// helper struct for common query operations.
type StatisticsSlice struct {
	data []*Statistics
}

// append implements BindableSlice.append
func (s *StatisticsSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Statistics))
}

// constant slice for all fields of the table "Statistics".
// nolint[gochecknoglobals]
var statisticsQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment", "is_visible", "expression"}

// returns fields, that should be used.
// nolint[gocyclo]
func StatisticsQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return statisticsQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Statistics_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(codegen.Statistics_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(codegen.Statistics_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(codegen.Statistics_NonUnique) == 1 {
		fields = append(fields, "non_unique")
	}

	if colSet.Bit(codegen.Statistics_IndexSchema) == 1 {
		fields = append(fields, "index_schema")
	}

	if colSet.Bit(codegen.Statistics_IndexName) == 1 {
		fields = append(fields, "index_name")
	}

	if colSet.Bit(codegen.Statistics_SeqInIndex) == 1 {
		fields = append(fields, "seq_in_index")
	}

	if colSet.Bit(codegen.Statistics_ColumnName) == 1 {
		fields = append(fields, "column_name")
	}

	if colSet.Bit(codegen.Statistics_Collation) == 1 {
		fields = append(fields, "collation")
	}

	if colSet.Bit(codegen.Statistics_Cardinality) == 1 {
		fields = append(fields, "cardinality")
	}

	if colSet.Bit(codegen.Statistics_SubPart) == 1 {
		fields = append(fields, "sub_part")
	}

	if colSet.Bit(codegen.Statistics_Packed) == 1 {
		fields = append(fields, "packed")
	}

	if colSet.Bit(codegen.Statistics_Nullable) == 1 {
		fields = append(fields, "nullable")
	}

	if colSet.Bit(codegen.Statistics_IndexType) == 1 {
		fields = append(fields, "index_type")
	}

	if colSet.Bit(codegen.Statistics_Comment) == 1 {
		fields = append(fields, "comment")
	}

	if colSet.Bit(codegen.Statistics_IndexComment) == 1 {
		fields = append(fields, "index_comment")
	}

	if colSet.Bit(codegen.Statistics_IsVisible) == 1 {
		fields = append(fields, "is_visible")
	}

	if colSet.Bit(codegen.Statistics_Expression) == 1 {
		fields = append(fields, "expression")
	}
	return fields
}

// StatisticsStore is used to query for 'Statistics' records.
type StatisticsStore struct {
	Store
}

// NewStatisticsStore return DAO Store for Statistics
func NewStatisticsStore(ctx *codegen.BaseContext, conn Execer) *StatisticsStore {
	s := &StatisticsStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *StatisticsStore) WithoutJoins() *StatisticsStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *StatisticsStore) Where(sql string) *StatisticsStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *StatisticsStore) OrderBy(sql string) *StatisticsStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *StatisticsStore) GroupBy(sql string) *StatisticsStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *StatisticsStore) Limit(n int) *StatisticsStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *StatisticsStore) Offset(n int) *StatisticsStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *StatisticsStore) JoinType(jt string) *StatisticsStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *StatisticsStore) Columns(cols ...int) *StatisticsStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *StatisticsStore) SetBits(colSet *big.Int) *StatisticsStore {
	s.colSet = colSet
	return s
}

func (s *Statistics) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaStatistics(&s.Statistics, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindInformationSchemaStatistics(s *codegen.Statistics, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Statistics_TableCatalog) == 1 {
		s.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_TableSchema) == 1 {
		s.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_TableName) == 1 {
		s.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_NonUnique) == 1 {
		s.NonUnique = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_IndexSchema) == 1 {
		s.IndexSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_IndexName) == 1 {
		if row[*col] == nil {
			s.IndexName = nil
		} else {
			s.IndexName = new(string)
			*s.IndexName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_SeqInIndex) == 1 {
		s.SeqInIndex = sdb.ToUInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_ColumnName) == 1 {
		if row[*col] == nil {
			s.ColumnName = nil
		} else {
			s.ColumnName = new(string)
			*s.ColumnName = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Collation) == 1 {
		if row[*col] == nil {
			s.Collation = nil
		} else {
			s.Collation = new(string)
			*s.Collation = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Cardinality) == 1 {
		if row[*col] == nil {
			s.Cardinality = nil
		} else {
			s.Cardinality = new(int64)
			*s.Cardinality = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_SubPart) == 1 {
		if row[*col] == nil {
			s.SubPart = nil
		} else {
			s.SubPart = new(int64)
			*s.SubPart = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Packed) == 1 {
		if row[*col] == nil {
			s.Packed = nil
		} else {
			s.Packed = new([]byte)
			*s.Packed = (row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Nullable) == 1 {
		s.Nullable = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_IndexType) == 1 {
		s.IndexType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Comment) == 1 {
		s.Comment = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_IndexComment) == 1 {
		s.IndexComment = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_IsVisible) == 1 {
		s.IsVisible = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Statistics_Expression) == 1 {
		if row[*col] == nil {
			s.Expression = nil
		} else {
			s.Expression = new(string)
			*s.Expression = sdb.ToString(row[*col])
		}
		*col++
	}
}

func (s *StatisticsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", StatisticsQueryFields(s.colSet))
	sql.Append(" FROM information_schema.STATISTICS A ")
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

// QueryCustom retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with 1:1 joined data.
func (s *StatisticsStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Statistics, error) {
	dto := &Statistics{}
	data := &StatisticsSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Statistics, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Statistics
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.STATISTICS' as a Statistics with 1:1 joined data.
func (s *StatisticsStore) One(args ...interface{}) (*codegen.Statistics, error) {
	data := &Statistics{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Statistics, nil
}

// Query retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with 1:1 joined data.
func (s *StatisticsStore) Query(args ...interface{}) ([]*codegen.Statistics, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// statisticsUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *StatisticsStore) statisticsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_NonUnique) == 1 {
		upsert = append(upsert, "non_unique = VALUES(non_unique)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexSchema) == 1 {
		upsert = append(upsert, "index_schema = VALUES(index_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexName) == 1 {
		upsert = append(upsert, "index_name = VALUES(index_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SeqInIndex) == 1 {
		upsert = append(upsert, "seq_in_index = VALUES(seq_in_index)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_ColumnName) == 1 {
		upsert = append(upsert, "column_name = VALUES(column_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Collation) == 1 {
		upsert = append(upsert, "collation = VALUES(collation)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Cardinality) == 1 {
		upsert = append(upsert, "cardinality = VALUES(cardinality)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SubPart) == 1 {
		upsert = append(upsert, "sub_part = VALUES(sub_part)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Packed) == 1 {
		upsert = append(upsert, "packed = VALUES(packed)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Nullable) == 1 {
		upsert = append(upsert, "nullable = VALUES(nullable)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexType) == 1 {
		upsert = append(upsert, "index_type = VALUES(index_type)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Comment) == 1 {
		upsert = append(upsert, "comment = VALUES(comment)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexComment) == 1 {
		upsert = append(upsert, "index_comment = VALUES(index_comment)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IsVisible) == 1 {
		upsert = append(upsert, "is_visible = VALUES(is_visible)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Expression) == 1 {
		upsert = append(upsert, "expression = VALUES(expression)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.STATISTICS")
	sql.Columns("table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment", "is_visible", "expression")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Statistics
func (s *StatisticsStore) Upsert(data ...*codegen.Statistics) (int64, error) {
	sql := s.statisticsUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "StatisticsUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the Statistics to the database.
func (s *StatisticsStore) Insert(data *codegen.Statistics) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.STATISTICS (")
	fields := StatisticsQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "information_schema.STATISTICS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Int("NonUnique", data.NonUnique).Str("IndexSchema", data.IndexSchema).Str("IndexName", logString(data.IndexName)).Uint("SeqInIndex", data.SeqInIndex).Str("ColumnName", logString(data.ColumnName)).Str("Collation", logString(data.Collation)).Int64("Cardinality", logInt64(data.Cardinality)).Int64("SubPart", logInt64(data.SubPart)).Bytes("Packed", logBytes(data.Packed)).Str("Nullable", data.Nullable).Str("IndexType", data.IndexType).Str("Comment", data.Comment).Str("IndexComment", data.IndexComment).Str("IsVisible", data.IsVisible).Str("Expression", logString(data.Expression)).Msg("sql")
	}
	_, err = s.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.NonUnique, data.IndexSchema, data.IndexName, data.SeqInIndex, data.ColumnName, data.Collation, data.Cardinality, data.SubPart, data.Packed, data.Nullable, data.IndexType, data.Comment, data.IndexComment, data.IsVisible, data.Expression)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Statistics in the database.
// nolint[gocyclo]
func (s *StatisticsStore) Update(data *codegen.Statistics) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.STATISTICS SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_NonUnique) == 1 {
		sql.AppendRaw(prepend, "non_unique = ?")
		prepend = ","
		args = append(args, data.NonUnique)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexSchema) == 1 {
		sql.AppendRaw(prepend, "index_schema = ?")
		prepend = ","
		args = append(args, data.IndexSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexName) == 1 {
		sql.AppendRaw(prepend, "index_name = ?")
		prepend = ","
		args = append(args, data.IndexName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SeqInIndex) == 1 {
		sql.AppendRaw(prepend, "seq_in_index = ?")
		prepend = ","
		args = append(args, data.SeqInIndex)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Collation) == 1 {
		sql.AppendRaw(prepend, "collation = ?")
		prepend = ","
		args = append(args, data.Collation)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Cardinality) == 1 {
		sql.AppendRaw(prepend, "cardinality = ?")
		prepend = ","
		args = append(args, data.Cardinality)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SubPart) == 1 {
		sql.AppendRaw(prepend, "sub_part = ?")
		prepend = ","
		args = append(args, data.SubPart)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Packed) == 1 {
		sql.AppendRaw(prepend, "packed = ?")
		prepend = ","
		args = append(args, data.Packed)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Nullable) == 1 {
		sql.AppendRaw(prepend, "nullable = ?")
		prepend = ","
		args = append(args, data.Nullable)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexType) == 1 {
		sql.AppendRaw(prepend, "index_type = ?")
		prepend = ","
		args = append(args, data.IndexType)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Comment) == 1 {
		sql.AppendRaw(prepend, "comment = ?")
		prepend = ","
		args = append(args, data.Comment)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexComment) == 1 {
		sql.AppendRaw(prepend, "index_comment = ?")
		prepend = ","
		args = append(args, data.IndexComment)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IsVisible) == 1 {
		sql.AppendRaw(prepend, "is_visible = ?")
		prepend = ","
		args = append(args, data.IsVisible)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Expression) == 1 {
		sql.AppendRaw(prepend, "expression = ?")
		args = append(args, data.Expression)
	}
	sql.Append(" WHERE ")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.STATISTICS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Statistics.
func (s *StatisticsStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.STATISTICS")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.STATISTICS.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *StatisticsStore) ToJSON(t *sdb.JsonBuffer, data *Statistics) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_NonUnique) == 1 {
		t.JD(prepend, "non_unique", data.NonUnique)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexSchema) == 1 {
		t.JS(prepend, "index_schema", data.IndexSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexName) == 1 {
		t.JS(prepend, "index_name", *data.IndexName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SeqInIndex) == 1 {
		t.JDu(prepend, "seq_in_index", data.SeqInIndex)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_ColumnName) == 1 {
		t.JS(prepend, "column_name", *data.ColumnName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Collation) == 1 {
		t.JS(prepend, "collation", *data.Collation)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Cardinality) == 1 {
		t.JD64(prepend, "cardinality", *data.Cardinality)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_SubPart) == 1 {
		t.JD64(prepend, "sub_part", *data.SubPart)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Packed) == 1 {
		t.JByte(prepend, "packed", *data.Packed)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Nullable) == 1 {
		t.JS(prepend, "nullable", data.Nullable)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexType) == 1 {
		t.JS(prepend, "index_type", data.IndexType)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Comment) == 1 {
		t.JS(prepend, "comment", data.Comment)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IndexComment) == 1 {
		t.JS(prepend, "index_comment", data.IndexComment)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_IsVisible) == 1 {
		t.JS(prepend, "is_visible", data.IsVisible)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Statistics_Expression) == 1 {
		t.JS(prepend, "expression", *data.Expression)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *StatisticsStore) ToJSONArray(w io.Writer, data []*Statistics, name string) {
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
