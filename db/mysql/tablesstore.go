package mysql

import (
	"database/sql"
	"io"
	"math/big"
	"time"

	codegen "bitbucket.org/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

// Tables represents a row from 'TABLES'.
type Tables struct {
	codegen.Tables
}

// new implements Bindable.new
func (ta *Tables) new() Bindable {
	return &Tables{}
}

// helper struct for common query operations.
type TablesSlice struct {
	data []*Tables
}

// append implements BindableSlice.append
func (ta *TablesSlice) append(d Bindable) {
	ta.data = append(ta.data, d.(*Tables))
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
	if colSet.Bit(Tables_TableCatalog) == 1 {
		fields = append(fields, "table_catalog")
	}

	if colSet.Bit(Tables_TableSchema) == 1 {
		fields = append(fields, "table_schema")
	}

	if colSet.Bit(Tables_TableName) == 1 {
		fields = append(fields, "table_name")
	}

	if colSet.Bit(Tables_TableType) == 1 {
		fields = append(fields, "table_type")
	}

	if colSet.Bit(Tables_Engine) == 1 {
		fields = append(fields, "engine")
	}

	if colSet.Bit(Tables_Version) == 1 {
		fields = append(fields, "version")
	}

	if colSet.Bit(Tables_RowFormat) == 1 {
		fields = append(fields, "row_format")
	}

	if colSet.Bit(Tables_TableRows) == 1 {
		fields = append(fields, "table_rows")
	}

	if colSet.Bit(Tables_AvgRowLength) == 1 {
		fields = append(fields, "avg_row_length")
	}

	if colSet.Bit(Tables_DataLength) == 1 {
		fields = append(fields, "data_length")
	}

	if colSet.Bit(Tables_MaxDataLength) == 1 {
		fields = append(fields, "max_data_length")
	}

	if colSet.Bit(Tables_IndexLength) == 1 {
		fields = append(fields, "index_length")
	}

	if colSet.Bit(Tables_DataFree) == 1 {
		fields = append(fields, "data_free")
	}

	if colSet.Bit(Tables_AutoIncrement) == 1 {
		fields = append(fields, "auto_increment")
	}

	if colSet.Bit(Tables_CreateTime) == 1 {
		fields = append(fields, "create_time")
	}

	if colSet.Bit(Tables_UpdateTime) == 1 {
		fields = append(fields, "update_time")
	}

	if colSet.Bit(Tables_CheckTime) == 1 {
		fields = append(fields, "check_time")
	}

	if colSet.Bit(Tables_TableCollation) == 1 {
		fields = append(fields, "table_collation")
	}

	if colSet.Bit(Tables_Checksum) == 1 {
		fields = append(fields, "checksum")
	}

	if colSet.Bit(Tables_CreateOptions) == 1 {
		fields = append(fields, "create_options")
	}

	if colSet.Bit(Tables_TableComment) == 1 {
		fields = append(fields, "table_comment")
	}
	return fields
}

// TablesStore is used to query for 'Tables' records.
type TablesStore struct {
	Store
}

// NewTablesStore return DAO Store for Tables
func NewTablesStore(conn Execer) *TablesStore {
	ta := &TablesStore{}
	ta.db = conn
	ta.withJoin = true
	ta.joinType = sdb.LEFT
	ta.batch = 1000
	return ta
}

// WithoutJoins won't execute JOIN when querying for records.
func (ta *TablesStore) WithoutJoins() *TablesStore {
	ta.withJoin = false
	return ta
}

// Where sets local sql, that will be appended to SELECT.
func (ta *TablesStore) Where(sql string) *TablesStore {
	ta.where = sql
	return ta
}

// OrderBy sets local sql, that will be appended to SELECT.
func (ta *TablesStore) OrderBy(sql string) *TablesStore {
	ta.orderBy = sql
	return ta
}

// GroupBy sets local sql, that will be appended to SELECT.
func (ta *TablesStore) GroupBy(sql string) *TablesStore {
	ta.groupBy = sql
	return ta
}

// Limit result set size
func (ta *TablesStore) Limit(n int) *TablesStore {
	ta.limit = n
	return ta
}

// Offset used, if a limit is provided
func (ta *TablesStore) Offset(n int) *TablesStore {
	ta.offset = n
	return ta
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (ta *TablesStore) JoinType(jt string) *TablesStore {
	ta.joinType = jt
	return ta
}

// Columns sets bits for specific columns.
func (ta *TablesStore) Columns(cols ...int) *TablesStore {
	ta.Store.Columns(cols...)
	return ta
}

// nolint[gocyclo]
func (ta *Tables) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(Tables_TableCatalog) == 1 {
		ta.TableCatalog = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableSchema) == 1 {
		ta.TableSchema = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableName) == 1 {
		ta.TableName = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableType) == 1 {
		ta.TableType = sdb.ToString(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_Engine) == 1 {
		if row[*col] == nil {
			ta.Engine = nil
		} else {
			ta.Engine = new(string)
			*ta.Engine = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_Version) == 1 {
		if row[*col] == nil {
			ta.Version = nil
		} else {
			ta.Version = new(uint64)
			*ta.Version = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_RowFormat) == 1 {
		if row[*col] == nil {
			ta.RowFormat = nil
		} else {
			ta.RowFormat = new(string)
			*ta.RowFormat = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableRows) == 1 {
		if row[*col] == nil {
			ta.TableRows = nil
		} else {
			ta.TableRows = new(uint64)
			*ta.TableRows = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_AvgRowLength) == 1 {
		if row[*col] == nil {
			ta.AvgRowLength = nil
		} else {
			ta.AvgRowLength = new(uint64)
			*ta.AvgRowLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_DataLength) == 1 {
		if row[*col] == nil {
			ta.DataLength = nil
		} else {
			ta.DataLength = new(uint64)
			*ta.DataLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_MaxDataLength) == 1 {
		if row[*col] == nil {
			ta.MaxDataLength = nil
		} else {
			ta.MaxDataLength = new(uint64)
			*ta.MaxDataLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_IndexLength) == 1 {
		if row[*col] == nil {
			ta.IndexLength = nil
		} else {
			ta.IndexLength = new(uint64)
			*ta.IndexLength = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_DataFree) == 1 {
		if row[*col] == nil {
			ta.DataFree = nil
		} else {
			ta.DataFree = new(uint64)
			*ta.DataFree = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_AutoIncrement) == 1 {
		if row[*col] == nil {
			ta.AutoIncrement = nil
		} else {
			ta.AutoIncrement = new(uint64)
			*ta.AutoIncrement = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_CreateTime) == 1 {
		if row[*col] == nil {
			ta.CreateTime = nil
		} else {
			ta.CreateTime = new(time.Time)
			*ta.CreateTime = sdb.ToTime(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_UpdateTime) == 1 {
		if row[*col] == nil {
			ta.UpdateTime = nil
		} else {
			ta.UpdateTime = new(time.Time)
			*ta.UpdateTime = sdb.ToTime(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_CheckTime) == 1 {
		if row[*col] == nil {
			ta.CheckTime = nil
		} else {
			ta.CheckTime = new(time.Time)
			*ta.CheckTime = sdb.ToTime(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableCollation) == 1 {
		if row[*col] == nil {
			ta.TableCollation = nil
		} else {
			ta.TableCollation = new(string)
			*ta.TableCollation = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_Checksum) == 1 {
		if row[*col] == nil {
			ta.Checksum = nil
		} else {
			ta.Checksum = new(uint64)
			*ta.Checksum = sdb.ToUInt64(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_CreateOptions) == 1 {
		if row[*col] == nil {
			ta.CreateOptions = nil
		} else {
			ta.CreateOptions = new(string)
			*ta.CreateOptions = sdb.ToString(row[*col])
		}
		*col++
	}
	if colSet == nil || colSet.Bit(Tables_TableComment) == 1 {
		ta.TableComment = sdb.ToString(row[*col])
		*col++
	}

}

func (ta *TablesStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", TablesQueryFields(ta.colSet))
	sql.Append("FROM information_schema.TABLES A")
	if ta.where != "" {
		sql.Append("WHERE", ta.where)
	}
	if ta.groupBy != "" {
		sql.Append("GROUP BY", ta.groupBy)
	}
	if ta.orderBy != "" {
		sql.Append("ORDER BY", ta.orderBy)
	}
	if ta.limit > 0 {
		sql.AppendRaw("LIMIT ", ta.limit)
		if ta.offset > 0 {
			sql.AppendRaw(",", ta.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'information_schema.TABLES' as a slice of Tables with 1:1 joined data.
func (ta *TablesStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Tables, error) {
	dto := &Tables{}
	data := &TablesSlice{}
	err := ta.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Tables, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Tables
	}
	return retValues, nil
}

// One retrieves a row from 'information_schema.TABLES' as a Tables with 1:1 joined data.
func (ta *TablesStore) One(args ...interface{}) (*codegen.Tables, error) {
	data := &Tables{}

	err := ta.one(data, ta.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Tables, nil
}

// Query retrieves many rows from 'information_schema.TABLES' as a slice of Tables with 1:1 joined data.
func (ta *TablesStore) Query(args ...interface{}) ([]*codegen.Tables, error) {
	stmt := ta.selectStatement()
	return ta.QueryCustom(stmt.Query(), args...)
}

// tablesUpsertStmt helper for generating Upsert statement.
// nolint[gocyclo]
func (ta *TablesStore) tablesUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCatalog) == 1 {
		upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableSchema) == 1 {
		upsert = append(upsert, "table_schema = VALUES(table_schema)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableName) == 1 {
		upsert = append(upsert, "table_name = VALUES(table_name)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableType) == 1 {
		upsert = append(upsert, "table_type = VALUES(table_type)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Engine) == 1 {
		upsert = append(upsert, "engine = VALUES(engine)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Version) == 1 {
		upsert = append(upsert, "version = VALUES(version)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_RowFormat) == 1 {
		upsert = append(upsert, "row_format = VALUES(row_format)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableRows) == 1 {
		upsert = append(upsert, "table_rows = VALUES(table_rows)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AvgRowLength) == 1 {
		upsert = append(upsert, "avg_row_length = VALUES(avg_row_length)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataLength) == 1 {
		upsert = append(upsert, "data_length = VALUES(data_length)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_MaxDataLength) == 1 {
		upsert = append(upsert, "max_data_length = VALUES(max_data_length)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_IndexLength) == 1 {
		upsert = append(upsert, "index_length = VALUES(index_length)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataFree) == 1 {
		upsert = append(upsert, "data_free = VALUES(data_free)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AutoIncrement) == 1 {
		upsert = append(upsert, "auto_increment = VALUES(auto_increment)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateTime) == 1 {
		upsert = append(upsert, "create_time = VALUES(create_time)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_UpdateTime) == 1 {
		upsert = append(upsert, "update_time = VALUES(update_time)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CheckTime) == 1 {
		upsert = append(upsert, "check_time = VALUES(check_time)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCollation) == 1 {
		upsert = append(upsert, "table_collation = VALUES(table_collation)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Checksum) == 1 {
		upsert = append(upsert, "checksum = VALUES(checksum)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateOptions) == 1 {
		upsert = append(upsert, "create_options = VALUES(create_options)")
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableComment) == 1 {
		upsert = append(upsert, "table_comment = VALUES(table_comment)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("information_schema.TABLES")
	sql.Columns("table_catalog", "table_schema", "table_name", "table_type", "engine", "version", "row_format", "table_rows", "avg_row_length", "data_length", "max_data_length", "index_length", "data_free", "auto_increment", "create_time", "update_time", "check_time", "table_collation", "checksum", "create_options", "table_comment")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// UpsertOne inserts the Tables to the database.
func (ta *TablesStore) UpsertOne(data *codegen.Tables) (int64, error) {
	return ta.Upsert([]*codegen.Tables{data})
}

// Upsert executes upsert for array of Tables
func (ta *TablesStore) Upsert(data []*codegen.Tables) (int64, error) {
	sql := ta.tablesUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "TablesUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := ta.db.Exec(sql.Query())
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

// Insert inserts the Tables to the database.
func (ta *TablesStore) Insert(data *codegen.Tables) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO information_schema.TABLES (")
	fields := TablesQueryFields(ta.colSet)
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
		log.Debug().Str("fn", "information_schema.TABLES.Insert").Str("stmt", sql.String()).Str("TableCatalog", data.TableCatalog).Str("TableSchema", data.TableSchema).Str("TableName", data.TableName).Str("TableType", data.TableType).Str("Engine", logString(data.Engine)).Uint64("Version", logUInt64(data.Version)).Str("RowFormat", logString(data.RowFormat)).Uint64("TableRows", logUInt64(data.TableRows)).Uint64("AvgRowLength", logUInt64(data.AvgRowLength)).Uint64("DataLength", logUInt64(data.DataLength)).Uint64("MaxDataLength", logUInt64(data.MaxDataLength)).Uint64("IndexLength", logUInt64(data.IndexLength)).Uint64("DataFree", logUInt64(data.DataFree)).Uint64("AutoIncrement", logUInt64(data.AutoIncrement)).Time("CreateTime", logTime(data.CreateTime)).Time("UpdateTime", logTime(data.UpdateTime)).Time("CheckTime", logTime(data.CheckTime)).Str("TableCollation", logString(data.TableCollation)).Uint64("Checksum", logUInt64(data.Checksum)).Str("CreateOptions", logString(data.CreateOptions)).Str("TableComment", data.TableComment).Msg("sql")
	}
	_, err = ta.db.Exec(sql.Query(), data.TableCatalog, data.TableSchema, data.TableName, data.TableType, data.Engine, data.Version, data.RowFormat, data.TableRows, data.AvgRowLength, data.DataLength, data.MaxDataLength, data.IndexLength, data.DataFree, data.AutoIncrement, data.CreateTime, data.UpdateTime, data.CheckTime, data.TableCollation, data.Checksum, data.CreateOptions, data.TableComment)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the Tables in the database.
// nolint[gocyclo]
func (ta *TablesStore) Update(data *codegen.Tables) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE information_schema.TABLES SET")
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCatalog) == 1 {
		sql.AppendRaw(prepend, "table_catalog = ?")
		prepend = ","
		args = append(args, data.TableCatalog)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableSchema) == 1 {
		sql.AppendRaw(prepend, "table_schema = ?")
		prepend = ","
		args = append(args, data.TableSchema)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableName) == 1 {
		sql.AppendRaw(prepend, "table_name = ?")
		prepend = ","
		args = append(args, data.TableName)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableType) == 1 {
		sql.AppendRaw(prepend, "table_type = ?")
		prepend = ","
		args = append(args, data.TableType)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Engine) == 1 {
		sql.AppendRaw(prepend, "engine = ?")
		prepend = ","
		args = append(args, data.Engine)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Version) == 1 {
		sql.AppendRaw(prepend, "version = ?")
		prepend = ","
		args = append(args, data.Version)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_RowFormat) == 1 {
		sql.AppendRaw(prepend, "row_format = ?")
		prepend = ","
		args = append(args, data.RowFormat)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableRows) == 1 {
		sql.AppendRaw(prepend, "table_rows = ?")
		prepend = ","
		args = append(args, data.TableRows)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AvgRowLength) == 1 {
		sql.AppendRaw(prepend, "avg_row_length = ?")
		prepend = ","
		args = append(args, data.AvgRowLength)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataLength) == 1 {
		sql.AppendRaw(prepend, "data_length = ?")
		prepend = ","
		args = append(args, data.DataLength)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_MaxDataLength) == 1 {
		sql.AppendRaw(prepend, "max_data_length = ?")
		prepend = ","
		args = append(args, data.MaxDataLength)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_IndexLength) == 1 {
		sql.AppendRaw(prepend, "index_length = ?")
		prepend = ","
		args = append(args, data.IndexLength)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataFree) == 1 {
		sql.AppendRaw(prepend, "data_free = ?")
		prepend = ","
		args = append(args, data.DataFree)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AutoIncrement) == 1 {
		sql.AppendRaw(prepend, "auto_increment = ?")
		prepend = ","
		args = append(args, data.AutoIncrement)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateTime) == 1 {
		sql.AppendRaw(prepend, "create_time = ?")
		prepend = ","
		args = append(args, data.CreateTime)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_UpdateTime) == 1 {
		sql.AppendRaw(prepend, "update_time = ?")
		prepend = ","
		args = append(args, data.UpdateTime)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CheckTime) == 1 {
		sql.AppendRaw(prepend, "check_time = ?")
		prepend = ","
		args = append(args, data.CheckTime)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCollation) == 1 {
		sql.AppendRaw(prepend, "table_collation = ?")
		prepend = ","
		args = append(args, data.TableCollation)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Checksum) == 1 {
		sql.AppendRaw(prepend, "checksum = ?")
		prepend = ","
		args = append(args, data.Checksum)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateOptions) == 1 {
		sql.AppendRaw(prepend, "create_options = ?")
		prepend = ","
		args = append(args, data.CreateOptions)
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableComment) == 1 {
		sql.AppendRaw(prepend, "table_comment = ?")
		args = append(args, data.TableComment)
	}
	sql.Append(" WHERE ")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.TABLES.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := ta.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Tables.
func (ta *TablesStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE information_schema.TABLES")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "information_schema.TABLES.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := ta.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (ta *TablesStore) ToJSON(t *sdb.JsonBuffer, data *Tables) {
	prepend := "{"
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCatalog) == 1 {
		t.JS(prepend, "table_catalog", data.TableCatalog)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableSchema) == 1 {
		t.JS(prepend, "table_schema", data.TableSchema)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableName) == 1 {
		t.JS(prepend, "table_name", data.TableName)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableType) == 1 {
		t.JS(prepend, "table_type", data.TableType)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Engine) == 1 {
		t.JS(prepend, "engine", *data.Engine)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Version) == 1 {
		t.JD64u(prepend, "version", *data.Version)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_RowFormat) == 1 {
		t.JS(prepend, "row_format", *data.RowFormat)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableRows) == 1 {
		t.JD64u(prepend, "table_rows", *data.TableRows)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AvgRowLength) == 1 {
		t.JD64u(prepend, "avg_row_length", *data.AvgRowLength)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataLength) == 1 {
		t.JD64u(prepend, "data_length", *data.DataLength)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_MaxDataLength) == 1 {
		t.JD64u(prepend, "max_data_length", *data.MaxDataLength)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_IndexLength) == 1 {
		t.JD64u(prepend, "index_length", *data.IndexLength)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_DataFree) == 1 {
		t.JD64u(prepend, "data_free", *data.DataFree)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_AutoIncrement) == 1 {
		t.JD64u(prepend, "auto_increment", *data.AutoIncrement)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateTime) == 1 {
		t.JT(prepend, "create_time", *data.CreateTime)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_UpdateTime) == 1 {
		t.JT(prepend, "update_time", *data.UpdateTime)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CheckTime) == 1 {
		t.JT(prepend, "check_time", *data.CheckTime)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableCollation) == 1 {
		t.JS(prepend, "table_collation", *data.TableCollation)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_Checksum) == 1 {
		t.JD64u(prepend, "checksum", *data.Checksum)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_CreateOptions) == 1 {
		t.JS(prepend, "create_options", *data.CreateOptions)
		prepend = ","
	}
	if ta.colSet == nil || ta.colSet.Bit(Tables_TableComment) == 1 {
		t.JS(prepend, "table_comment", data.TableComment)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (ta *TablesStore) ToJSONArray(w io.Writer, data []*Tables, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		ta.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
