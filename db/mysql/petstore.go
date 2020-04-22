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

// Pet represents a row from 'pet'.
type Pet struct {
	codegen.Pet
}

// new implements Bindable.new
func (pe *Pet) new() Bindable {
	return &Pet{}
}

// helper struct for common query operations.
type PetSlice struct {
	data []*Pet
}

// append implements BindableSlice.append
func (pe *PetSlice) append(d Bindable) {
	pe.data = append(pe.data, d.(*Pet))
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
	if colSet.Bit(Pet_ID) == 1 {
		fields = append(fields, "id")
	}

	if colSet.Bit(Pet_PersonID) == 1 {
		fields = append(fields, "person_id")
	}

	if colSet.Bit(Pet_TagID) == 1 {
		fields = append(fields, "tag_id")
	}

	if colSet.Bit(Pet_Species) == 1 {
		fields = append(fields, "species")
	}
	return fields
}

// PetStore is used to query for 'Pet' records.
type PetStore struct {
	Store
}

// NewPetStore return DAO Store for Pet
func NewPetStore(conn Execer) *PetStore {
	pe := &PetStore{}
	pe.db = conn
	pe.withJoin = true
	pe.joinType = sdb.LEFT
	pe.batch = 1000
	return pe
}

// WithoutJoins won't execute JOIN when querying for records.
func (pe *PetStore) WithoutJoins() *PetStore {
	pe.withJoin = false
	return pe
}

// Where sets local sql, that will be appended to SELECT.
func (pe *PetStore) Where(sql string) *PetStore {
	pe.where = sql
	return pe
}

// OrderBy sets local sql, that will be appended to SELECT.
func (pe *PetStore) OrderBy(sql string) *PetStore {
	pe.orderBy = sql
	return pe
}

// GroupBy sets local sql, that will be appended to SELECT.
func (pe *PetStore) GroupBy(sql string) *PetStore {
	pe.groupBy = sql
	return pe
}

// Limit result set size
func (pe *PetStore) Limit(n int) *PetStore {
	pe.limit = n
	return pe
}

// Offset used, if a limit is provided
func (pe *PetStore) Offset(n int) *PetStore {
	pe.offset = n
	return pe
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (pe *PetStore) JoinType(jt string) *PetStore {
	pe.joinType = jt
	return pe
}

// Columns sets bits for specific columns.
func (pe *PetStore) Columns(cols ...int) *PetStore {
	pe.Store.Columns(cols...)
	return pe
}

// nolint[gocyclo]
func (pe *Pet) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindFakeBenchmarkPet(&pe.Pet, row, withJoin, colSet, col)
}

func BindFakeBenchmarkPet(pe *codegen.Pet, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(Pet_ID) == 1 {
		pe.ID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Pet_PersonID) == 1 {
		pe.PersonID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Pet_TagID) == 1 {
		pe.TagID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(Pet_Species) == 1 {
		pe.Species = sdb.ToString(row[*col])
		*col++
	}
	if withJoin {
		belongsto := &Person{}
		pe.BelongsTo = &belongsto.Person
		belongsto.bind(row, false, colSet, col)
		hastag := &Tag{}
		pe.HasTag = &hastag.Tag
		hastag.bind(row, false, colSet, col)
	}
}

func (pe *PetStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", PetQueryFields(pe.colSet))
	if pe.withJoin {
		sql.Fields(", ", "B", PersonQueryFields(pe.colSet))
		sql.Fields(", ", "C", TagQueryFields(pe.colSet))
		sql.Append(" FROM fake_benchmark.pet A")
		sql.Append(pe.joinType, " JOIN fake_benchmark.person B ON (A.person_id = B.id)")
		sql.Append(pe.joinType, " JOIN fake_benchmark.tag C ON (A.tag_id = C.id)")
	} else {
		sql.Append(" FROM fake_benchmark.pet A")
	}
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

// QueryCustom retrieves many rows from 'fake_benchmark.pet' as a slice of Pet with 1:1 joined data.
func (pe *PetStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Pet, error) {
	dto := &Pet{}
	data := &PetSlice{}
	err := pe.queryCustom(data, dto, stmt, args...)
	if err != nil {
		log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Pet, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Pet
	}
	return retValues, nil
}

// One retrieves a row from 'fake_benchmark.pet' as a Pet with 1:1 joined data.
func (pe *PetStore) One(args ...interface{}) (*codegen.Pet, error) {
	data := &Pet{}

	err := pe.one(data, pe.selectStatement(), args...)
	if err != nil {
		log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Pet, nil
}

// Query retrieves many rows from 'fake_benchmark.pet' as a slice of Pet with 1:1 joined data.
func (pe *PetStore) Query(args ...interface{}) ([]*codegen.Pet, error) {
	stmt := pe.selectStatement()
	return pe.QueryCustom(stmt.Query(), args...)
}

// petUpsertStmt helper for generating Upsert statement.
// nolint[gocyclo]
func (pe *PetStore) petUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if pe.colSet == nil || pe.colSet.Bit(Pet_PersonID) == 1 {
		upsert = append(upsert, "person_id = VALUES(person_id)")
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_TagID) == 1 {
		upsert = append(upsert, "tag_id = VALUES(tag_id)")
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_Species) == 1 {
		upsert = append(upsert, "species = VALUES(species)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("fake_benchmark.pet")
	sql.Columns("id", "person_id", "tag_id", "species")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Pet
func (pe *PetStore) Upsert(data ...*codegen.Pet) (int64, error) {
	sql := pe.petUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "PetUpsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query())
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

// Insert inserts the Pet to the database.
func (pe *PetStore) Insert(data *codegen.Pet) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("INSERT INTO fake_benchmark.pet (")
	fields := PetQueryFields(pe.colSet)
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
		log.Debug().Str("fn", "fake_benchmark.pet.Insert").Str("stmt", sql.String()).Int("ID", data.ID).Int("PersonID", data.PersonID).Int("TagID", data.TagID).Str("Species", data.Species).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), data.ID, data.PersonID, data.TagID, data.Species)
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

// Update updates the Pet in the database.
// nolint[gocyclo]
func (pe *PetStore) Update(data *codegen.Pet) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE fake_benchmark.pet SET")
	if pe.colSet == nil || pe.colSet.Bit(Pet_PersonID) == 1 {
		sql.AppendRaw(prepend, "person_id = ?")
		prepend = ","
		args = append(args, data.PersonID)
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_TagID) == 1 {
		sql.AppendRaw(prepend, "tag_id = ?")
		prepend = ","
		args = append(args, data.TagID)
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_Species) == 1 {
		sql.AppendRaw(prepend, "species = ?")
		args = append(args, data.Species)
	}
	sql.Append(" WHERE id = ?")
	args = append(args, data.ID)
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.pet.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Delete deletes the Pet from the database.
func (pe *PetStore) Delete(data *codegen.Pet) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.pet WHERE")
	sql.Append("id = ?")

	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.pet.Delete").Str("stmt", sql.String()).Int("ID", data.ID).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query(), data.ID)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteSlice delets all slice element from the database.
func (pe *PetStore) DeleteSlice(data []*codegen.Pet) (int64, error) {
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
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.pet.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteByQuery uses a where condition to delete entries.
func (pe *PetStore) DeleteByQuery(args ...interface{}) (int64, error) {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.pet")
	if pe.where == "" {
		return 0, errors.New("no where condition set")
	}
	sql.Append("WHERE", pe.where)
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.pet.DeleteByQuery").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}

	res, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Pet.
func (pe *PetStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE fake_benchmark.pet")
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "fake_benchmark.pet.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := pe.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}
	return err
}

// OneByID retrieves a row from 'fake_benchmark.pet' as a Pet.
//
// Generated from index 'primary'.
// nolint[goconst]
func (pe *PetStore) OneByID(id int) (*codegen.Pet, error) {
	pe.where = "A.id = ?"
	return pe.One(id)
}

// QueryByPersonID retrieves multiple rows from 'fake_benchmark.pet' as a slice of Pet.
//
// Generated from index 'fk_pet_person_idx'.
// nolint[goconst]
func (pe *PetStore) QueryByPersonID(personid int) ([]*codegen.Pet, error) {
	pe.where = "A.person_id = ?"
	return pe.Query(personid)
}

// QueryByTagID retrieves multiple rows from 'fake_benchmark.pet' as a slice of Pet.
//
// Generated from index 'fk_pet_tag_idx'.
// nolint[goconst]
func (pe *PetStore) QueryByTagID(tagid int) ([]*codegen.Pet, error) {
	pe.where = "A.tag_id = ?"
	return pe.Query(tagid)
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (pe *PetStore) ToJSON(t *sdb.JsonBuffer, data *Pet) {
	prepend := "{"
	if pe.colSet == nil || pe.colSet.Bit(Pet_ID) == 1 {
		t.JD(prepend, "id", data.ID)
		prepend = ","
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_PersonID) == 1 {
		t.JD(prepend, "person_id", data.PersonID)
		prepend = ","
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_TagID) == 1 {
		t.JD(prepend, "tag_id", data.TagID)
		prepend = ","
	}
	if pe.colSet == nil || pe.colSet.Bit(Pet_Species) == 1 {
		t.JS(prepend, "species", data.Species)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (pe *PetStore) ToJSONArray(w io.Writer, data []*Pet, name string) {
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