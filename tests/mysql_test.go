package tests

import (
	"testing"

	"bitbucket.com/seambiz/logging"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/ziutek/mymysql/thrsafe" // or native
)

func TestValues(t *testing.T) {
	logging.Init(true, "")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// columns are prefixed with "o" since we used sqlstruct to generate them
	columns := []string{"id", "tinybool", "smallint", "mediumint", "int", "integer", "bigint", "utinyint", "usmallint", "umediumint", "uint",
		"uinteger", "ubigint", "float", "double", "decimal", "numeric", "bit", "year", "date", "time", "datetime", "timestamp", "char",
		"varchar", "tinytext", "text", "mediumtext", "longtext", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob"}

	// expect query to fetch order and user, match it with regexp
	mock.ExpectQuery("SELECT (.+) FROM codegen.tests A").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("10,1,2,3,4,5,6,7,8,9,10," +
			"11,12,1.0,2.0,3.0,4.0,1,2017,2017-01-03,16:01:13,2017-02-04 15:33:12,2017-02-04 15:33:44,a," +
			"b,c,d,e,f,g,h,i,j,k,l"))

	data, err := NewTestsStore(db).Query()
	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if len(data) != 1 {
		t.Errorf("Expected 1 row, but got %d instead", len(data))
	}
	d := data[0]
	if d.ID != 10 {
		t.Errorf("Expected d.ID to be 10")
	}
	if !d.Tinybool {
		t.Errorf("Expected true")
	}
	if d.Smallint != 2 {
		t.Errorf("Expected d.Smallint to be 2")
	}
}

func TestNull(t *testing.T) {
	logging.Init(true, "")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// columns are prefixed with "o" since we used sqlstruct to generate them
	columns := []string{"id", "tinybool", "smallint", "mediumint", "int", "integer", "bigint", "utinyint", "usmallint", "umediumint", "uint",
		"uinteger", "ubigint", "float", "double", "decimal", "numeric", "bit", "year", "date", "time", "datetime", "timestamp", "char",
		"varchar", "tinytext", "text", "mediumtext", "longtext", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob"}

	// expect query to fetch order and user, match it with regexp
	mock.ExpectQuery("SELECT (.+) FROM codegen.tests A").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil))

	data, err := NewTestsStore(db).Query()
	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if len(data) != 1 {
		t.Errorf("Expected 1 row, but got %d instead", len(data))
	}
	d := data[0]
	if d.ID != 0 {
		t.Errorf("Expected d.ID to be 0")
	}
	if d.Tinybool {
		t.Errorf("Expected false")
	}
	if d.Smallint != 0 {
		t.Errorf("Expected d.Smallint to be 0")
	}
}
