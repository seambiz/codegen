package mysql

import (
	"database/sql"
	"io"
	"math/big"
	"time"

	codegen "bitbucket.org/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// Tables represents a row from 'TABLES'.
type Tables struct {
	codegen.Tables
}

// new implements Bindable.new
func (s *Tables) new() Bindable {
	return &Tables{}
}

// helper struct for common query operations.
type TablesSlice struct {
	data []*Tables
}

// append implements BindableSlice.append
func (s *TablesSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Tables))
}

// constant slice for all fields of the table "Tables".
// nolint[gochecknoglobals]
var tablesQueryFieldsAll = []string{"table_catalog", "table_schema", "table_name", "table_type", "engine", "version", "row_format", "table_rows", "avg_row_length", "data_length", "max_data_length", "index_length", "data_free", "auto_increment", "create_time", "update_time", "check_time", "table_collation", "checksum", "create_options", "table_comment"}

// returns fields, that should be used.
// nolint[gocyclo]
func TablesQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return tablesQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Tables_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(codegen.Tables_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(codegen.Tables_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(codegen.Tables_TableType) == 1 {
		fields = append(fields, "table_type")
	}

	if colSet.Bit(codegen.Tables_Engine) == 1 {
		fields = append(fields, "engine")
	}

	if colSet.Bit(codegen.Tables_Version) == 1 {
		fields = append(fields, "version")
	}

	if colSet.Bit(codegen.Tables_RowFormat) == 1 {
		fields = append(fields, "row_format")
	}

	if colSet.Bit(codegen.Tables_TableRows) == 1 {
		fields = append(fields, "table_rows")
	}

	if colSet.Bit(codegen.Tables_AvgRowLength) == 1 {
		fields = append(fields, "avg_row_length")
	}

	if colSet.Bit(codegen.Tables_DataLength) == 1 {
		fields = append(fields, "data_length")
	}

	if colSet.Bit(codegen.Tables_MaxDataLength) == 1 {
		fields = append(fields, "max_data_length")
	}

	if colSet.Bit(codegen.Tables_IndexLength) == 1 {
		fields = append(fields, "index_length")
	}

	if colSet.Bit(codegen.Tables_DataFree) == 1 {
		fields = append(fields, "data_free")
	}

	if colSet.Bit(codegen.Tables_AutoIncrement) == 1 {
		fields = append(fields, "auto_increment")
	}

	if colSet.Bit(codegen.Tables_CreateTime) == 1 {
		fields = append(fields, "create_time")
	}

	if colSet.Bit(codegen.Tables_UpdateTime) == 1 {
		fields = append(fields, "update_time")
	}

	if colSet.Bit(codegen.Tables_CheckTime) == 1 {
		fields = append(fields, "check_time")
	}

	if colSet.Bit(codegen.Tables_TableCollation) == 1 {
		fields = append(fields, "table_collation")
	}

	if colSet.Bit(codegen.Tables_Checksum) == 1 {
		fields = append(fields, "checksum")
	}

	if colSet.Bit(codegen.Tables_CreateOptions) == 1 {
		fields = append(fields, "create_options")
	}

	if colSet.Bit(codegen.Tables_TableComment) == 1 {
		fields = append(fields, "table_comment")
	}
	return fields
}

// TablesStore is used to query for 'Tables' records.
type TablesStore struct {
	Store
}

// NewTablesStore return DAO Store for Tables
func NewTablesStore(ctx *codegen.BaseContext, conn Execer) *TablesStore {
	s := &TablesStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *TablesStore) WithoutJoins() *TablesStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *TablesStore) Where(sql string) *TablesStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *TablesStore) OrderBy(sql string) *TablesStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *TablesStore) GroupBy(sql string) *TablesStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *TablesStore) Limit(n int) *TablesStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *TablesStore) Offset(n int) *TablesStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *TablesStore) JoinType(jt string) *TablesStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *TablesStore) Columns(cols ...int) *TablesStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *TablesStore) SetBits(colSet *big.Int) *TablesStore {
	s.colSet = colSet
	return s
}

func (s *Tables) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindInformationSchemaTables(&s.Tables, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindInformationSchemaTables(s *codegen.Tables, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Tables_TableCatalog) == 1 {
		s.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableSchema) == 1 {
		s.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableName) == 1 {
		s.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableType) == 1 {
		s.TableType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_Engine) == 1 {
		if row[*col] == nil {
			s.Engine = nil
		} else {
			s.Engine = new(string)
			*s.Engine = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_Version) == 1 {
		if row[*col] == nil {
			s.Version = nil
		} else {
			s.Version = new(int)
			*s.Version = sdb.ToInt(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_RowFormat) == 1 {
		if row[*col] == nil {
			s.RowFormat = nil
		} else {
			s.RowFormat = new(string)
			*s.RowFormat = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableRows) == 1 {
		if row[*col] == nil {
			s.TableRows = nil
		} else {
			s.TableRows = new(uint64)
			*s.TableRows = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_AvgRowLength) == 1 {
		if row[*col] == nil {
			s.AvgRowLength = nil
		} else {
			s.AvgRowLength = new(uint64)
			*s.AvgRowLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_DataLength) == 1 {
		if row[*col] == nil {
			s.DataLength = nil
		} else {
			s.DataLength = new(uint64)
			*s.DataLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_MaxDataLength) == 1 {
		if row[*col] == nil {
			s.MaxDataLength = nil
		} else {
			s.MaxDataLength = new(uint64)
			*s.MaxDataLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_IndexLength) == 1 {
		if row[*col] == nil {
			s.IndexLength = nil
		} else {
			s.IndexLength = new(uint64)
			*s.IndexLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_DataFree) == 1 {
		if row[*col] == nil {
			s.DataFree = nil
		} else {
			s.DataFree = new(uint64)
			*s.DataFree = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_AutoIncrement) == 1 {
		if row[*col] == nil {
			s.AutoIncrement = nil
		} else {
			s.AutoIncrement = new(uint64)
			*s.AutoIncrement = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_CreateTime) == 1 {
		s.CreateTime = sdb.ToTime(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_UpdateTime) == 1 {
		if row[*col] == nil {
			s.UpdateTime = nil
		} else {
			s.UpdateTime = new(time.Time)
			*s.UpdateTime = sdb.ToTime(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_CheckTime) == 1 {
		if row[*col] == nil {
			s.CheckTime = nil
		} else {
			s.CheckTime = new(time.Time)
			*s.CheckTime = sdb.ToTime(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableCollation) == 1 {
		if row[*col] == nil {
			s.TableCollation = nil
		} else {
			s.TableCollation = new(string)
			*s.TableCollation = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_Checksum) == 1 {
		if row[*col] == nil {
			s.Checksum = nil
		} else {
			s.Checksum = new(int64)
			*s.Checksum = sdb.ToInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_CreateOptions) == 1 {
		if row[*col] == nil {
			s.CreateOptions = nil
		} else {
			s.CreateOptions = new(string)
			*s.CreateOptions = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tables_TableComment) == 1 {
		if row[*col] == nil {
			s.TableComment = nil
		} else {
			s.TableComment = new(string)
			*s.TableComment = sdb.ToString(row[*col])
		}
		*col++
	}
}

func (s *TablesStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", TablesQueryFields(s.colSet))
	sql.Append(" FROM information_schema.TABLES A ")
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

// QueryCustom retrieves many rows from 'information_schema.TABLES' as a slice of Tables with 1:1 joined data.
func (s *TablesStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Tables, error) {
	dto := &Tables{}
	data := &TablesSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Tables, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Tables
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.TABLES' as a Tables with 1:1 joined data.
func (s *TablesStore) One(args ...interface{}) (*codegen.Tables, error) {
	data := &Tables{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Tables, nil
}

// Query retrieves many rows from 'information_schema.TABLES' as a slice of Tables with 1:1 joined data.
func (s *TablesStore) Query(args ...interface{}) ([]*codegen.Tables, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// tablesUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *TablesStore) tablesUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableType) == 1 {
		upsert = append(upsert, "table_type = VALUES(table_type)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Engine) == 1 {
		upsert = append(upsert, "engine = VALUES(engine)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Version) == 1 {
		upsert = append(upsert, "version = VALUES(version)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_RowFormat) == 1 {
		upsert = append(upsert, "row_format = VALUES(row_format)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableRows) == 1 {
		upsert = append(upsert, "table_rows = VALUES(table_rows)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AvgRowLength) == 1 {
		upsert = append(upsert, "avg_row_length = VALUES(avg_row_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataLength) == 1 {
		upsert = append(upsert, "data_length = VALUES(data_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_MaxDataLength) == 1 {
		upsert = append(upsert, "max_data_length = VALUES(max_data_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_IndexLength) == 1 {
		upsert = append(upsert, "index_length = VALUES(index_length)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataFree) == 1 {
		upsert = append(upsert, "data_free = VALUES(data_free)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AutoIncrement) == 1 {
		upsert = append(upsert, "auto_increment = VALUES(auto_increment)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateTime) == 1 {
		upsert = append(upsert, "create_time = VALUES(create_time)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_UpdateTime) == 1 {
		upsert = append(upsert, "update_time = VALUES(update_time)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CheckTime) == 1 {
		upsert = append(upsert, "check_time = VALUES(check_time)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCollation) == 1 {
		upsert = append(upsert, "table_collation = VALUES(table_collation)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Checksum) == 1 {
		upsert = append(upsert, "checksum = VALUES(checksum)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateOptions) == 1 {
		upsert = append(upsert, "create_options = VALUES(create_options)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableComment) == 1 {
		upsert = append(upsert, "table_comment = VALUES(table_comment)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.TABLES")
	sql.Columns("table_catalog", "table_schema", "table_name", "table_type", "engine", "version", "row_format", "table_rows", "avg_row_length", "data_length", "max_data_length", "index_length", "data_free", "auto_increment", "create_time", "update_time", "check_time", "table_collation", "checksum", "create_options", "table_comment")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Tables
func (s *TablesStore) Upsert(data ...*codegen.Tables) (int64, error) {
	sql := s.tablesUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "TablesUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the Tables to the database.
func (s *TablesStore) Insert(data *codegen.Tables) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO information_schema.TABLES (")
	fields := TablesQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "information_schema.TABLES.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("TableType", data.TableType).Str("Engine", logString(data.Engine)).Int("Version", logInt(data.Version)).Str("RowFormat", logString(data.RowFormat)).Uint64("TableRows", logUInt64(data.TableRows)).Uint64("AvgRowLength", logUInt64(data.AvgRowLength)).Uint64("DataLength", logUInt64(data.DataLength)).Uint64("MaxDataLength", logUInt64(data.MaxDataLength)).Uint64("IndexLength", logUInt64(data.IndexLength)).Uint64("DataFree", logUInt64(data.DataFree)).Uint64("AutoIncrement", logUInt64(data.AutoIncrement)).Time("CreateTime", data.CreateTime).Time("UpdateTime", logTime(data.UpdateTime)).Time("CheckTime", logTime(data.CheckTime)).Str("TableCollation", logString(data.TableCollation)).Int64("Checksum", logInt64(data.Checksum)).Str("CreateOptions", logString(data.CreateOptions)).Str("TableComment", logString(data.TableComment)).Msg("sql")
	}
	_, err = s.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.TableType, data.Engine, data.Version, data.RowFormat, data.TableRows, data.AvgRowLength, data.DataLength, data.MaxDataLength, data.IndexLength, data.DataFree, data.AutoIncrement, data.CreateTime, data.UpdateTime, data.CheckTime, data.TableCollation, data.Checksum, data.CreateOptions, data.TableComment)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Tables in the database.
// nolint[gocyclo]
func (s *TablesStore) Update(data *codegen.Tables) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.TABLES SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableType) == 1 {
		sql.AppendRaw(prepend, "table_type = ?")
		prepend = ","
		args = append(args, data.TableType)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Engine) == 1 {
		sql.AppendRaw(prepend, "engine = ?")
		prepend = ","
		args = append(args, data.Engine)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Version) == 1 {
		sql.AppendRaw(prepend, "version = ?")
		prepend = ","
		args = append(args, data.Version)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_RowFormat) == 1 {
		sql.AppendRaw(prepend, "row_format = ?")
		prepend = ","
		args = append(args, data.RowFormat)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableRows) == 1 {
		sql.AppendRaw(prepend, "table_rows = ?")
		prepend = ","
		args = append(args, data.TableRows)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AvgRowLength) == 1 {
		sql.AppendRaw(prepend, "avg_row_length = ?")
		prepend = ","
		args = append(args, data.AvgRowLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataLength) == 1 {
		sql.AppendRaw(prepend, "data_length = ?")
		prepend = ","
		args = append(args, data.DataLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_MaxDataLength) == 1 {
		sql.AppendRaw(prepend, "max_data_length = ?")
		prepend = ","
		args = append(args, data.MaxDataLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_IndexLength) == 1 {
		sql.AppendRaw(prepend, "index_length = ?")
		prepend = ","
		args = append(args, data.IndexLength)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataFree) == 1 {
		sql.AppendRaw(prepend, "data_free = ?")
		prepend = ","
		args = append(args, data.DataFree)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AutoIncrement) == 1 {
		sql.AppendRaw(prepend, "auto_increment = ?")
		prepend = ","
		args = append(args, data.AutoIncrement)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateTime) == 1 {
		sql.AppendRaw(prepend, "create_time = ?")
		prepend = ","
		args = append(args, data.CreateTime)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_UpdateTime) == 1 {
		sql.AppendRaw(prepend, "update_time = ?")
		prepend = ","
		args = append(args, data.UpdateTime)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CheckTime) == 1 {
		sql.AppendRaw(prepend, "check_time = ?")
		prepend = ","
		args = append(args, data.CheckTime)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCollation) == 1 {
		sql.AppendRaw(prepend, "table_collation = ?")
		prepend = ","
		args = append(args, data.TableCollation)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Checksum) == 1 {
		sql.AppendRaw(prepend, "checksum = ?")
		prepend = ","
		args = append(args, data.Checksum)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateOptions) == 1 {
		sql.AppendRaw(prepend, "create_options = ?")
		prepend = ","
		args = append(args, data.CreateOptions)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableComment) == 1 {
		sql.AppendRaw(prepend, "table_comment = ?")
		args = append(args, data.TableComment)
	}
	sql.Append(" WHERE ")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.TABLES.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Tables.
func (s *TablesStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.TABLES")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "information_schema.TABLES.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *TablesStore) ToJSON(t *sdb.JsonBuffer, data *Tables) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableType) == 1 {
		t.JS(prepend, "table_type", data.TableType)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Engine) == 1 {
		t.JS(prepend, "engine", *data.Engine)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Version) == 1 {
		t.JD(prepend, "version", *data.Version)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_RowFormat) == 1 {
		t.JS(prepend, "row_format", *data.RowFormat)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableRows) == 1 {
		t.JD64u(prepend, "table_rows", *data.TableRows)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AvgRowLength) == 1 {
		t.JD64u(prepend, "avg_row_length", *data.AvgRowLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataLength) == 1 {
		t.JD64u(prepend, "data_length", *data.DataLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_MaxDataLength) == 1 {
		t.JD64u(prepend, "max_data_length", *data.MaxDataLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_IndexLength) == 1 {
		t.JD64u(prepend, "index_length", *data.IndexLength)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_DataFree) == 1 {
		t.JD64u(prepend, "data_free", *data.DataFree)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_AutoIncrement) == 1 {
		t.JD64u(prepend, "auto_increment", *data.AutoIncrement)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateTime) == 1 {
		t.JT(prepend, "create_time", data.CreateTime)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_UpdateTime) == 1 {
		t.JT(prepend, "update_time", *data.UpdateTime)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CheckTime) == 1 {
		t.JT(prepend, "check_time", *data.CheckTime)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableCollation) == 1 {
		t.JS(prepend, "table_collation", *data.TableCollation)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_Checksum) == 1 {
		t.JD64(prepend, "checksum", *data.Checksum)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_CreateOptions) == 1 {
		t.JS(prepend, "create_options", *data.CreateOptions)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tables_TableComment) == 1 {
		t.JS(prepend, "table_comment", *data.TableComment)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *TablesStore) ToJSONArray(w io.Writer, data []*Tables, name string) {
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
