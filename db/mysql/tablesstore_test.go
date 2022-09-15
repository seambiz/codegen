package mysql

import (
	"testing"
	"time"

	codegen "bitbucket.org/codegen"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

// Insert inserts the Tables to the database.
func TestTablesInsert(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("INSERT INTO information_schema.TABLES (table_catalog, table_schema, table_name, table_type, engine, version, row_format, table_rows, avg_row_length, data_length, max_data_length, index_length, data_free, auto_increment, create_time, update_time, check_time, table_collation, checksum, create_options, table_comment) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? )").
		WithArgs("", "", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, time.Time{}, nil, nil, nil, nil, nil, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	store := NewTablesStore(ctx, db)
	err = store.Insert(&codegen.Tables{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTablesUpdate(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("UPDATE information_schema.TABLES SET table_catalog = ?,table_schema = ?,table_name = ?,table_type = ?,engine = ?,version = ?,row_format = ?,table_rows = ?,avg_row_length = ?,data_length = ?,max_data_length = ?,index_length = ?,data_free = ?,auto_increment = ?,create_time = ?,update_time = ?,check_time = ?,table_collation = ?,checksum = ?,create_options = ?,table_comment = ? WHERE ").
		WithArgs("", "", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, time.Time{}, nil, nil, nil, nil, nil, nil).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewTablesStore(ctx, db)
	aff, err := store.Update(&codegen.Tables{})
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

func TestTablesSelectWithoutJoin(t *testing.T) {
	ctx := &codegen.Context{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.table_catalog", "A.table_schema", "A.table_name", "A.table_type", "A.engine", "A.version", "A.row_format", "A.table_rows", "A.avg_row_length", "A.data_length", "A.max_data_length", "A.index_length", "A.data_free", "A.auto_increment", "A.create_time", "A.update_time", "A.check_time", "A.table_collation", "A.checksum", "A.create_options", "A.table_comment"}).
		AddRow("", "", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, time.Time{}, nil, nil, nil, nil, nil, nil).
		AddRow("", "", "", "", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, time.Time{}, nil, nil, nil, nil, nil, nil)

	mock.ExpectQuery("SELECT A.table_catalog, A.table_schema, A.table_name, A.table_type, A.engine, A.version, A.row_format, A.table_rows, A.avg_row_length, A.data_length, A.max_data_length, A.index_length, A.data_free, A.auto_increment, A.create_time, A.update_time, A.check_time, A.table_collation, A.checksum, A.create_options, A.table_comment FROM information_schema.TABLES A").
		WillReturnRows(rows)
	store := NewTablesStore(ctx, db).WithoutJoins()
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
