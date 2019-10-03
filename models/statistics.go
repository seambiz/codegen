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
var statisticsQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment"}

// returns fields, that should be used.
// nolint[gocyclo]
func StatisticsQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return statisticsQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(StatisticsTableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}
	if colSet.Bit(StatisticsTableSchema) == 1 {
		fields = append(fields, "table_schema")
	}
	if colSet.Bit(StatisticsTableName) == 1 {
		fields = append(fields, "table_name")
	}
	if colSet.Bit(StatisticsNonUnique) == 1 {
		fields = append(fields, "non_unique")
	}
	if colSet.Bit(StatisticsIndexSchema) == 1 {
		fields = append(fields, "index_schema")
	}
	if colSet.Bit(StatisticsIndexName) == 1 {
		fields = append(fields, "index_name")
	}
	if colSet.Bit(StatisticsSeqInIndex) == 1 {
		fields = append(fields, "seq_in_index")
	}
	if colSet.Bit(StatisticsColumnName) == 1 {
		fields = append(fields, "column_name")
	}
	if colSet.Bit(StatisticsCollation) == 1 {
		fields = append(fields, "collation")
	}
	if colSet.Bit(StatisticsCardinality) == 1 {
		fields = append(fields, "cardinality")
	}
	if colSet.Bit(StatisticsSubPart) == 1 {
		fields = append(fields, "sub_part")
	}
	if colSet.Bit(StatisticsPacked) == 1 {
		fields = append(fields, "packed")
	}
	if colSet.Bit(StatisticsNullable) == 1 {
		fields = append(fields, "nullable")
	}
	if colSet.Bit(StatisticsIndexType) == 1 {
		fields = append(fields, "index_type")
	}
	if colSet.Bit(StatisticsComment) == 1 {
		fields = append(fields, "comment")
	}
	if colSet.Bit(StatisticsIndexComment) == 1 {
		fields = append(fields, "index_comment")
	}
	return fields
}

// Statistics represents a row from 'information_schema.STATISTICS'.
type Statistics struct {
	TableCatalog string `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema  string `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName    string `json:"TABLE_NAME" db:"table_name"`
	NonUnique    int64  `json:"NON_UNIQUE" db:"non_unique"`
	IndexSchema  string `json:"INDEX_SCHEMA" db:"index_schema"`
	IndexName    string `json:"INDEX_NAME" db:"index_name"`
	SeqInIndex   int64  `json:"SEQ_IN_INDEX" db:"seq_in_index"`
	ColumnName   string `json:"COLUMN_NAME" db:"column_name"`
	Collation    string `json:"COLLATION" db:"collation"`
	Cardinality  int64  `json:"CARDINALITY" db:"cardinality"`
	SubPart      int64  `json:"SUB_PART" db:"sub_part"`
	Packed       string `json:"PACKED" db:"packed"`
	Nullable     string `json:"NULLABLE" db:"nullable"`
	IndexType    string `json:"INDEX_TYPE" db:"index_type"`
	Comment      string `json:"COMMENT" db:"comment"`
	IndexComment string `json:"INDEX_COMMENT" db:"index_comment"`
}

// new implements DTO.new
func (st *Statistics) new() DTO {
	return &Statistics{}
}

// helper struct for common query operations.
type StatisticsSlice struct {
	data []*Statistics
}

// append implements DTOSlice.append
func (st *StatisticsSlice) append(d DTO) {
	st.data = append(st.data, d.(*Statistics))
}

// Columns to be used for various statements.
func (st *StatisticsStore) Columns(cols ...int) *StatisticsStore {
	st.colSet = big.NewInt(0)
	for _, col := range cols {
		st.colSet.SetBit(st.colSet, col, 1)
	}
	return st
}

// IsEmpty checks if primary key fields are zero.
func (st *Statistics) IsEmpty() bool {
	return true
}

// StatisticsStore is used to query for 'Statistics' records.
type StatisticsStore struct {
	Store
}

// NewStatisticsStore return DAO Store for Statistics
func NewStatisticsStore(conn *sql.DB) *StatisticsStore {
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

// nolint[gocyclo]
func (st *Statistics) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(StatisticsTableCatalog) == 1 {
		st.TableCatalog = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsTableSchema) == 1 {
		st.TableSchema = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsTableName) == 1 {
		st.TableName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsNonUnique) == 1 {
		st.NonUnique = convert.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsIndexSchema) == 1 {
		st.IndexSchema = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsIndexName) == 1 {
		st.IndexName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsSeqInIndex) == 1 {
		st.SeqInIndex = convert.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsColumnName) == 1 {
		st.ColumnName = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsCollation) == 1 {
		st.Collation = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsCardinality) == 1 {
		st.Cardinality = convert.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsSubPart) == 1 {
		st.SubPart = convert.ToInt64(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsPacked) == 1 {
		st.Packed = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsNullable) == 1 {
		st.Nullable = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsIndexType) == 1 {
		st.IndexType = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsComment) == 1 {
		st.Comment = convert.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(StatisticsIndexComment) == 1 {
		st.IndexComment = convert.ToString(row[*col])
		*col++
	}
}
func (st *StatisticsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", StatisticsQueryFields(st.colSet))
	sql.Append("FROM information_schema.STATISTICS A")
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

// One retrieves a row from 'information_schema.STATISTICS' as a Statistics with possible joined data.
func (st *StatisticsStore) One(args ...interface{}) (*Statistics, error) {
	data := &Statistics{}

	err := st.one(data, st.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return data, nil
}

// Query retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with possible joined data.
func (st *StatisticsStore) Query(args ...interface{}) ([]*Statistics, error) {
	stmt := st.selectStatement()
	return st.QueryCustom(stmt.Query(), args...)
}

// QueryCustom retrieves many rows from 'information_schema.STATISTICS' as a slice of Statistics with possible joined data.
func (st *StatisticsStore) QueryCustom(stmt string, args ...interface{}) ([]*Statistics, error) {
	dto := &Statistics{}
	data := &StatisticsSlice{}
	err := st.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	return data.data, nil
}

// statisticsUpsertStmt helper for generating Upserts general statement
// nolint[gocyclo]
func (st *StatisticsStore) statisticsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableCatalog) == 1 {
		upsert = append(upsert, "TABLE_CATALOG = VALUES(TABLE_CATALOG)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableSchema) == 1 {
		upsert = append(upsert, "TABLE_SCHEMA = VALUES(TABLE_SCHEMA)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableName) == 1 {
		upsert = append(upsert, "TABLE_NAME = VALUES(TABLE_NAME)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNonUnique) == 1 {
		upsert = append(upsert, "NON_UNIQUE = VALUES(NON_UNIQUE)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexSchema) == 1 {
		upsert = append(upsert, "INDEX_SCHEMA = VALUES(INDEX_SCHEMA)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexName) == 1 {
		upsert = append(upsert, "INDEX_NAME = VALUES(INDEX_NAME)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSeqInIndex) == 1 {
		upsert = append(upsert, "SEQ_IN_INDEX = VALUES(SEQ_IN_INDEX)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsColumnName) == 1 {
		upsert = append(upsert, "COLUMN_NAME = VALUES(COLUMN_NAME)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCollation) == 1 {
		upsert = append(upsert, "COLLATION = VALUES(COLLATION)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCardinality) == 1 {
		upsert = append(upsert, "CARDINALITY = VALUES(CARDINALITY)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSubPart) == 1 {
		upsert = append(upsert, "SUB_PART = VALUES(SUB_PART)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsPacked) == 1 {
		upsert = append(upsert, "PACKED = VALUES(PACKED)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNullable) == 1 {
		upsert = append(upsert, "NULLABLE = VALUES(NULLABLE)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexType) == 1 {
		upsert = append(upsert, "INDEX_TYPE = VALUES(INDEX_TYPE)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsComment) == 1 {
		upsert = append(upsert, "COMMENT = VALUES(COMMENT)")
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexComment) == 1 {
		upsert = append(upsert, "INDEX_COMMENT = VALUES(INDEX_COMMENT)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.STATISTICS")
	sql.Columns("table_catalog", "table_schema", "table_name", "non_unique", "index_schema", "index_name", "seq_in_index", "column_name", "collation", "cardinality", "sub_part", "packed", "nullable", "index_type", "comment", "index_comment")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// UpsertOne inserts the Statistics to the database.
func (st *StatisticsStore) UpsertOne(data *Statistics) (int64, error) {
	return st.Upsert([]*Statistics{data})
}

// Upsert executes upsert for array of Statistics
func (st *StatisticsStore) Upsert(data []*Statistics) (int64, error) {
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
func (st *StatisticsStore) Insert(data *Statistics) error {
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
		log.Debug().Str("fn", "information_schema.STATISTICS.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Int64("NonUnique", data.NonUnique).Str("IndexSchema", data.IndexSchema).Str("IndexName", data.IndexName).Int64("SeqInIndex", data.SeqInIndex).Str("ColumnName", data.ColumnName).Str("Collation", data.Collation).Int64("Cardinality", data.Cardinality).Int64("SubPart", data.SubPart).Str("Packed", data.Packed).Str("Nullable", data.Nullable).Str("IndexType", data.IndexType).Str("Comment", data.Comment).Str("IndexComment", data.IndexComment).Msg("sql")
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
func (st *StatisticsStore) Update(data *Statistics) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.STATISTICS SET")
	if st.colSet == nil || st.colSet.Bit(StatisticsTableCatalog) == 1 {
		sql.AppendRaw(prepend, "TABLE_CATALOG = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableSchema) == 1 {
		sql.AppendRaw(prepend, "TABLE_SCHEMA = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableName) == 1 {
		sql.AppendRaw(prepend, "TABLE_NAME = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNonUnique) == 1 {
		sql.AppendRaw(prepend, "NON_UNIQUE = ?")
		prepend = ","
		args = append(args, data.NonUnique)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexSchema) == 1 {
		sql.AppendRaw(prepend, "INDEX_SCHEMA = ?")
		prepend = ","
		args = append(args, data.IndexSchema)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexName) == 1 {
		sql.AppendRaw(prepend, "INDEX_NAME = ?")
		prepend = ","
		args = append(args, data.IndexName)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSeqInIndex) == 1 {
		sql.AppendRaw(prepend, "SEQ_IN_INDEX = ?")
		prepend = ","
		args = append(args, data.SeqInIndex)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsColumnName) == 1 {
		sql.AppendRaw(prepend, "COLUMN_NAME = ?")
		prepend = ","
		args = append(args, data.ColumnName)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCollation) == 1 {
		sql.AppendRaw(prepend, "COLLATION = ?")
		prepend = ","
		args = append(args, data.Collation)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCardinality) == 1 {
		sql.AppendRaw(prepend, "CARDINALITY = ?")
		prepend = ","
		args = append(args, data.Cardinality)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSubPart) == 1 {
		sql.AppendRaw(prepend, "SUB_PART = ?")
		prepend = ","
		args = append(args, data.SubPart)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsPacked) == 1 {
		sql.AppendRaw(prepend, "PACKED = ?")
		prepend = ","
		args = append(args, data.Packed)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNullable) == 1 {
		sql.AppendRaw(prepend, "NULLABLE = ?")
		prepend = ","
		args = append(args, data.Nullable)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexType) == 1 {
		sql.AppendRaw(prepend, "INDEX_TYPE = ?")
		prepend = ","
		args = append(args, data.IndexType)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsComment) == 1 {
		sql.AppendRaw(prepend, "COMMENT = ?")
		prepend = ","
		args = append(args, data.Comment)
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexComment) == 1 {
		sql.AppendRaw(prepend, "INDEX_COMMENT = ?")
		args = append(args, data.IndexComment)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.STATISTICS.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err :=
		st.db.Exec(sql.Query(), args...)
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
func (st *StatisticsStore) ToJSON(t *buffer.TemplateBuffer, data *Statistics) {
	prepend := "{"
	if st.colSet == nil || st.colSet.Bit(StatisticsTableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsTableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNonUnique) == 1 {
		t.JD64(prepend, "non_unique", data.NonUnique)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexSchema) == 1 {
		t.JS(prepend, "index_schema", data.IndexSchema)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexName) == 1 {
		t.JS(prepend, "index_name", data.IndexName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSeqInIndex) == 1 {
		t.JD64(prepend, "seq_in_index", data.SeqInIndex)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsColumnName) == 1 {
		t.JS(prepend, "column_name", data.ColumnName)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCollation) == 1 {
		t.JS(prepend, "collation", data.Collation)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsCardinality) == 1 {
		t.JD64(prepend, "cardinality", data.Cardinality)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsSubPart) == 1 {
		t.JD64(prepend, "sub_part", data.SubPart)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsPacked) == 1 {
		t.JS(prepend, "packed", data.Packed)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsNullable) == 1 {
		t.JS(prepend, "nullable", data.Nullable)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexType) == 1 {
		t.JS(prepend, "index_type", data.IndexType)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsComment) == 1 {
		t.JS(prepend, "comment", data.Comment)
		prepend = ","
	}
	if st.colSet == nil || st.colSet.Bit(StatisticsIndexComment) == 1 {
		t.JS(prepend, "index_comment", data.IndexComment)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (st *StatisticsStore) ToJSONArray(w io.Writer, data []*Statistics, name string) {
	t := buffer.NewTemplateBuffer()
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
