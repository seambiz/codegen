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

// Tag represents a row from 'tag'.
type Tag struct {
	codegen.Tag
}

// new implements Bindable.new
func (s *Tag) new() Bindable {
	return &Tag{}
}

// helper struct for common query operations.
type TagSlice struct {
	data []*Tag
}

// append implements BindableSlice.append
func (s *TagSlice) append(d Bindable) {
	s.data = append(s.data, d.(*Tag))
}

// constant slice for all fields of the table "Tag".
// nolint[gochecknoglobals]
var tagQueryFieldsAll = []string{"id", "name"}

// returns fields, that should be used.
// nolint[gocyclo]
func TagQueryFields(colSet *big.Int) []string {
	if colSet == nil {
		return tagQueryFieldsAll
	}

	fields := []string{}
	if colSet.Bit(codegen.Tag_ID) == 1 {
		fields = append(fields, "id")
	}

	if colSet.Bit(codegen.Tag_Name) == 1 {
		fields = append(fields, "name")
	}
	return fields
}

// TagStore is used to query for 'Tag' records.
type TagStore struct {
	Store
}

// NewTagStore return DAO Store for Tag
func NewTagStore(ctx *codegen.BaseContext, conn Execer) *TagStore {
	s := &TagStore{}
	s.db = conn
	s.withJoin = true
	s.joinType = sdb.LEFT
	s.batch = 1000
	s.log = ctx.Log
	return s
}

// WithoutJoins won't execute JOIN when querying for records.
func (s *TagStore) WithoutJoins() *TagStore {
	s.withJoin = false
	return s
}

// Where sets local sql, that will be appended to SELECT.
func (s *TagStore) Where(sql string) *TagStore {
	s.where = sql
	return s
}

// OrderBy sets local sql, that will be appended to SELECT.
func (s *TagStore) OrderBy(sql string) *TagStore {
	s.orderBy = sql
	return s
}

// GroupBy sets local sql, that will be appended to SELECT.
func (s *TagStore) GroupBy(sql string) *TagStore {
	s.groupBy = sql
	return s
}

// Limit result set size
func (s *TagStore) Limit(n int) *TagStore {
	s.limit = n
	return s
}

// Offset used, if a limit is provided
func (s *TagStore) Offset(n int) *TagStore {
	s.offset = n
	return s
}

// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).
func (s *TagStore) JoinType(jt string) *TagStore {
	s.joinType = jt
	return s
}

// Columns sets bits for specific columns.
func (s *TagStore) Columns(cols ...int) *TagStore {
	s.Store.Columns(cols...)
	return s
}

// SetBits sets complete BitSet for use in UpdatePartial.
func (s *TagStore) SetBits(colSet *big.Int) *TagStore {
	s.colSet = colSet
	return s
}

func (s *Tag) bind(row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	BindFakeBenchmarkTag(&s.Tag, row, withJoin, colSet, col)
}

// nolint:gocyclo
func BindFakeBenchmarkTag(s *codegen.Tag, row []sql.RawBytes, withJoin bool, colSet *big.Int, col *int) {
	if colSet == nil || colSet.Bit(codegen.Tag_ID) == 1 {
		s.ID = sdb.ToInt(row[*col])
		*col++
	}
	if colSet == nil || colSet.Bit(codegen.Tag_Name) == 1 {
		s.Name = sdb.ToString(row[*col])
		*col++
	}
}

func (s *TagStore) selectStatement() *sdb.SQLStatement {
	sql := sdb.NewSQLStatement()
	sql.Append("SELECT")
	sql.Fields("", "A", TagQueryFields(s.colSet))
	sql.Append(" FROM fake_benchmark.tag A ")
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

// QueryCustom retrieves many rows from 'fake_benchmark.tag' as a slice of Tag with 1:1 joined data.
func (s *TagStore) QueryCustom(stmt string, args ...interface{}) ([]*codegen.Tag, error) {
	dto := &Tag{}
	data := &TagSlice{}
	err := s.queryCustom(data, dto, stmt, args...)
	if err != nil {
		s.log.Error().Err(err).Msg("querycustom")
		return nil, err
	}
	retValues := make([]*codegen.Tag, len(data.data))
	for i := range data.data {
		retValues[i] = &data.data[i].Tag
	}
	return retValues, nil
}

// One retrieves a row from 'fake_benchmark.tag' as a Tag with 1:1 joined data.
func (s *TagStore) One(args ...interface{}) (*codegen.Tag, error) {
	data := &Tag{}

	err := s.one(data, s.selectStatement(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("query one")
		return nil, err
	}
	return &data.Tag, nil
}

// Query retrieves many rows from 'fake_benchmark.tag' as a slice of Tag with 1:1 joined data.
func (s *TagStore) Query(args ...interface{}) ([]*codegen.Tag, error) {
	stmt := s.selectStatement()
	return s.QueryCustom(stmt.Query(), args...)
}

// tagUpsertStmt helper for generating Upsert statement.
// nolint:gocyclo
func (s *TagStore) tagUpsertStmt() *sdb.UpsertStatement {
	upsert := []string{}
	if s.colSet == nil || s.colSet.Bit(codegen.Tag_Name) == 1 {
		upsert = append(upsert, "name = VALUES(name)")
	}
	sql := &sdb.UpsertStatement{}
	sql.InsertInto("fake_benchmark.tag")
	sql.Columns("id", "name")
	sql.OnDuplicateKeyUpdate(upsert)
	return sql
}

// Upsert executes upsert for array of Tag
func (s *TagStore) Upsert(data ...*codegen.Tag) (int64, error) {
	sql := s.tagUpsertStmt()

	for _, d := range data {
		sql.Record(d)
	}

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "TagUpsert").Str("stmt", sql.String()).Msg("sql")
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

// Insert inserts the Tag to the database.
func (s *TagStore) Insert(data *codegen.Tag) error {
	var err error
	sql := sdb.NewSQLStatement()
	sql.AppendRaw("INSERT INTO fake_benchmark.tag (")
	fields := TagQueryFields(s.colSet)
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
		s.log.Trace().Str("fn", "fake_benchmark.tag.Insert").Str("stmt", sql.String()).Int("ID", data.ID).Str("Name", data.Name).Msg("sql")
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

// Update updates the Tag in the database.
// nolint[gocyclo]
func (s *TagStore) Update(data *codegen.Tag) (int64, error) {
	sql := sdb.NewSQLStatement()
	var prepend string
	args := []interface{}{}
	sql.Append("UPDATE fake_benchmark.tag SET")
	if s.colSet == nil || s.colSet.Bit(codegen.Tag_Name) == 1 {
		sql.AppendRaw(prepend, "name = ?")
		args = append(args, data.Name)
	}
	sql.Append(" WHERE id = ?")
	args = append(args, data.ID)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.tag.Update").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Delete deletes the Tag from the database.
func (s *TagStore) Delete(data *codegen.Tag) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.tag WHERE")
	sql.Append("id = ?")

	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.tag.Delete").Str("stmt", sql.String()).Int("ID", data.ID).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query(), data.ID)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteSlice delets all slice element from the database.
func (s *TagStore) DeleteSlice(data []*codegen.Tag) (int64, error) {
	var err error

	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.tag WHERE")
	sql.AppendRaw("id IN (")
	for i := range data {
		if i > 0 {
			sql.AppendRaw(",")
		}
		sql.AppendInt(data[i].ID)
	}
	sql.Append(")")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.tag.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// DeleteByQuery uses a where condition to delete entries.
func (s *TagStore) DeleteByQuery(args ...interface{}) (int64, error) {
	var err error
	sql := sdb.NewSQLStatement()
	sql.Append("DELETE FROM fake_benchmark.tag")
	if s.where == "" {
		return 0, errors.New("no where condition set")
	}
	sql.Append("WHERE", s.where)
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.tag.DeleteByQuery").Str("stmt", sql.String()).Interface("args", args).Msg("sql")
	}

	res, err := s.db.Exec(sql.Query(), args...)
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}

// Truncate deletes all rows from Tag.
func (s *TagStore) Truncate() error {
	sql := sdb.NewSQLStatement()
	sql.Append("TRUNCATE fake_benchmark.tag")
	if s.log.Trace().Enabled() {
		s.log.Trace().Str("fn", "fake_benchmark.tag.Truncate").Str("stmt", sql.String()).Msg("sql")
	}
	_, err := s.db.Exec(sql.Query())
	if err != nil {
		s.log.Error().Err(err).Msg("exec")
	}
	return err
}

// OneByID retrieves a row from 'fake_benchmark.tag' as a Tag.
//
// Generated from index 'primary'.
// nolint[goconst]
func (s *TagStore) OneByID(id int) (*codegen.Tag, error) {
	s.where = "A.id = ?"
	return s.One(id)
}

// ToJSON writes a single object to the buffer.
// nolint[gocylco]
func (s *TagStore) ToJSON(t *sdb.JsonBuffer, data *Tag) {
	prepend := "{"
	if s.colSet == nil || s.colSet.Bit(codegen.Tag_ID) == 1 {
		t.JD(prepend, "id", data.ID)
		prepend = ","
	}
	if s.colSet == nil || s.colSet.Bit(codegen.Tag_Name) == 1 {
		t.JS(prepend, "name", data.Name)
	}
	t.S(`}`)
}

// ToJSONArray writes a slice to the named array.
func (s *TagStore) ToJSONArray(w io.Writer, data []*Tag, name string) {
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
