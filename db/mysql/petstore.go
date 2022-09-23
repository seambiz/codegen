package mysql

import (
	"database/sql"
	"errors"
	"io"
	"math/big"

	codegen "github.com/seambiz/codegen"

	"github.com/seambiz/seambiz/sdb"
)

// GENERATED BY CODEGEN.

// Pet represents a row from 'pet'.
type Pet struct {
	codegen.Pet
}

// new implements Bindable.new
func (s *Pet) new() Bindable {
	return &Pet{}
}

// helper struct for common query operations.
type PetSlice struct {
	data []*Pet
}

// append implements BindableSlice.append
func (s *PetSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Pet))
}

// constant slice for all fields of the table "Pet".
// nolint[gochecknoglobals]
var petQueryFieldsAll = []string{"id", "person_id", "tag_id", "species"}

// returns fields, that should be used.
// nolint[gocyclo]
func PetQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return petQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Pet_ID) == 1 {
		fields = append(fields, "id")
	}

	if colSet.Bit(codegen.Pet_PersonID) == 1 {
		fields = append(fields, "person_id")
	}

	if colSet.Bit(codegen.Pet_TagID) == 1 {
		fields = append(fields, "tag_id")
	}

	if colSet.Bit(codegen.Pet_Species) == 1 {
		fields = append(fields, "species")
	}
	return fields
}

// PetStore is used to query for 'Pet' records.
type PetStore struct {
	Store
	ctx *codegen.Context
}

// NewPetStore return DAO Store for Pet
func NewPetStore(ctx *codegen.Context, conn Execer) *PetStore {
	s := &PetStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	s.ctx = ctx
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *PetStore) WithoutJoins() *PetStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *PetStore) Where(sql string) *PetStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *PetStore) OrderBy(sql string) *PetStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *PetStore) GroupBy(sql string) *PetStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *PetStore) Limit(n int) *PetStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *PetStore) Offset(n int) *PetStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *PetStore) JoinType(jt string) *PetStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *PetStore) Columns(cols ...int) *PetStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *PetStore) SetBits(colSet *big.Int) *PetStore {
	s.colSet = colSet
	return s
}

func (s *Pet) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Pet_ID) == 1 {
		s.ID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Pet_PersonID) == 1 {
		s.PersonID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Pet_TagID) == 1 {
		s.TagID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Pet_Species) == 1 {
		s.Species = sdb.ToString(row[*col])
		*col++
	}
	if withJoin {
		belongsto := &Person{}
		s.BelongsTo = &belongsto.Person
		belongsto.bind(row, false, colSet, col)
		hastag := &Tag{}
		s.HasTag = &hastag.Tag
		hastag.bind(row, false, colSet, col)
	}
}

func (s *PetStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", PetQueryFields(s.colSet))
	if s.withJoin {
		sql.Fields(", ", "B", PersonQueryFields(s.colSet))
		sql.Fields(", ", "C", TagQueryFields(s.colSet))
		sql.Append(" FROM fake_benchmark.pet A ")
		sql.Append(s.joinType, " JOIN fake_benchmark.person B ON (A.person_id = B.id) ")
		sql.Append(s.joinType, " JOIN fake_benchmark.tag C ON (A.tag_id = C.id) ")
	} else {
		sql.Append(" FROM fake_benchmark.pet A ")
	}
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

// QueryCustom retrieves many rows from 'fake_benchmark.pet' as a slice of Pet with 1:1 joined data.
func (s *PetStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Pet, error) {
	dto := &Pet{}
	data := &PetSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Pet, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Pet
	}
	return retValues, nil
}

// One retrieves a row from 'fake_benchmark.pet' as a Pet with 1:1 joined data.
func (s *PetStore) One(args ...interface{}) (*codegen.Pet, error) {
	data := &Pet{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Pet, nil
}

// Query retrieves many rows from 'fake_benchmark.pet' as a slice of Pet with 1:1 joined data.
func (s *PetStore) Query(args ...interface{}) ([]*codegen.Pet, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// petUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *PetStore) petUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_PersonID) == 1 {
		upsert = append(upsert, "person_id = VALUES(person_id)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_TagID) == 1 {
		upsert = append(upsert, "tag_id = VALUES(tag_id)")
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_Species) == 1 {
		upsert = append(upsert, "species = VALUES(species)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("fake_benchmark.pet")
	sql.Columns("id", "person_id", "tag_id", "species")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Pet
func (s *PetStore) Upsert(data ...*codegen.Pet) (int64, error) {
	sql := s.petUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "PetUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the Pet to the database.
func (s *PetStore) Insert(data *codegen.Pet) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO fake_benchmark.pet (")
	fields := PetQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "fake_benchmark.pet.Insert").Str("stmt", sql.String()).Int("ID", data.ID).Int("PersonID", data.PersonID).Int("TagID", data.TagID).Str("Species", data.Species).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), data.ID, data.PersonID, data.TagID, data.Species)
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

// Update updates the Pet in the database.
// nolint[gocyclo]
func (s *PetStore) Update(data *codegen.Pet) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE fake_benchmark.pet SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_PersonID) == 1 {
		sql.AppendRaw(prepend, "person_id = ?")
		prepend = ","
		args = append(args, data.PersonID)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_TagID) == 1 {
		sql.AppendRaw(prepend, "tag_id = ?")
		prepend = ","
		args = append(args, data.TagID)
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_Species) == 1 {
		sql.AppendRaw(prepend, "species = ?")
		args = append(args, data.Species)
	}
	sql.Append(" WHERE id = ?")
	args = append(args, data.ID)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.pet.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Delete deletes the Pet from the database.
func (s *PetStore) Delete(data *codegen.Pet) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.pet WHERE")
	sql.Append("id = ?")

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.pet.Delete").Str("stmt", sql.String()).Int("ID", data.ID).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), data.ID)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteSlice delets all slice element from the database.
func (s *PetStore) DeleteSlice(data []*codegen.Pet) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.pet WHERE")
	sql.AppendRaw("id IN (")
	for i := range data {
		if i > 0 {
			sql.AppendRaw(",")
		}
		sql.AppendInt(data[i].ID)
	}
	sql.Append(")")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.pet.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteByQuery uses a where condition to delete entries.
func (s *PetStore) DeleteByQuery(args ...interface{}) (int64, error) {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.pet")
	if s.where == "" {
		return 0, errors.New("no where condition set")
	}
	sql.Append("WHERE", s.where)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.pet.DeleteByQuery").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}

	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Pet.
func (s *PetStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE fake_benchmark.pet")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.pet.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// OneByID retrieves a row from 'fake_benchmark.pet' as a Pet.
//
// Generated from index 'primary'.
// nolint[goconst]
func (s *PetStore) OneByID(id int) (*codegen.Pet, error) {
	s.where = "A.id = ?"
	return s.One(id)
}

// QueryByPersonID retrieves multiple rows from 'fake_benchmark.pet' as a slice of Pet.
//
// Generated from index 'fk_pet_person_idx'.
// nolint[goconst]
func (s *PetStore) QueryByPersonID(personid int) ([]*codegen.Pet, error) {
	s.where = "A.person_id = ?"
	return s.Query(personid)
}

// QueryByTagID retrieves multiple rows from 'fake_benchmark.pet' as a slice of Pet.
//
// Generated from index 'fk_pet_tag_idx'.
// nolint[goconst]
func (s *PetStore) QueryByTagID(tagid int) ([]*codegen.Pet, error) {
	s.where = "A.tag_id = ?"
	return s.Query(tagid)
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *PetStore) ToJSON(t *sdb.JsonBuffer, data *Pet) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_ID) == 1 {
		t.JD(prepend, "id", data.ID)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_PersonID) == 1 {
		t.JD(prepend, "person_id", data.PersonID)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_TagID) == 1 {
		t.JD(prepend, "tag_id", data.TagID)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Pet_Species) == 1 {
		t.JS(prepend, "species", data.Species)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *PetStore) ToJSONArray(w io.Writer, data []*Pet, name string) {
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
