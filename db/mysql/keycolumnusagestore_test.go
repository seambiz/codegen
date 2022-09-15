package mysql

import (
	"testing"

	codegen "bitbucket.org/codegen"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

// Insert inserts the KeyColumnUsage to the database.
func TestKeyColumnUsageInsert(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("INSERT INTO information_schema.KEY_COLUMN_USAGE (constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, column_name, ordinal_position, position_in_unique_constraint, referenced_table_schema, referenced_table_name, referenced_column_name) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? )").
		WithArgs("", "", nil, "", "", "", nil, 0, nil, nil, nil, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	store := NewKeyColumnUsageStore(ctx, db)
	err = store.Insert(&codegen.KeyColumnUsage{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestKeyColumnUsageUpdate(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("UPDATE information_schema.KEY_COLUMN_USAGE SET constraint_catalog = ?,constraint_schema = ?,constraint_name = ?,table_catalog = ?,table_schema = ?,table_name = ?,column_name = ?,ordinal_position = ?,position_in_unique_constraint = ?,referenced_table_schema = ?,referenced_table_name = ?,referenced_column_name = ? WHERE ").
		WithArgs("", "", nil, "", "", "", nil, 0, nil, nil, nil, nil).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewKeyColumnUsageStore(ctx, db)
	aff, err := store.Update(&codegen.KeyColumnUsage{})
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

func TestKeyColumnUsageSelectWithoutJoin(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.constraint_catalog", "A.constraint_schema", "A.constraint_name", "A.table_catalog", "A.table_schema", "A.table_name", "A.column_name", "A.ordinal_position", "A.position_in_unique_constraint", "A.referenced_table_schema", "A.referenced_table_name", "A.referenced_column_name"}).
		AddRow("", "", nil, "", "", "", nil, 0, nil, nil, nil, nil).
		AddRow("", "", nil, "", "", "", nil, 0, nil, nil, nil, nil)

	mock.ExpectQuery("SELECT A.constraint_catalog, A.constraint_schema, A.constraint_name, A.table_catalog, A.table_schema, A.table_name, A.column_name, A.ordinal_position, A.position_in_unique_constraint, A.referenced_table_schema, A.referenced_table_name, A.referenced_column_name FROM information_schema.KEY_COLUMN_USAGE A").
		WillReturnRows(rows)
	store := NewKeyColumnUsageStore(ctx, db).WithoutJoins()
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

// ^^ END OF GENERATED BY CODEGEN. ^^
