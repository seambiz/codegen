package mysql

import (
	"database/sql"
	"errors"
	"io"
	"math/big"

	codegen "bitbucket.org/codegen"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

// Person represents a row from 'person'.
type Person struct {
	codegen.Person
}

// new implements Bindable.new
func (pe *Person) new() Bindable {
	return &Person{}
}

// helper struct for common query operations.
type PersonSlice struct {
	data []*Person
}

// append implements BindableSlice.append
func (pe *PersonSlice) append(d Bindable) {
	pe.data = append(pe.data, d.(*Person))
}

// constant slice for all fields of the table "Person".
// nolint[gochecknoglobals]
var personQueryFieldsAll = []string{"id", "name"}

// returns fields, that should be used.
// nolint[gocyclo]
func PersonQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return personQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Person_ID) == 1 {
		fields = append(fields, "id")
	}

	if colSet.Bit(codegen.Person_Name) == 1 {
		fields = append(fields, "name")
	}
	return fields
}

// PersonStore is used to query for 'Person' records.
type PersonStore struct {
	Store
}

// NewPersonStore return DAO Store for Person
func NewPersonStore(conn Execer) *PersonStore {
	pe := &PersonStore{}
	pe.db = conn
	pe.withJoin = true
	pe.joinType = sdb.LEFT
	pe.batch = 1000
	return pe
}

// WithoutJoins won't execute JOIN when querying for records.
func (pe *PersonStore) WithoutJoins() *PersonStore {
	pe.withJoin = false
	return pe
}

// Where sets local sql, that will be appended to SELECT.
func (pe *PersonStore) Where(sql string) *PersonStore {
	pe.where = sql
	return pe
}

// OrderBy sets local sql, that will be appended to SELECT.
func (pe *PersonStore) OrderBy(sql string) *PersonStore {
	pe.orderBy = sql
	return pe
}

// GroupBy sets local sql, that will be appended to SELECT.
func (pe *PersonStore) GroupBy(sql string) *PersonStore {
	pe.groupBy = sql
	return pe
}

// Limit result set size
func (pe *PersonStore) Limit(n int) *PersonStore {
	pe.limit = n
	return pe
}

// Offset used, if a limit is provided
func (pe *PersonStore) Offset(n int) *PersonStore {
	pe.offset = n
	return pe
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (pe *PersonStore) JoinType(jt string) *PersonStore {
	pe.joinType = jt
	return pe
}

// Columns sets bits for specific columns.
func (pe *PersonStore) Columns(cols ...int) *PersonStore {
	pe.Store.Columns(cols...)
	return pe
}

// nolint[gocyclo]
func (pe *Person) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindFakeBenchmarkPerson(&pe.Person, row, withJoin, colSet, col)
}

func BindFakeBenchmarkPerson(pe *codegen.Person, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Person_ID) == 1 {
		pe.ID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Person_Name) == 1 {
		pe.Name = sdb.ToString(row[*col])
		*col++
	}
}

func (pe *PersonStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", PersonQueryFields(pe.colSet))
	sql.Append(" FROM fake_benchmark.person A")
	if pe.where != "" {
		sql.Append("WHERE", pe.where)
	}
	if pe.groupBy != "" {
		sql.Append("GROUP BY", pe.groupBy)
	}
	if pe.orderBy != "" {
		sql.Append("ORDER BY", pe.orderBy)
	}
	if pe.limit > 0 {
		sql.AppendRaw("LIMIT ", pe.limit)
		if pe.offset > 0 {
			sql.AppendRaw(",", pe.offset)
		}
	}
	return sql
}

// QueryCustom retrieves many rows from 'fake_benchmark.person' as a slice of Person with 1:1 joined data.
func (pe *PersonStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Person, error) {
	dto := &Person{}
	data := &PersonSlice{}
	err := pe.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Person, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Person
	}
	return retValues, nil
}

// One retrieves a row from 'fake_benchmark.person' as a Person with 1:1 joined data.
func (pe *PersonStore) One(args ...interface{}) (*codegen.Person, error) {
	data := &Person{}

	err := pe.one(data, pe.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Person, nil
}

// Query retrieves many rows from 'fake_benchmark.person' as a slice of Person with 1:1 joined data.
func (pe *PersonStore) Query(args ...interface{}) ([]*codegen.Person, error) {
	stmt := pe.selectStatement()
	return pe.QueryCustom(stmt.Query(), args...)
}

// EagerFetch Pets eagerly fetches N records from referenced table 'pet'.
func (pe *PersonStore) EagerFetchPets(data []*codegen.Person) error {
	where := sdb.NewSQLStatement()
	where.AppendRaw("person_id IN (")
	for i, d := range data {
		if i > 0 {
			where.Append(",")
		}
		where.AppendInt(d.ID)
	}
	where.Append(")")

	details, err := NewPetStore(pe.db).Where(where.Query()).OrderBy("A.person_id DESC, A.id DESC").Query()
	if err != nil {
		log.Error().Err(err).Msg("fetch details")
		return err
	}
	for i := range data {
		for j := len(details) - 1; j >= 0; j-- {
			if details[j].PersonID == data[i].ID {
				data[i].Pets = append(data[i].Pets, details[j])
				details = append(details[:j], details[j+1:]...)
			}
		}
	}
	return nil
}

// Insert inserts the Person to the database.
func (pe *PersonStore) Insert(data *codegen.Person) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO fake_benchmark.person (")
	fields := PersonQueryFields(pe.colSet)
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
		log.Debug().Str("fn", "fake_benchmark.person.Insert").Str("stmt", sql.String()).Int("ID", data.ID).Str("Name", data.Name).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), data.ID, data.Name)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("lastinsertid")
		return err
	}

	// set primary key and existence
	data.ID = int(id)

	return nil
}

// Update updates the Person in the database.
// nolint[gocyclo]
func (pe *PersonStore) Update(data *codegen.Person) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE fake_benchmark.person SET")
	if pe.colSet == nil || pe.colSet.Bit(codegen.Person_Name) == 1 {
		sql.AppendRaw(prepend, "name = ?")
		args = append(args, data.Name)
	}
	sql.Append(" WHERE id = ?")
	args = append(args, data.ID)
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.person.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Delete deletes the Person from the database.
func (pe *PersonStore) Delete(data *codegen.Person) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.person WHERE")
	sql.Append("id = ?")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.person.Delete").Str("stmt", sql.String()).Int("ID", data.ID).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), data.ID)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteSlice delets all slice element from the database.
func (pe *PersonStore) DeleteSlice(data []*codegen.Person) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.person WHERE")
	sql.AppendRaw("id IN (")
	for i := range data {
		if i > 0 {
			sql.AppendRaw(",")
		}
		sql.AppendInt(data[i].ID)
	}
	sql.Append(")")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.person.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteByQuery uses a where condition to delete entries.
func (pe *PersonStore) DeleteByQuery(args ...interface{}) (int64, error) {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.person")
	if pe.where == "" {
		return 0, errors.New("no where condition set")
	}
	sql.Append("WHERE", pe.where)
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.person.DeleteByQuery").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}

	res, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Person.
func (pe *PersonStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE fake_benchmark.person")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.person.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// OneByID retrieves a row from 'fake_benchmark.person' as a Person.
//
// Generated from index 'primary'.
// nolint[goconst]
func (pe *PersonStore) OneByID(id int) (*codegen.Person, error) {
	pe.where = "A.id = ?"
	return pe.One(id)
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (pe *PersonStore) ToJSON(t *sdb.JsonBuffer, data *Person) {
	prepend := "{"
	if pe.colSet == nil || pe.colSet.Bit(codegen.Person_ID) == 1 {
		t.JD(prepend, "id", data.ID)
		prepend = ","
	}
	if pe.colSet == nil || pe.colSet.Bit(codegen.Person_Name) == 1 {
		t.JS(prepend, "name", data.Name)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (pe *PersonStore) ToJSONArray(w io.Writer, data []*Person, name string) {
	t := sdb.NewJsonBuffer()
	t.SS(`{"`, name, `":[`)
	for i := range data {
		if i > 0 {
			t.S(",")
		}
		pe.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
