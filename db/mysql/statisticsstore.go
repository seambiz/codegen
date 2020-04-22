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

// Statistics represents a row from 'STATISTICS'.
type Statistics struct {
	codegen.Statistics
}

// new implements Bindable.new
func (st *Statistics) new() Bindable {
	return &Statistics{}
}

// helper struct for common query operations.
type StatisticsSlice struct {
	data []*Statistics
}

// append implements BindableSlice.append
func (st *StatisticsSlice) append(d Bindable) {
	st.data = append(st.data, d.(*Statistics))
}

// constant slice for all fields of the table "Statistics".
// nolint[gochecknoglobals]
var statisticsQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment"}

// returns fields, that should be used.
// nolint[gocyclo]
func StatisticsQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return statisticsQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(Statistics_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(Statistics_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(Statistics_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(Statistics_NonUnique) == 1 {
		fields = append(fields, "non_unique")
	}

	if colSet.Bit(Statistics_IndexSchema) == 1 {
		fields = append(fields, "index_schema")
	}

	if colSet.Bit(Statistics_IndexName) == 1 {
		fields = append(fields, "index_name")
	}

	if colSet.Bit(Statistics_SeqInIndex) == 1 {
		fields = append(fields, "seq_in_index")
	}

	if colSet.Bit(Statistics_ColumnName) == 1 {
		fields = append(fields, "column_name")
	}

	if colSet.Bit(Statistics_Collation) == 1 {
		fields = append(fields, "collation")
	}

	if colSet.Bit(Statistics_Cardinality) == 1 {
		fields = append(fields, "cardinality")
	}

	if colSet.Bit(Statistics_SubPart) == 1 {
		fields = append(fields, "sub_part")
	}

	if colSet.Bit(Statistics_Packed) == 1 {
		fields = append(fields, "packed")
	}

	if colSet.Bit(Statistics_Nullable) == 1 {
		fields = append(fields, "nullable")
	}

	if colSet.Bit(Statistics_IndexType) == 1 {
		fields = append(fields, "index_type")
	}

	if colSet.Bit(Statistics_Comment) == 1 {
		fields = append(fields, "comment")
	}

	if colSet.Bit(Statistics_IndexComment) == 1 {
		fields = append(fields, "index_comment")
	}
	return fields
}

// StatisticsStore is used to query for 'Statistics' records.
type StatisticsStore struct {
	Store
}

// NewStatisticsStore return DAO Store for Statistics
func NewStatisticsStore(conn Execer) *StatisticsStore {
	st := &StatisticsStore{}
	st.db = conn
	st.withJoin = true
	st.joinType = sdb.LEFT
	st.batch = 1000
	return st
}

// WithoutJoins won't execute JOIN when querying for records.
func (st *StatisticsStore) WithoutJoins() *StatisticsStore {
	st.withJoin = false
	return st
}

// Where sets local sql, that will be appended to SELECT.
func (st *StatisticsStore) Where(sql string) *StatisticsStore {
	st.where = sql
	return st
}

// OrderBy sets local sql, that will be appended to SELECT.
func (st *StatisticsStore) OrderBy(sql string) *StatisticsStore {
	st.orderBy = sql
	return st
}

// GroupBy sets local sql, that will be appended to SELECT.
func (st *StatisticsStore) GroupBy(sql string) *StatisticsStore {
	st.groupBy = sql
	return st
}

// Limit result set size
func (st *StatisticsStore) Limit(n int) *StatisticsStore {
	st.limit = n
	return st
}

// Offset used, if a limit is provided
func (st *StatisticsStore) Offset(n int) *StatisticsStore {
	st.offset = n
	return st
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (st *StatisticsStore) JoinType(jt string) *StatisticsStore {
	st.joinType = jt
	return st
}

// Columns sets bits for specific columns.
func (st *StatisticsStore) Columns(cols ...int) *StatisticsStore {
	st.Store.Columns(cols...)
	return st
}

// nolint[gocyclo]
func (st *Statistics) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaStatistics(&st.Statistics, row, withJoin, colSet, col)
}

func BindInformationSchemaStatistics(st *codegen.Statistics, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(Statistics_TableCatalog) == 1 {
		st.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_TableSchema) == 1 {
		st.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_TableName) == 1 {
		st.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_NonUnique) == 1 {
		st.NonUnique = sdb.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_IndexSchema) == 1 {
		st.IndexSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_IndexName) == 1 {
		st.IndexName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_SeqInIndex) == 1 {
		st.SeqInIndex = sdb.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_ColumnName) == 1 {
		st.ColumnName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_Collation) == 1 {
		if row[*col] == nil {
			st.Collation = nil
		} else {
			st.Collation = new(string)
			*st.Collation = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_Cardinality) == 1 {
		if row[*col] == nil {
			st.Cardinality = nil
		} else {
			st.Cardinality = new(int64)
			*st.Cardinality = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_SubPart) == 1 {
		if row[*col] == nil {
			st.SubPart = nil
		} else {
			st.SubPart = new(int64)
			*st.SubPart = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_Packed) == 1 {
		if row[*col] == nil {
			st.Packed = nil
		} else {
			st.Packed = new(string)
			*st.Packed = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_Nullable) == 1 {
		st.Nullable = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_IndexType) == 1 {
		st.IndexType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_Comment) == 1 {
		if row[*col] == nil {
			st.Comment = nil
		} else {
			st.Comment = new(string)
			*st.Comment = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Statistics_IndexComment) == 1 {
		st.IndexComment = sdb.ToString(row[*col])
		*col++
	}
}

func (st *StatisticsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", StatisticsQueryFields(st.colSet))
	sql.Append(" FROM information_schema.STATISTICS A")
	if st.where != "" {
		sql.Append("WHERE", st.where)
	}
	if st.groupBy != "" {
		sql.Append("GROUP BY", st.groupBy)
	}
	if st.orderBy != "" {
		sql.Append("ORDER BY", st.orderBy)
	}
	if st.limit > 0 {
		sql.AppendRaw("LIMIT ", st.limit)
		if st.offset > 0 {
			sql.AppendRaw(",", st.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with 1:1 joined data.
func (st *StatisticsStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Statistics, error) {
	dto := &Statistics{}
	data := &StatisticsSlice{}
	err := st.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Statistics, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Statistics
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.STATISTICS' as a Statistics with 1:1 joined data.
func (st *StatisticsStore) One(args ...interface{}) (*codegen.Statistics, error) {
	data := &Statistics{}

	err := st.one(data, st.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Statistics, nil
}

// Query retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with 1:1 joined data.
func (st *StatisticsStore) Query(args ...interface{}) ([]*codegen.Statistics, error) {
	stmt := st.selectStatement()
	return st.QueryCustom(stmt.Query(), args...)
}

// statisticsUpsertStmt helper for generating Upsert statement.
// nolint[gocyclo]
func (st *StatisticsStore) statisticsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_NonUnique) == 1 {
		upsert = append(upsert, "non_unique = VALUES(non_unique)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexSchema) == 1 {
		upsert = append(upsert, "index_schema = VALUES(index_schema)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexName) == 1 {
		upsert = append(upsert, "index_name = VALUES(index_name)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SeqInIndex) == 1 {
		upsert = append(upsert, "seq_in_index = VALUES(seq_in_index)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_ColumnName) == 1 {
		upsert = append(upsert, "column_name = VALUES(column_name)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Collation) == 1 {
		upsert = append(upsert, "collation = VALUES(collation)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Cardinality) == 1 {
		upsert = append(upsert, "cardinality = VALUES(cardinality)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SubPart) == 1 {
		upsert = append(upsert, "sub_part = VALUES(sub_part)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Packed) == 1 {
		upsert = append(upsert, "packed = VALUES(packed)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Nullable) == 1 {
		upsert = append(upsert, "nullable = VALUES(nullable)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexType) == 1 {
		upsert = append(upsert, "index_type = VALUES(index_type)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Comment) == 1 {
		upsert = append(upsert, "comment = VALUES(comment)")
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexComment) == 1 {
		upsert = append(upsert, "index_comment = VALUES(index_comment)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.STATISTICS")
	sql.Columns("table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Statistics
func (st *StatisticsStore) Upsert(data ...*codegen.Statistics) (int64, error) {
	sql := st.statisticsUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "StatisticsUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := st.db.Exec(sql.Query())
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

// Insert inserts the Statistics to the database.
func (st *StatisticsStore) Insert(data *codegen.Statistics) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO information_schema.STATISTICS (")
	fields := StatisticsQueryFields(st.colSet)
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
		log.Debug().Str("fn", "information_schema.STATISTICS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Int64("NonUnique", data.NonUnique).Str("IndexSchema", data.IndexSchema).Str("IndexName", data.IndexName).Int64("SeqInIndex", data.SeqInIndex).Str("ColumnName", data.ColumnName).Str("Collation", logString(data.Collation)).Int64("Cardinality", logInt64(data.Cardinality)).Int64("SubPart", logInt64(data.SubPart)).Str("Packed", logString(data.Packed)).Str("Nullable", data.Nullable).Str("IndexType", data.IndexType).Str("Comment", logString(data.Comment)).Str("IndexComment", data.IndexComment).Msg("sql")
	}
	_, err = st.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.NonUnique, data.IndexSchema, data.IndexName, data.SeqInIndex, data.ColumnName, data.Collation, data.Cardinality, data.SubPart, data.Packed, data.Nullable, data.IndexType, data.Comment, data.IndexComment)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Statistics in the database.
// nolint[gocyclo]
func (st *StatisticsStore) Update(data *codegen.Statistics) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.STATISTICS SET")
	if st.colSet == nil || st.colSet.Bit(Statistics_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_NonUnique) == 1 {
		sql.AppendRaw(prepend, "non_unique = ?")
		prepend = ","
		args = append(args, data.NonUnique)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexSchema) == 1 {
		sql.AppendRaw(prepend, "index_schema = ?")
		prepend = ","
		args = append(args, data.IndexSchema)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexName) == 1 {
		sql.AppendRaw(prepend, "index_name = ?")
		prepend = ","
		args = append(args, data.IndexName)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SeqInIndex) == 1 {
		sql.AppendRaw(prepend, "seq_in_index = ?")
		prepend = ","
		args = append(args, data.SeqInIndex)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_ColumnName) == 1 {
		sql.AppendRaw(prepend, "column_name = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Collation) == 1 {
		sql.AppendRaw(prepend, "collation = ?")
		prepend = ","
		args = append(args, data.Collation)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Cardinality) == 1 {
		sql.AppendRaw(prepend, "cardinality = ?")
		prepend = ","
		args = append(args, data.Cardinality)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SubPart) == 1 {
		sql.AppendRaw(prepend, "sub_part = ?")
		prepend = ","
		args = append(args, data.SubPart)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Packed) == 1 {
		sql.AppendRaw(prepend, "packed = ?")
		prepend = ","
		args = append(args, data.Packed)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Nullable) == 1 {
		sql.AppendRaw(prepend, "nullable = ?")
		prepend = ","
		args = append(args, data.Nullable)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexType) == 1 {
		sql.AppendRaw(prepend, "index_type = ?")
		prepend = ","
		args = append(args, data.IndexType)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Comment) == 1 {
		sql.AppendRaw(prepend, "comment = ?")
		prepend = ","
		args = append(args, data.Comment)
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexComment) == 1 {
		sql.AppendRaw(prepend, "index_comment = ?")
		args = append(args, data.IndexComment)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.STATISTICS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := st.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Statistics.
func (st *StatisticsStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.STATISTICS")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.STATISTICS.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := st.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (st *StatisticsStore) ToJSON(t *sdb.JsonBuffer, data *Statistics) {
	prepend := "{"
	if st.colSet == nil || st.colSet.Bit(Statistics_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_NonUnique) == 1 {
		t.JD64(prepend, "non_unique", data.NonUnique)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexSchema) == 1 {
		t.JS(prepend, "index_schema", data.IndexSchema)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexName) == 1 {
		t.JS(prepend, "index_name", data.IndexName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SeqInIndex) == 1 {
		t.JD64(prepend, "seq_in_index", data.SeqInIndex)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_ColumnName) == 1 {
		t.JS(prepend, "column_name", data.ColumnName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Collation) == 1 {
		t.JS(prepend, "collation", *data.Collation)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Cardinality) == 1 {
		t.JD64(prepend, "cardinality", *data.Cardinality)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_SubPart) == 1 {
		t.JD64(prepend, "sub_part", *data.SubPart)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Packed) == 1 {
		t.JS(prepend, "packed", *data.Packed)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Nullable) == 1 {
		t.JS(prepend, "nullable", data.Nullable)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexType) == 1 {
		t.JS(prepend, "index_type", data.IndexType)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_Comment) == 1 {
		t.JS(prepend, "comment", *data.Comment)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(Statistics_IndexComment) == 1 {
		t.JS(prepend, "index_comment", data.IndexComment)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (st *StatisticsStore) ToJSONArray(w io.Writer, data []*Statistics, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		st.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
