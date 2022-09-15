package mysql

import (
	"testing"

	codegen "bitbucket.org/codegen"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

// Insert inserts the Pet to the database.
func TestPetInsert(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("INSERT INTO fake_benchmark.pet (id, person_id, tag_id, species) VALUES ( ? , ? , ? , ? )").
		WithArgs(0, 0, 0, "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	store := NewPetStore(ctx, db)
	err = store.Insert(&codegen.Pet{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPetUpdate(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("UPDATE fake_benchmark.pet SET person_id = ?,tag_id = ?,species = ? WHERE id = ?").
		WithArgs(0, 0, "", 0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewPetStore(ctx, db)
	aff, err := store.Update(&codegen.Pet{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if aff != 1 {
		t.Errorf("a single row should be affected: %d", aff)
	}
}

func TestPetSelectWithoutJoin(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.id", "A.person_id", "A.tag_id", "A.species"}).
		AddRow(0, 0, 0, "").
		AddRow(0, 0, 0, "")

	mock.ExpectQuery("SELECT A.id, A.person_id, A.tag_id, A.species FROM fake_benchmark.pet A").
		WillReturnRows(rows)
	store := NewPetStore(ctx, db).WithoutJoins()
	data, err := store.Query()
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(data) != 2 {
		t.Errorf("number of rows != 2: %d", len(data))
	}
}

func TestPetSelectJoin(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.id", "A.person_id", "A.tag_id", "A.species", "B.id", "B.name", "C.id", "C.name"}).
		AddRow(0, 0, 0, "", 1, "", 1, "").
		AddRow(0, 0, 0, "", 1, "", 1, "")

	mock.ExpectQuery("SELECT A.id, A.person_id, A.tag_id, A.species, B.id, B.name, C.id, C.name FROM fake_benchmark.pet A LEFT JOIN fake_benchmark.person B ON (A.person_id = B.id) LEFT JOIN fake_benchmark.tag C ON (A.tag_id = C.id)").
		WillReturnRows(rows)
	store := NewPetStore(ctx, db)
	data, err := store.Query()
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(data) != 2 {
		t.Errorf("number of rows != 2: %d", len(data))
	}

	if data[0].BelongsTo == nil {
		t.Errorf("join did not work: %v", data[0].BelongsTo)
	}

	if data[0].HasTag == nil {
		t.Errorf("join did not work: %v", data[0].HasTag)
	}
}

func TestPetDelete(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.pet WHERE id = ?").
		WithArgs(0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	aff, err := NewPetStore(ctx, db).Delete(&codegen.Pet{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if aff != 1 {
		t.Errorf("a single row should be affected: %d", aff)
	}
}

func TestPetDeleteSlice(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.pet WHERE id IN (0,0)").
		WillReturnResult(sqlmock.NewResult(0, 2))
	aff, err := NewPetStore(ctx, db).DeleteSlice([]*codegen.Pet{{}, {}})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	if aff != 2 {
		t.Errorf("two rows should be affected: %d", aff)
	}
}

// ^^ END OF GENERATED BY CODEGEN. ^^
