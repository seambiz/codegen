package mysql

import (
	"database/sql"
	"io"
	"math/big"
	"time"

	codegen "github.com/seambiz/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// TABLES represents a row from 'TABLES'.
type TABLES struct {
	codegen.TABLES
}

// new implements Bindable.new
func (s *TABLES) new() Bindable {
	return &TABLES{}
}

// helper struct for common query operations.
type TABLESSlice struct {
	data []*TABLES
}

// append implements BindableSlice.append
func (s *TABLESSlice) append(d Bindable) {
	s.data = append(s.data, d.(*TABLES))
}



// constant slice for all fields of the table "TABLES".
// nolint[gochecknoglobals]
var tablesQueryFieldsAll = []string{"table_catalog" , "table_schema" , "table_name" , "table_type" , "engine" , "version" , "row_format" , "table_rows" , "avg_row_length" , "data_length" , "max_data_length" , "index_length" , "data_free" , "auto_increment" , "create_time" , "update_time" , "check_time" , "table_collation" , "checksum" , "create_options" , "table_comment"}

// returns fields, that should be used.
// nolint[gocyclo]
func TABLESQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return tablesQueryFieldsAll
	}

	fields := []string{}
    if colSet.Bit(codegen.TABLES_TABLECATALOG) == 1 {
		fields = append(fields, "table_catalog")
	}
	
    if colSet.Bit(codegen.TABLES_TABLESCHEMA) == 1 {
		fields = append(fields, "table_schema")
	}
	
    if colSet.Bit(codegen.TABLES_TABLENAME) == 1 {
		fields = append(fields, "table_name")
	}
	
    if colSet.Bit(codegen.TABLES_TABLETYPE) == 1 {
		fields = append(fields, "table_type")
	}
	
    if colSet.Bit(codegen.TABLES_ENGINE) == 1 {
		fields = append(fields, "engine")
	}
	
    if colSet.Bit(codegen.TABLES_VERSION) == 1 {
		fields = append(fields, "version")
	}
	
    if colSet.Bit(codegen.TABLES_ROWFORMAT) == 1 {
		fields = append(fields, "row_format")
	}
	
    if colSet.Bit(codegen.TABLES_TABLEROWS) == 1 {
		fields = append(fields, "table_rows")
	}
	
    if colSet.Bit(codegen.TABLES_AVGROWLENGTH) == 1 {
		fields = append(fields, "avg_row_length")
	}
	
    if colSet.Bit(codegen.TABLES_DATALENGTH) == 1 {
		fields = append(fields, "data_length")
	}
	
    if colSet.Bit(codegen.TABLES_MAXDATALENGTH) == 1 {
		fields = append(fields, "max_data_length")
	}
	
    if colSet.Bit(codegen.TABLES_INDEXLENGTH) == 1 {
		fields = append(fields, "index_length")
	}
	
    if colSet.Bit(codegen.TABLES_DATAFREE) == 1 {
		fields = append(fields, "data_free")
	}
	
    if colSet.Bit(codegen.TABLES_AUTOINCREMENT) == 1 {
		fields = append(fields, "auto_increment")
	}
	
    if colSet.Bit(codegen.TABLES_CREATETIME) == 1 {
		fields = append(fields, "create_time")
	}
	
    if colSet.Bit(codegen.TABLES_UPDATETIME) == 1 {
		fields = append(fields, "update_time")
	}
	
    if colSet.Bit(codegen.TABLES_CHECKTIME) == 1 {
		fields = append(fields, "check_time")
	}
	
    if colSet.Bit(codegen.TABLES_TABLECOLLATION) == 1 {
		fields = append(fields, "table_collation")
	}
	
    if colSet.Bit(codegen.TABLES_CHECKSUM) == 1 {
		fields = append(fields, "checksum")
	}
	
    if colSet.Bit(codegen.TABLES_CREATEOPTIONS) == 1 {
		fields = append(fields, "create_options")
	}
	
    if colSet.Bit(codegen.TABLES_TABLECOMMENT) == 1 {
		fields = append(fields, "table_comment")
	}
	return fields
}
// TABLESStore is used to query for 'TABLES' records.
type TABLESStore struct {
	Store
	ctx *codegen.Context
}

// NewTABLESStore return DAO Store for TABLES
func NewTABLESStore(ctx *codegen.Context, conn Execer) *TABLESStore {
	s := &TABLESStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	s.ctx = ctx
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *TABLESStore) WithoutJoins() *TABLESStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *TABLESStore) Where(sql string) *TABLESStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *TABLESStore) OrderBy(sql string) *TABLESStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *TABLESStore) GroupBy(sql string) *TABLESStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *TABLESStore) Limit(n int) *TABLESStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *TABLESStore) Offset(n int) *TABLESStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *TABLESStore) JoinType(jt string) *TABLESStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *TABLESStore) Columns(cols ...int) *TABLESStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *TABLESStore) SetBits(colSet *big.Int) *TABLESStore {
	s.colSet = colSet
	return s
}
func (s *TABLES) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLECATALOG) == 1 {
		s.TABLECATALOG = sdb.ToString(row[*col])
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLESCHEMA) == 1 {
		s.TABLESCHEMA = sdb.ToString(row[*col])
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLENAME) == 1 {
		s.TABLENAME = sdb.ToString(row[*col])
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLETYPE) == 1 {
		s.TABLETYPE = sdb.ToString(row[*col])
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_ENGINE) == 1 {
		if row[*col] == nil {
			s.ENGINE = nil
			} else {
				s.ENGINE = new(string)
				*s.ENGINE = sdb.ToString(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_VERSION) == 1 {
		if row[*col] == nil {
			s.VERSION = nil
			} else {
				s.VERSION = new(int)
				*s.VERSION = sdb.ToInt(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_ROWFORMAT) == 1 {
		if row[*col] == nil {
			s.ROWFORMAT = nil
			} else {
				s.ROWFORMAT = new(string)
				*s.ROWFORMAT = sdb.ToString(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLEROWS) == 1 {
		if row[*col] == nil {
			s.TABLEROWS = nil
			} else {
				s.TABLEROWS = new(uint64)
				*s.TABLEROWS = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_AVGROWLENGTH) == 1 {
		if row[*col] == nil {
			s.AVGROWLENGTH = nil
			} else {
				s.AVGROWLENGTH = new(uint64)
				*s.AVGROWLENGTH = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_DATALENGTH) == 1 {
		if row[*col] == nil {
			s.DATALENGTH = nil
			} else {
				s.DATALENGTH = new(uint64)
				*s.DATALENGTH = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_MAXDATALENGTH) == 1 {
		if row[*col] == nil {
			s.MAXDATALENGTH = nil
			} else {
				s.MAXDATALENGTH = new(uint64)
				*s.MAXDATALENGTH = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_INDEXLENGTH) == 1 {
		if row[*col] == nil {
			s.INDEXLENGTH = nil
			} else {
				s.INDEXLENGTH = new(uint64)
				*s.INDEXLENGTH = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_DATAFREE) == 1 {
		if row[*col] == nil {
			s.DATAFREE = nil
			} else {
				s.DATAFREE = new(uint64)
				*s.DATAFREE = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_AUTOINCREMENT) == 1 {
		if row[*col] == nil {
			s.AUTOINCREMENT = nil
			} else {
				s.AUTOINCREMENT = new(uint64)
				*s.AUTOINCREMENT = sdb.ToUInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_CREATETIME) == 1 {
		s.CREATETIME = sdb.ToTime(row[*col])
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_UPDATETIME) == 1 {
		if row[*col] == nil {
			s.UPDATETIME = nil
			} else {
				s.UPDATETIME = new(time.Time)
				*s.UPDATETIME = sdb.ToTime(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_CHECKTIME) == 1 {
		if row[*col] == nil {
			s.CHECKTIME = nil
			} else {
				s.CHECKTIME = new(time.Time)
				*s.CHECKTIME = sdb.ToTime(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLECOLLATION) == 1 {
		if row[*col] == nil {
			s.TABLECOLLATION = nil
			} else {
				s.TABLECOLLATION = new(string)
				*s.TABLECOLLATION = sdb.ToString(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_CHECKSUM) == 1 {
		if row[*col] == nil {
			s.CHECKSUM = nil
			} else {
				s.CHECKSUM = new(int64)
				*s.CHECKSUM = sdb.ToInt64(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_CREATEOPTIONS) == 1 {
		if row[*col] == nil {
			s.CREATEOPTIONS = nil
			} else {
				s.CREATEOPTIONS = new(string)
				*s.CREATEOPTIONS = sdb.ToString(row[*col])
			}
		*col++
	}
    if colSet == nil || colSet.Bit(codegen.TABLES_TABLECOMMENT) == 1 {
		if row[*col] == nil {
			s.TABLECOMMENT = nil
			} else {
				s.TABLECOMMENT = new(string)
				*s.TABLECOMMENT = sdb.ToString(row[*col])
			}
		*col++
	}}

func (s *TABLESStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", TABLESQueryFields(s.colSet))
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
// QueryCustom retrieves many rows from 'information_schema.TABLES' as a slice of TABLES with 1:1 joined data.
func (s *TABLESStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.TABLES, error) {
    dto := &TABLES{}
    data := &TABLESSlice{}
    err := s.queryCustom(data, dto, stmt, args...)
    if err != nil {
        s.log.Error().Err(err).Msg("querycustom")
        return nil, err
    }
    retValues := make([]*codegen.TABLES, len(data.data))
    for i := range data.data {
        retValues[i] = &data.data[i].TABLES
    }
    return retValues, nil
}
// One retrieves a row from 'information_schema.TABLES' as a TABLES with 1:1 joined data.
func (s *TABLESStore) One(args ...interface{}) (*codegen.TABLES, error) {
    data := &TABLES{}

 	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
        s.log.Error().Err(err).Msg("query one")
        return nil, err
	}
	return &data.TABLES, nil
}
// Query retrieves many rows from 'information_schema.TABLES' as a slice of TABLES with 1:1 joined data.
func (s *TABLESStore) Query(args ...interface{}) ([]*codegen.TABLES, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}



// tABLESUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *TABLESStore) tABLESUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECATALOG) == 1 {
			    upsert = append(upsert, "table_catalog = VALUES(table_catalog)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLESCHEMA) == 1 {
			    upsert = append(upsert, "table_schema = VALUES(table_schema)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLENAME) == 1 {
			    upsert = append(upsert, "table_name = VALUES(table_name)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLETYPE) == 1 {
			    upsert = append(upsert, "table_type = VALUES(table_type)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ENGINE) == 1 {
			    upsert = append(upsert, "engine = VALUES(engine)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_VERSION) == 1 {
			    upsert = append(upsert, "version = VALUES(version)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ROWFORMAT) == 1 {
			    upsert = append(upsert, "row_format = VALUES(row_format)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLEROWS) == 1 {
			    upsert = append(upsert, "table_rows = VALUES(table_rows)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AVGROWLENGTH) == 1 {
			    upsert = append(upsert, "avg_row_length = VALUES(avg_row_length)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATALENGTH) == 1 {
			    upsert = append(upsert, "data_length = VALUES(data_length)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_MAXDATALENGTH) == 1 {
			    upsert = append(upsert, "max_data_length = VALUES(max_data_length)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_INDEXLENGTH) == 1 {
			    upsert = append(upsert, "index_length = VALUES(index_length)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATAFREE) == 1 {
			    upsert = append(upsert, "data_free = VALUES(data_free)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AUTOINCREMENT) == 1 {
			    upsert = append(upsert, "auto_increment = VALUES(auto_increment)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATETIME) == 1 {
			    upsert = append(upsert, "create_time = VALUES(create_time)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_UPDATETIME) == 1 {
			    upsert = append(upsert, "update_time = VALUES(update_time)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKTIME) == 1 {
			    upsert = append(upsert, "check_time = VALUES(check_time)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOLLATION) == 1 {
			    upsert = append(upsert, "table_collation = VALUES(table_collation)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKSUM) == 1 {
			    upsert = append(upsert, "checksum = VALUES(checksum)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATEOPTIONS) == 1 {
			    upsert = append(upsert, "create_options = VALUES(create_options)")
			}
        	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOMMENT) == 1 {
			    upsert = append(upsert, "table_comment = VALUES(table_comment)")
			}
	sql := &sdb.UpsertStatement{}
		sql.InsertInto("information_schema.TABLES")
	sql.Columns("table_catalog","table_schema","table_name","table_type","engine","version","row_format","table_rows","avg_row_length","data_length","max_data_length","index_length","data_free","auto_increment","create_time","update_time","check_time","table_collation","checksum","create_options","table_comment",)
    sql.OnDuplicateKeyUpdate(upsert)
	return sql	
}

// Upsert executes upsert for array of TABLES
func (s *TABLESStore) Upsert(data ...*codegen.TABLES) (int64, error) {
	sql := s.tABLESUpsertStmt()
	
	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "TABLESUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the TABLES to the database.
func (s *TABLESStore) Insert(data *codegen.TABLES) error {
    var err error
    sql := sdb.NewSQLStatement()
		sql.AppendRaw("INSERT INTO information_schema.TABLES (")
	fields := TABLESQueryFields(s.colSet)
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
	    s.log.Trace().Str("fn", "information_schema.TABLES.Insert").Str("stmt", sql.String()).Str("TABLECATALOG", data.TABLECATALOG).Str("TABLESCHEMA", data.TABLESCHEMA).Str("TABLENAME", data.TABLENAME).Str("TABLETYPE", data.TABLETYPE).Str("ENGINE", logString(data.ENGINE)).Int("VERSION", logInt(data.VERSION)).Str("ROWFORMAT", logString(data.ROWFORMAT)).Uint64("TABLEROWS", logUInt64(data.TABLEROWS)).Uint64("AVGROWLENGTH", logUInt64(data.AVGROWLENGTH)).Uint64("DATALENGTH", logUInt64(data.DATALENGTH)).Uint64("MAXDATALENGTH", logUInt64(data.MAXDATALENGTH)).Uint64("INDEXLENGTH", logUInt64(data.INDEXLENGTH)).Uint64("DATAFREE", logUInt64(data.DATAFREE)).Uint64("AUTOINCREMENT", logUInt64(data.AUTOINCREMENT)).Time("CREATETIME", data.CREATETIME).Time("UPDATETIME", logTime(data.UPDATETIME)).Time("CHECKTIME", logTime(data.CHECKTIME)).Str("TABLECOLLATION", logString(data.TABLECOLLATION)).Int64("CHECKSUM", logInt64(data.CHECKSUM)).Str("CREATEOPTIONS", logString(data.CREATEOPTIONS)).Str("TABLECOMMENT", logString(data.TABLECOMMENT)).Msg("sql")
    }
		_, err =s.db.Exec(sql.Query(),data.TABLECATALOG,data.TABLESCHEMA,data.TABLENAME,data.TABLETYPE,data.ENGINE,data.VERSION,data.ROWFORMAT,data.TABLEROWS,data.AVGROWLENGTH,data.DATALENGTH,data.MAXDATALENGTH,data.INDEXLENGTH,data.DATAFREE,data.AUTOINCREMENT,data.CREATETIME,data.UPDATETIME,data.CHECKTIME,data.TABLECOLLATION,data.CHECKSUM,data.CREATEOPTIONS,data.TABLECOMMENT)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	return nil
}

// Update updates the TABLES in the database.
// nolint[gocyclo]
func (s *TABLESStore) Update(data *codegen.TABLES) (int64, error) {
    sql := sdb.NewSQLStatement()
    var prepend string
    args := []interface{}{}
        sql.Append("UPDATE information_schema.TABLES SET")
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECATALOG) == 1 {
            sql.AppendRaw(prepend, "table_catalog = ?")
                prepend = ","
            args = append(args, data.TABLECATALOG)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLESCHEMA) == 1 {
            sql.AppendRaw(prepend, "table_schema = ?")
                prepend = ","
            args = append(args, data.TABLESCHEMA)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLENAME) == 1 {
            sql.AppendRaw(prepend, "table_name = ?")
                prepend = ","
            args = append(args, data.TABLENAME)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLETYPE) == 1 {
            sql.AppendRaw(prepend, "table_type = ?")
                prepend = ","
            args = append(args, data.TABLETYPE)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ENGINE) == 1 {
            sql.AppendRaw(prepend, "engine = ?")
                prepend = ","
            args = append(args, data.ENGINE)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_VERSION) == 1 {
            sql.AppendRaw(prepend, "version = ?")
                prepend = ","
            args = append(args, data.VERSION)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ROWFORMAT) == 1 {
            sql.AppendRaw(prepend, "row_format = ?")
                prepend = ","
            args = append(args, data.ROWFORMAT)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLEROWS) == 1 {
            sql.AppendRaw(prepend, "table_rows = ?")
                prepend = ","
            args = append(args, data.TABLEROWS)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AVGROWLENGTH) == 1 {
            sql.AppendRaw(prepend, "avg_row_length = ?")
                prepend = ","
            args = append(args, data.AVGROWLENGTH)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATALENGTH) == 1 {
            sql.AppendRaw(prepend, "data_length = ?")
                prepend = ","
            args = append(args, data.DATALENGTH)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_MAXDATALENGTH) == 1 {
            sql.AppendRaw(prepend, "max_data_length = ?")
                prepend = ","
            args = append(args, data.MAXDATALENGTH)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_INDEXLENGTH) == 1 {
            sql.AppendRaw(prepend, "index_length = ?")
                prepend = ","
            args = append(args, data.INDEXLENGTH)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATAFREE) == 1 {
            sql.AppendRaw(prepend, "data_free = ?")
                prepend = ","
            args = append(args, data.DATAFREE)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AUTOINCREMENT) == 1 {
            sql.AppendRaw(prepend, "auto_increment = ?")
                prepend = ","
            args = append(args, data.AUTOINCREMENT)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATETIME) == 1 {
            sql.AppendRaw(prepend, "create_time = ?")
                prepend = ","
            args = append(args, data.CREATETIME)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_UPDATETIME) == 1 {
            sql.AppendRaw(prepend, "update_time = ?")
                prepend = ","
            args = append(args, data.UPDATETIME)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKTIME) == 1 {
            sql.AppendRaw(prepend, "check_time = ?")
                prepend = ","
            args = append(args, data.CHECKTIME)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOLLATION) == 1 {
            sql.AppendRaw(prepend, "table_collation = ?")
                prepend = ","
            args = append(args, data.TABLECOLLATION)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKSUM) == 1 {
            sql.AppendRaw(prepend, "checksum = ?")
                prepend = ","
            args = append(args, data.CHECKSUM)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATEOPTIONS) == 1 {
            sql.AppendRaw(prepend, "create_options = ?")
                prepend = ","
            args = append(args, data.CREATEOPTIONS)
        }
        if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOMMENT) == 1 {
            sql.AppendRaw(prepend, "table_comment = ?")
            args = append(args, data.TABLECOMMENT)
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
// Truncate deletes all rows from TABLES.
func (s *TABLESStore) Truncate() error {
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
func (s *TABLESStore) ToJSON(t *sdb.JsonBuffer, data *TABLES) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECATALOG) == 1 {
			t.JS(prepend, "table_catalog", data.TABLECATALOG)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLESCHEMA) == 1 {
			t.JS(prepend, "table_schema", data.TABLESCHEMA)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLENAME) == 1 {
			t.JS(prepend, "table_name", data.TABLENAME)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLETYPE) == 1 {
			t.JS(prepend, "table_type", data.TABLETYPE)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ENGINE) == 1 {
			t.JS(prepend, "engine", *data.ENGINE)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_VERSION) == 1 {
			t.JD(prepend, "version", *data.VERSION)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_ROWFORMAT) == 1 {
			t.JS(prepend, "row_format", *data.ROWFORMAT)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLEROWS) == 1 {
			t.JD64u(prepend, "table_rows", *data.TABLEROWS)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AVGROWLENGTH) == 1 {
			t.JD64u(prepend, "avg_row_length", *data.AVGROWLENGTH)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATALENGTH) == 1 {
			t.JD64u(prepend, "data_length", *data.DATALENGTH)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_MAXDATALENGTH) == 1 {
			t.JD64u(prepend, "max_data_length", *data.MAXDATALENGTH)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_INDEXLENGTH) == 1 {
			t.JD64u(prepend, "index_length", *data.INDEXLENGTH)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_DATAFREE) == 1 {
			t.JD64u(prepend, "data_free", *data.DATAFREE)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_AUTOINCREMENT) == 1 {
			t.JD64u(prepend, "auto_increment", *data.AUTOINCREMENT)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATETIME) == 1 {
			t.JT(prepend, "create_time", data.CREATETIME)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_UPDATETIME) == 1 {
			t.JT(prepend, "update_time", *data.UPDATETIME)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKTIME) == 1 {
			t.JT(prepend, "check_time", *data.CHECKTIME)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOLLATION) == 1 {
			t.JS(prepend, "table_collation", *data.TABLECOLLATION)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CHECKSUM) == 1 {
			t.JD64(prepend, "checksum", *data.CHECKSUM)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_CREATEOPTIONS) == 1 {
			t.JS(prepend, "create_options", *data.CREATEOPTIONS)
			prepend = ","
		}
	if s.colSet == nil || s.colSet.Bit(codegen.TABLES_TABLECOMMENT) == 1 {
			t.JS(prepend, "table_comment", *data.TABLECOMMENT)
		}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *TABLESStore) ToJSONArray(w io.Writer, data []*TABLES, name string) {
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
