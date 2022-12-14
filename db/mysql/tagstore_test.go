package mysql

import (
	"testing"

	codegen "github.com/seambiz/codegen"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

// Insert inserts the Tag to the database.
func TestTagInsert(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("INSERT INTO fake_benchmark.tag (id, name) VALUES ( ? , ? )").
		WithArgs(0, "").
		WillReturnResult(sqlmock.NewResult(1, 1))
	store := NewTagStore(ctx, db)
	err = store.Insert(&codegen.Tag{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTagUpdate(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("UPDATE fake_benchmark.tag SET name = ? WHERE id = ?").
		WithArgs("", 0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewTagStore(ctx, db)
	aff, err := store.Update(&codegen.Tag{})
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

func TestTagSelectWithoutJoin(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.id", "A.name"}).
		AddRow(0, "").
		AddRow(0, "")

	mock.ExpectQuery("SELECT A.id, A.name FROM fake_benchmark.tag A").
		WillReturnRows(rows)
	store := NewTagStore(ctx, db).WithoutJoins()
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

func TestTagDelete(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.tag WHERE id = ?").
		WithArgs(0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	aff, err := NewTagStore(ctx, db).Delete(&codegen.Tag{})
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

func TestTagDeleteSlice(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.tag WHERE id IN (0,0)").
		WillReturnResult(sqlmock.NewResult(0, 2))
	aff, err := NewTagStore(ctx, db).DeleteSlice([]*codegen.Tag{{}, {}})
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
