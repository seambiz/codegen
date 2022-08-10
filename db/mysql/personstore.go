package mysql

import (
	"database/sql"
	"errors"
	"io"
	"math/big"

	codegen "bitbucket.org/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

// Person represents a row from 'person'.
type Person struct {
	codegen.Person
}

// new implements Bindable.new
func (s *Person) new() Bindable {
	return &Person{}
}

// helper struct for common query operations.
type PersonSlice struct {
	data []*Person
}

// append implements BindableSlice.append
func (s *PersonSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Person))
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
func NewPersonStore(ctx *codegen.BaseContext, conn Execer) *PersonStore {
	s := &PersonStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *PersonStore) WithoutJoins() *PersonStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *PersonStore) Where(sql string) *PersonStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *PersonStore) OrderBy(sql string) *PersonStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *PersonStore) GroupBy(sql string) *PersonStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *PersonStore) Limit(n int) *PersonStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *PersonStore) Offset(n int) *PersonStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *PersonStore) JoinType(jt string) *PersonStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *PersonStore) Columns(cols ...int) *PersonStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *PersonStore) SetBits(colSet *big.Int) *PersonStore {
	s.colSet = colSet
	return s
}

func (s *Person) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindFakeBenchmarkPerson(&s.Person, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindFakeBenchmarkPerson(s *codegen.Person, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Person_ID) == 1 {
		s.ID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Person_Name) == 1 {
		s.Name = sdb.ToString(row[*col])
		*col++
	}
}

func (s *PersonStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", PersonQueryFields(s.colSet))
	sql.Append(" FROM fake_benchmark.person A ")
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

// QueryCustom retrieves many rows from 'fake_benchmark.person' as a slice of Person with 1:1 joined data.
func (s *PersonStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Person, error) {
	dto := &Person{}
	data := &PersonSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Person, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Person
	}
	return retValues, nil
}

// One retrieves a row from 'fake_benchmark.person' as a Person with 1:1 joined data.
func (s *PersonStore) One(args ...interface{}) (*codegen.Person, error) {
	data := &Person{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Person, nil
}

// Query retrieves many rows from 'fake_benchmark.person' as a slice of Person with 1:1 joined data.
func (s *PersonStore) Query(args ...interface{}) ([]*codegen.Person, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// EagerFetch Pets eagerly fetches N records from referenced table 'pet'.
func (s *PersonStore) EagerFetchPets(fkStore *PetStore, data []*codegen.Person) error {
	if len(data) == 0 {
		return nil
	}

	where := sdb.NewSQLStatement()
	where.AppendRaw("person_id IN (")
	for i, d := range data {
		if i > 0 {
			where.Append(",")
		}
		where.AppendInt(d.ID)
	}
	where.Append(")")

	details, err := fkStore.Where(where.Query()).OrderBy("A.person_id DESC, A.id DESC").Query()
	if err != nil {
		s.log.Error().Err(err).Msg("fetch details")
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

// personUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *PersonStore) personUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Person_Name) == 1 {
		upsert = append(upsert, "name = VALUES(name)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("fake_benchmark.person")
	sql.Columns("id", "name")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Person
func (s *PersonStore) Upsert(data ...*codegen.Person) (int64, error) {
	sql := s.personUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "PersonUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the Person to the database.
func (s *PersonStore) Insert(data *codegen.Person) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO fake_benchmark.person (")
	fields := PersonQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "fake_benchmark.person.Insert").Str("stmt", sql.String()).Int("ID", data.ID).Str("Name", data.Name).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), data.ID, data.Name)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		s.log.Error().Err(err).Msg("lastinsertid")
		return err
	}

	// set primary key and existence
	data.ID = int(id)

	return nil
}

// Update updates the Person in the database.
// nolint[gocyclo]
func (s *PersonStore) Update(data *codegen.Person) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE fake_benchmark.person SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Person_Name) == 1 {
		sql.AppendRaw(prepend, "name = ?")
		args = append(args, data.Name)
	}
	sql.Append(" WHERE id = ?")
	args = append(args, data.ID)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.person.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Delete deletes the Person from the database.
func (s *PersonStore) Delete(data *codegen.Person) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.person WHERE")
	sql.Append("id = ?")

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.person.Delete").Str("stmt", sql.String()).Int("ID", data.ID).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), data.ID)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteSlice delets all slice element from the database.
func (s *PersonStore) DeleteSlice(data []*codegen.Person) (int64, error) {
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
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.person.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteByQuery uses a where condition to delete entries.
func (s *PersonStore) DeleteByQuery(args ...interface{}) (int64, error) {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.person")
	if s.where == "" {
		return 0, errors.New("no where condition set")
	}
	sql.Append("WHERE", s.where)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.person.DeleteByQuery").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}

	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Person.
func (s *PersonStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE fake_benchmark.person")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.person.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// OneByID retrieves a row from 'fake_benchmark.person' as a Person.
//
// Generated from index 'primary'.
// nolint[goconst]
func (s *PersonStore) OneByID(id int) (*codegen.Person, error) {
	s.where = "A.id = ?"
	return s.One(id)
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *PersonStore) ToJSON(t *sdb.JsonBuffer, data *Person) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Person_ID) == 1 {
		t.JD(prepend, "id", data.ID)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Person_Name) == 1 {
		t.JS(prepend, "name", data.Name)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *PersonStore) ToJSONArray(w io.Writer, data []*Person, name string) {
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

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
