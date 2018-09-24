// Package tests contains the types for schema 'codegen'.
package tests

import (
	"database/sql"
	"time"

	"bitbucket.org/seambiz/seambiz/sdb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// LEFT Join
	LEFT = "LEFT"
)

/*

	t.ID, _ = strconv.Atoi(string(row[0]))
	t.Tinybool = toBool(row[1])
	t.Smallint, _ = strconv.Atoi(string(row[2]))
	t.Mediumint, _ = strconv.Atoi(string(row[3]))
	t.Int, _ = strconv.Atoi(string(row[4]))
	t.Integer, _ = strconv.Atoi(string(row[5]))
	t.Bigint, _ = strconv.ParseInt(string(row[6]), 10, 64)
	t.Utinyint = toUInt(row[7])
	t.Usmallint = toUInt(row[8])
	t.Umediumint = toUInt(row[9])
	t.UInt = toUInt(row[10])
	t.UInteger = toUInt(row[11])
	t.Ubigint, _ = strconv.ParseUint(string(row[12]), 0, 0)
	t.Float, _ = strconv.ParseFloat(string(row[13]), 64)
	t.Double, _ = strconv.ParseFloat(string(row[14]), 64)
	t.Decimal, _ = strconv.ParseFloat(string(row[15]), 64)
	t.Numeric, _ = strconv.ParseFloat(string(row[16]), 64)
	t.Bit = toBool(row[17])
	t.Year, _ = strconv.Atoi(string(row[18]))
	t.Date, _ = time.ParseInLocation("2006-01-02", string(row[19]), time.Local)
	t.Time, _ = time.ParseInLocation("15:04:05", string(row[20]), time.Local)
	t.Datetime, _ = time.ParseInLocation("2006-01-02 15:04:05", string(row[21]), time.Local)
	t.Timestamp, _ = time.ParseInLocation("2006-01-02 15:04:05", string(row[22]), time.Local)
	t.Char = string(row[23])
	t.Varchar = string(row[24])
	t.Tinytext = string(row[25])
	t.Text = string(row[26])
	t.Mediumtext = string(row[27])
	t.Longtext = string(row[28])
	t.Binary = row[29]
	t.Varbinary = row[30]
	t.Tinyblob = row[31]
	t.Blob = row[32]
	t.Mediumblob = row[33]
	t.Longblob = row[34]

*/

// seems working!!!
// GENERATED BY CODEGEN. DO NOT EDIT.
var testsQueryFields = []string{"id", "tinybool", "smallint", "mediumint", "int", "integer", "bigint", "utinyint", "usmallint", "umediumint", "uint", "uinteger", "ubigint", "float", "double", "decimal", "numeric", "bit", "year", "date", "time", "datetime", "timestamp", "char", "varchar", "tinytext", "text", "mediumtext", "longtext", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob"}

// Tests represents a row from 'codegen.tests'.
type Tests struct {
	ID         int       `json:"id" db:"id"`
	Tinybool   bool      `json:"tinybool" db:"tinybool"`
	Smallint   int       `json:"smallint" db:"smallint"`
	Mediumint  int       `json:"mediumint" db:"mediumint"`
	Int        int       `json:"int" db:"int"`
	Integer    int       `json:"integer" db:"integer"`
	Bigint     int64     `json:"bigint" db:"bigint"`
	Utinyint   uint      `json:"utinyint" db:"utinyint"`
	Usmallint  uint      `json:"usmallint" db:"usmallint"`
	Umediumint uint      `json:"umediumint" db:"umediumint"`
	UInt       uint      `json:"uint" db:"uint"`
	UInteger   uint      `json:"uinteger" db:"uinteger"`
	Ubigint    uint64    `json:"ubigint" db:"ubigint"`
	Float      float32   `json:"float" db:"float"`
	Double     float64   `json:"double" db:"double"`
	Decimal    float64   `json:"decimal" db:"decimal"`
	Numeric    float64   `json:"numeric" db:"numeric"`
	Bit        bool      `json:"bit" db:"bit"`
	Year       int       `json:"year" db:"year"`
	Date       time.Time `json:"date" db:"date"`
	Time       time.Time `json:"time" db:"time"`
	Datetime   time.Time `json:"datetime" db:"datetime"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
	Char       string    `json:"char" db:"char"`
	Varchar    string    `json:"varchar" db:"varchar"`
	Tinytext   string    `json:"tinytext" db:"tinytext"`
	Text       string    `json:"text" db:"text"`
	Mediumtext string    `json:"mediumtext" db:"mediumtext"`
	Longtext   string    `json:"longtext" db:"longtext"`
	Binary     []byte    `json:"binary" db:"binary"`
	Varbinary  []byte    `json:"varbinary" db:"varbinary"`
	Tinyblob   []byte    `json:"tinyblob" db:"tinyblob"`
	Blob       []byte    `json:"blob" db:"blob"`
	Mediumblob []byte    `json:"mediumblob" db:"mediumblob"`
	Longblob   []byte    `json:"longblob" db:"longblob"`
}

// IsEmpty checks if primary key fields are zero.
func (t *Tests) IsEmpty() bool {
	return t.ID == 0
}

// TestsStore is used to query for 'Tests' records.
type TestsStore struct {
	db       *sql.DB
	withJoin bool
	joinType string
	where    string
	orderby  string
}

// NewTestsStore return DAO Store for Tests
func NewTestsStore(conn *sql.DB) *TestsStore {
	t := &TestsStore{}
	t.db = conn
	t.withJoin = true
	t.joinType = LEFT
	return t
}

// WithoutJoins won't execute JOIN when querying for records.
func (t *TestsStore) WithoutJoins() *TestsStore {
	t.withJoin = false
	return t
}

// Where sets local sql, that will be appended to SELECT.
func (t *TestsStore) Where(sql string) *TestsStore {
	t.where = sql
	return t
}

// OrderBy sets local sql, that will be appended to SELECT.
func (t *TestsStore) OrderBy(sql string) *TestsStore {
	t.orderby = sql
	return t
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (t *TestsStore) JoinType(jt string) *TestsStore {
	t.joinType = jt
	return t
}
func (t *Tests) bind(row []sql.RawBytes, withJoin bool) {
	t.ID = sdb.ToInt(row[0])
	t.Tinybool = sdb.ToBool(row[1])
	t.Smallint = sdb.ToInt(row[2])
	t.Mediumint = sdb.ToInt(row[3])
	t.Int = sdb.ToInt(row[4])
	t.Integer = sdb.ToInt(row[5])
	t.Bigint = sdb.ToInt64(row[6])
	t.Utinyint = sdb.ToUInt(row[7])
	t.Usmallint = sdb.ToUInt(row[8])
	t.Umediumint = sdb.ToUInt(row[9])
	t.UInt = sdb.ToUInt(row[10])
	t.UInteger = sdb.ToUInt(row[11])
	t.Ubigint = sdb.ToUInt64(row[12])
	t.Float = sdb.ToFloat32(row[13])
	t.Double = sdb.ToFloat64(row[14])
	t.Decimal = sdb.ToFloat64(row[15])
	t.Numeric = sdb.ToFloat64(row[16])
	t.Bit = sdb.ToBool(row[17])
	t.Year = sdb.ToInt(row[18])
	t.Date = sdb.ToTime(row[19])
	t.Time = sdb.ToTime(row[20])
	t.Datetime = sdb.ToTime(row[21])
	t.Timestamp = sdb.ToTime(row[22])
	t.Char = string(row[23])
	t.Varchar = string(row[24])
	t.Tinytext = string(row[25])
	t.Text = string(row[26])
	t.Mediumtext = string(row[27])
	t.Longtext = string(row[28])
	t.Binary = row[29]
	t.Varbinary = row[30]
	t.Tinyblob = row[31]
	t.Blob = row[32]
	t.Mediumblob = row[33]
	t.Longblob = row[34]
}
func (t *TestsStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", testsQueryFields)
	if t.withJoin {
		sql.Append("FROM codegen.tests A")
	} else {
		sql.Append("FROM codegen.tests A")
	}
	if t.where != "" {
		sql.Append("WHERE", t.where)
	}
	if t.orderby != "" {
		sql.Append("ORDER BY", t.orderby)
	}
	return sql
}

// One retrieves a row from 'codegen.tests' as a Tests with possible joined data.
func (t *TestsStore) One(args ...interface{}) (*Tests, error) {
	var err error
	data := Tests{}

	stmt := t.selectStatement()
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "TestsStore.One").Str("stmt", stmt.String()).Interface("args", args).Msg("stmt")
	}
	rows, err := t.db.Query(stmt.Query(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	if rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		data.bind(values, t.withJoin)
	} else {
		return nil, sql.ErrNoRows
	}
	return &data, err
}

// Query retrieves many rows from 'codegen.tests' as a slice of Tests with possible joined data.
func (t *TestsStore) Query(args ...interface{}) ([]*Tests, error) {
	stmt := t.selectStatement()
	return t.QueryCustom(stmt.Query(), args...)
}

// QueryCustom retrieves many rows from 'codegen.tests' as a slice of Tests with possible joined data.
func (t *TestsStore) QueryCustom(stmt string, args ...interface{}) ([]*Tests, error) {
	var err error
	res := []*Tests{}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "TestsStore.Query").Str("stmt", stmt).Interface("args", args)).Msg("stmt")
	}
	rows, err := t.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		data := Tests{}
		data.bind(values, t.withJoin)
		res = append(res, &data)
	}
	return res, err
}

// Delete deletes the Tests from the database.
func (t *TestsStore) Delete(data *Tests) error {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM codegen.tests WHERE")
	sql.Append("id = ?")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "codegen.tests.Delete", zap.String("stmt", sql.String()), zap.Int("ID", data.ID))
	}
	_, err = t.db.Exec(sql.Query(),
		data.ID)
	if err != nil {
		log.Error().Err(err)
	}

	return err
}

// Insert inserts the Tests to the database.
func (t *TestsStore) Insert(data *Tests) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO codegen.tests (")
	sql.Fields("", "", testsQueryFields)
	sql.Append(") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "codegen.tests.Insert", zap.String("stmt", sql.String()), zap.Int("ID", data.ID), zap.Bool("Tinybool", data.Tinybool), zap.Int("Smallint", data.Smallint), zap.Int("Mediumint", data.Mediumint), zap.Int("Int", data.Int), zap.Int("Integer", data.Integer), zap.Int64("Bigint", data.Bigint), zap.Uint("Utinyint", data.Utinyint), zap.Uint("Usmallint", data.Usmallint), zap.Uint("Umediumint", data.Umediumint), zap.Uint("UInt", data.UInt), zap.Uint("UInteger", data.UInteger), zap.Uint64("Ubigint", data.Ubigint), zap.Float32("Float", data.Float), zap.Float64("Double", data.Double), zap.Float64("Decimal", data.Decimal), zap.Float64("Numeric", data.Numeric), zap.Bool("Bit", data.Bit), zap.Int("Year", data.Year), zap.Time("Date", data.Date), zap.Time("Time", data.Time), zap.Time("Datetime", data.Datetime), zap.Time("Timestamp", data.Timestamp), zap.String("Char", data.Char), zap.String("Varchar", data.Varchar), zap.String("Tinytext", data.Tinytext), zap.String("Text", data.Text), zap.String("Mediumtext", data.Mediumtext), zap.String("Longtext", data.Longtext), zap.ByteString("Binary", data.Binary), zap.ByteString("Varbinary", data.Varbinary), zap.ByteString("Tinyblob", data.Tinyblob), zap.ByteString("Blob", data.Blob), zap.ByteString("Mediumblob", data.Mediumblob), zap.ByteString("Longblob", data.Longblob))
	}
	res, err := t.db.Exec(sql.Query(), data.ID, data.Tinybool, data.Smallint, data.Mediumint, data.Int, data.Integer, data.Bigint, data.Utinyint, data.Usmallint, data.Umediumint, data.UInt, data.UInteger, data.Ubigint, data.Float, data.Double, data.Decimal, data.Numeric, data.Bit, data.Year, data.Date, data.Time, data.Datetime, data.Timestamp, data.Char, data.Varchar, data.Tinytext, data.Text, data.Mediumtext, data.Longtext, data.Binary, data.Varbinary, data.Tinyblob, data.Blob, data.Mediumblob, data.Longblob)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	// set primary key and existence
	data.ID = int(id)

	return nil
}

// Truncate deletes all rows from Tests.
func (t *TestsStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE codegen.tests")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "codegen.tests.Truncate", zap.String("stmt", sql.String()))
	}
	_, err := t.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err)
	}
	return err
}

// Update updates the Tests in the database.
func (t *TestsStore) Update(data *Tests) (int64, error) {
	sql := sdb.NewSQLStatement()
	sql.Append("UPDATE codegen.tests SET")
	sql.Append("tinybool = ?, smallint = ?, mediumint = ?, int = ?, integer = ?, bigint = ?, utinyint = ?, usmallint = ?, umediumint = ?, uint = ?, uinteger = ?, ubigint = ?, float = ?, double = ?, decimal = ?, numeric = ?, bit = ?, year = ?, date = ?, time = ?, datetime = ?, timestamp = ?, char = ?, varchar = ?, tinytext = ?, text = ?, mediumtext = ?, longtext = ?, binary = ?, varbinary = ?, tinyblob = ?, blob = ?, mediumblob = ?, longblob = ?")
	sql.Append("WHERE id = ?")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "codegen.tests.Update", zap.String("stmt", sql.String()), zap.Int("ID", data.ID), zap.Bool("Tinybool", data.Tinybool), zap.Int("Smallint", data.Smallint), zap.Int("Mediumint", data.Mediumint), zap.Int("Int", data.Int), zap.Int("Integer", data.Integer), zap.Int64("Bigint", data.Bigint), zap.Uint("Utinyint", data.Utinyint), zap.Uint("Usmallint", data.Usmallint), zap.Uint("Umediumint", data.Umediumint), zap.Uint("UInt", data.UInt), zap.Uint("UInteger", data.UInteger), zap.Uint64("Ubigint", data.Ubigint), zap.Float32("Float", data.Float), zap.Float64("Double", data.Double), zap.Float64("Decimal", data.Decimal), zap.Float64("Numeric", data.Numeric), zap.Bool("Bit", data.Bit), zap.Int("Year", data.Year), zap.Time("Date", data.Date), zap.Time("Time", data.Time), zap.Time("Datetime", data.Datetime), zap.Time("Timestamp", data.Timestamp), zap.String("Char", data.Char), zap.String("Varchar", data.Varchar), zap.String("Tinytext", data.Tinytext), zap.String("Text", data.Text), zap.String("Mediumtext", data.Mediumtext), zap.String("Longtext", data.Longtext), zap.ByteString("Binary", data.Binary), zap.ByteString("Varbinary", data.Varbinary), zap.ByteString("Tinyblob", data.Tinyblob), zap.ByteString("Blob", data.Blob), zap.ByteString("Mediumblob", data.Mediumblob), zap.ByteString("Longblob", data.Longblob))
	}
	res, err := t.db.Exec(sql.Query(), data.Tinybool, data.Smallint, data.Mediumint, data.Int, data.Integer, data.Bigint, data.Utinyint, data.Usmallint, data.Umediumint, data.UInt, data.UInteger, data.Ubigint, data.Float, data.Double, data.Decimal, data.Numeric, data.Bit, data.Year, data.Date, data.Time, data.Datetime, data.Timestamp, data.Char, data.Varchar, data.Tinytext, data.Text, data.Mediumtext, data.Longtext, data.Binary, data.Varbinary, data.Tinyblob, data.Blob, data.Mediumblob, data.Longblob, data.ID)
	if err != nil {
		log.Error().Err(err)
		return 0, err
	}
	return res.RowsAffected()
}

// testsUpsertStmt helper for generating Upserts general statement
func testsUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{
		"tinybool = VALUES(tinybool)",
		"smallint = VALUES(smallint)",
		"mediumint = VALUES(mediumint)",
		"int = VALUES(int)",
		"integer = VALUES(integer)",
		"bigint = VALUES(bigint)",
		"utinyint = VALUES(utinyint)",
		"usmallint = VALUES(usmallint)",
		"umediumint = VALUES(umediumint)",
		"uint = VALUES(uint)",
		"uinteger = VALUES(uinteger)",
		"ubigint = VALUES(ubigint)",
		"float = VALUES(float)",
		"double = VALUES(double)",
		"decimal = VALUES(decimal)",
		"numeric = VALUES(numeric)",
		"bit = VALUES(bit)",
		"year = VALUES(year)",
		"date = VALUES(date)",
		"time = VALUES(time)",
		"datetime = VALUES(datetime)",
		"timestamp = VALUES(timestamp)",
		"char = VALUES(char)",
		"varchar = VALUES(varchar)",
		"tinytext = VALUES(tinytext)",
		"text = VALUES(text)",
		"mediumtext = VALUES(mediumtext)",
		"longtext = VALUES(longtext)",
		"binary = VALUES(binary)",
		"varbinary = VALUES(varbinary)",
		"tinyblob = VALUES(tinyblob)",
		"blob = VALUES(blob)",
		"mediumblob = VALUES(mediumblob)",
		"longblob = VALUES(longblob)",
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("codegen.tests")
	sql.Columns("id", "tinybool", "smallint", "mediumint", "int", "integer", "bigint", "utinyint", "usmallint", "umediumint", "uint", "uinteger", "ubigint", "float", "double", "decimal", "numeric", "bit", "year", "date", "time", "datetime", "timestamp", "char", "varchar", "tinytext", "text", "mediumtext", "longtext", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// UpsertOne inserts the Tests to the database.
func (t *TestsStore) UpsertOne(data *Tests) error {
	return t.Upsert([]*Tests{data})
}

// Upsert executes upsert for array of Tests
func (t *TestsStore) Upsert(data []*Tests) error {
	sql := testsUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "TestsUpsert", zap.String("stmt", sql.String()))
	}
	_, err := t.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
// yes it does!!!
