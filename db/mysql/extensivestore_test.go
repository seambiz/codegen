package mysql

import (
	"testing"
	"time"

	codegen "bitbucket.org/codegen"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

// Insert inserts the Extensive to the database.
func TestExtensiveInsert(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("INSERT INTO fake_benchmark.extensive (id, tinyint, tinyint_unsigned, tinyint_null, smallint, smallint_unsigned, smallint_null, int, int_null, int_unsigned, bigint, bigint_null, bigint_unsigned, varchar, varchar_null, float, float_null, double, double_null, decimal, decimal_null, numeric, numeric_null, created_at, updated_at, tinyint1, tinyint1_null, year, year_null, date, date_null, time, time_null, datetime, datetime_null, timestamp, timestamp_null, char, char_null, tinytext, tinytext_null, text, text_null, mediumtext, mediumtext_null, longtext, longtext_null, binary, binary_null, varbinary, varbinary_null, tinyblob, tinyblob_null, blob, blob_null, mediumblob, mediumblob_null, longblob, longblob_null) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? )").
		WithArgs(0, 0, 0, nil, 0, 0, nil, 0, nil, 0, 0, nil, 0, "", nil, 0.0, nil, 0.0, nil, 0.0, nil, 0.0, nil, TimeSec5{}, TimeSec5{}, false, nil, 0, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, "", nil, "", nil, "", nil, "", nil, "", nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	store := NewExtensiveStore(ctx, db)
	err = store.Insert(&codegen.Extensive{})
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExtensiveUpdate(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("UPDATE fake_benchmark.extensive SET tinyint = ?,tinyint_unsigned = ?,tinyint_null = ?,smallint = ?,smallint_unsigned = ?,smallint_null = ?,int = ?,int_null = ?,int_unsigned = ?,bigint = ?,bigint_null = ?,bigint_unsigned = ?,varchar = ?,varchar_null = ?,float = ?,float_null = ?,double = ?,double_null = ?,decimal = ?,decimal_null = ?,numeric = ?,numeric_null = ?,created_at = ?,updated_at = ?,tinyint1 = ?,tinyint1_null = ?,year = ?,year_null = ?,date = ?,date_null = ?,time = ?,time_null = ?,datetime = ?,datetime_null = ?,timestamp = ?,timestamp_null = ?,char = ?,char_null = ?,tinytext = ?,tinytext_null = ?,text = ?,text_null = ?,mediumtext = ?,mediumtext_null = ?,longtext = ?,longtext_null = ?,binary = ?,binary_null = ?,varbinary = ?,varbinary_null = ?,tinyblob = ?,tinyblob_null = ?,blob = ?,blob_null = ?,mediumblob = ?,mediumblob_null = ?,longblob = ?,longblob_null = ? WHERE id = ?").
		WithArgs(0, 0, nil, 0, 0, nil, 0, nil, 0, 0, nil, 0, "", nil, 0.0, nil, 0.0, nil, 0.0, nil, 0.0, nil, 0, TimeSec5{}, false, nil, 0, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, "", nil, "", nil, "", nil, "", nil, "", nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, 0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewExtensiveStore(ctx, db)
	aff, err := store.Update(&codegen.Extensive{})
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

func TestExtensiveSelectWithoutJoin(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"A.id", "A.tinyint", "A.tinyint_unsigned", "A.tinyint_null", "A.smallint", "A.smallint_unsigned", "A.smallint_null", "A.int", "A.int_null", "A.int_unsigned", "A.bigint", "A.bigint_null", "A.bigint_unsigned", "A.varchar", "A.varchar_null", "A.float", "A.float_null", "A.double", "A.double_null", "A.decimal", "A.decimal_null", "A.numeric", "A.numeric_null", "A.created_at", "A.updated_at", "A.tinyint1", "A.tinyint1_null", "A.year", "A.year_null", "A.date", "A.date_null", "A.time", "A.time_null", "A.datetime", "A.datetime_null", "A.timestamp", "A.timestamp_null", "A.char", "A.char_null", "A.tinytext", "A.tinytext_null", "A.text", "A.text_null", "A.mediumtext", "A.mediumtext_null", "A.longtext", "A.longtext_null", "A.binary", "A.binary_null", "A.varbinary", "A.varbinary_null", "A.tinyblob", "A.tinyblob_null", "A.blob", "A.blob_null", "A.mediumblob", "A.mediumblob_null", "A.longblob", "A.longblob_null"}).
		AddRow(0, 0, 0, nil, 0, 0, nil, 0, nil, 0, 0, nil, 0, "", nil, 0.0, nil, 0.0, nil, 0.0, nil, 0.0, nil, 0, 0, false, nil, 0, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, "", nil, "", nil, "", nil, "", nil, "", nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil).
		AddRow(0, 0, 0, nil, 0, 0, nil, 0, nil, 0, 0, nil, 0, "", nil, 0.0, nil, 0.0, nil, 0.0, nil, 0.0, nil, 0, 0, false, nil, 0, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, time.Time{}, nil, "", nil, "", nil, "", nil, "", nil, "", nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil, []byte(nil), nil)

	mock.ExpectQuery("SELECT A.id, A.tinyint, A.tinyint_unsigned, A.tinyint_null, A.smallint, A.smallint_unsigned, A.smallint_null, A.int, A.int_null, A.int_unsigned, A.bigint, A.bigint_null, A.bigint_unsigned, A.varchar, A.varchar_null, A.float, A.float_null, A.double, A.double_null, A.decimal, A.decimal_null, A.numeric, A.numeric_null, A.created_at, A.updated_at, A.tinyint1, A.tinyint1_null, A.year, A.year_null, A.date, A.date_null, A.time, A.time_null, A.datetime, A.datetime_null, A.timestamp, A.timestamp_null, A.char, A.char_null, A.tinytext, A.tinytext_null, A.text, A.text_null, A.mediumtext, A.mediumtext_null, A.longtext, A.longtext_null, A.binary, A.binary_null, A.varbinary, A.varbinary_null, A.tinyblob, A.tinyblob_null, A.blob, A.blob_null, A.mediumblob, A.mediumblob_null, A.longblob, A.longblob_null FROM fake_benchmark.extensive A").
		WillReturnRows(rows)
	store := NewExtensiveStore(ctx, db).WithoutJoins()
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

func TestExtensiveDelete(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.extensive WHERE id = ?").
		WithArgs(0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	aff, err := NewExtensiveStore(ctx, db).Delete(&codegen.Extensive{})
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

func TestExtensiveDeleteSlice(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.
		ExpectExec("DELETE FROM fake_benchmark.extensive WHERE id IN (0,0)").
		WillReturnResult(sqlmock.NewResult(0, 2))
	aff, err := NewExtensiveStore(ctx, db).DeleteSlice([]*codegen.Extensive{{}, {}})
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
