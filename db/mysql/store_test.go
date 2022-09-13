package mysql

import (
	"database/sql"
	"testing"

	codegen "bitbucket.org/codegen"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
	"github.com/seambiz/seambiz/sdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Insert inserts the Tables to the database.
func TestStoreMapScan(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	stmt := sdb.NewSQLStatement()
	stmt.Append("SELECT * FROM dummy.table")

	rows := sqlmock.NewRows([]string{"id", "col", "null"}).
		AddRow(1, "test", nil)
	mock.
		ExpectQuery("SELECT * FROM dummy.table").
		WillReturnRows(rows)

	res := map[string]sql.RawBytes{}
	err = NewStore(db, ctx.Log).SQL(stmt.Query()).MapScan(res)
	if err != nil {
		t.Fatalf("SQL error '%s'", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Equal(t, 1, sdb.ToInt(res["id"]), "int")
	assert.Equal(t, "test", sdb.ToString(res["col"]), "string")
	assert.Nil(t, res["null"], "nil value")
}

// TestStoreSelfJoin just shows the specific use-case some code generators have problem with.
func TestStoreSelfJoin(t *testing.T) {
	ctx := &codegen.BaseContext{Log: &zerolog.Logger{}}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "id", "name"}).
		AddRow(1, "person", 2, "another person")
	mock.ExpectQuery("SELECT A.id, A.name , B.id, B.name FROM fake_benchmark.Person A INNER JOIN fake_benchmark B ON (A.id = B.id+1)").
		WillReturnRows(rows)

	store := NewStore(db, ctx.Log).
		SelectFields("A", PersonQueryFields).
		SelectFields("B", PersonQueryFields).
		SQL("FROM fake_benchmark.Person A ").
		SQL("INNER JOIN fake_benchmark B ON (A.id = B.id+1)")
	var dest struct {
		codegen.Person
		Person2 struct {
			codegen.Person
		}
	}
	err = store.OneBind(&dest)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

	assert.Equal(t, 1, dest.ID)
	assert.Equal(t, "person", dest.Name)
	assert.Equal(t, 2, dest.Person2.ID)
	assert.Equal(t, "another person", dest.Person2.Name)
}

// TestStoreOne shows the different functions to query for a single column.
func TestStoreOne(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sql := sdb.NewSQLStatement()
	sql.Append("SELECT 1 FROM dummy.table")

	t.Run("OneInt", func(t *testing.T) {
		store := NewStore(db).SQL(sql.String())
		rows := sqlmock.NewRows([]string{"count"}).
			AddRow(1)

		mock.ExpectQuery("SELECT 1 FROM dummy.table").
			WillReturnRows(rows)
		res, err := store.OneInt()
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Equal(t, 1, res, "single int")
	})

	t.Run("OneString", func(t *testing.T) {
		store := NewStore(db).SQL(sql.String())
		rows := sqlmock.NewRows([]string{"count"}).
			AddRow(1)

		mock.ExpectQuery("SELECT 1 FROM dummy.table").
			WillReturnRows(rows)
		res, err := store.OneString()
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Equal(t, "1", res, "single string")
	})

	t.Run("OneBool", func(t *testing.T) {
		store := NewStore(db).SQL(sql.String())
		rows := sqlmock.NewRows([]string{"count"}).
			AddRow(1)

		mock.ExpectQuery("SELECT 1 FROM dummy.table").
			WillReturnRows(rows)
		res, err := store.OneBool()
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Equal(t, true, res, "single bool")
	})

	t.Run("OneValue", func(t *testing.T) {
		store := NewStore(db).SQL(sql.String())
		rows := sqlmock.NewRows([]string{"count"}).
			AddRow(1)

		mock.ExpectQuery("SELECT 1 FROM dummy.table").
			WillReturnRows(rows)
		res, err := store.OneValue(func(b []byte) interface{} {
			return sdb.ToInt64(b)
		})
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Equal(t, int64(1), res, "single value")
	})
}

func prepareMockColumns(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("SELECT A.id, A.species FROM fake_benchmark.pet A WHERE A.id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species"}).
			AddRow(1, "cat"))
}

// TestStoreRawSQL show different methods to query for specific columns.
func TestStoreRawSQL(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	prepareMockColumns(mock)
	prepareMockColumns(mock)
	prepareMockColumns(mock)

	/**********************************************************************
	 * PetStore binding
	**********************************************************************/
	pet, err := NewPetStore(db).
		Columns(codegen.Pet_ID, codegen.Pet_Species).
		Where("A.id = ?").
		WithoutJoins().
		One(1)
	assert.Nil(t, err)
	assert.NotNil(t, pet)

	/**********************************************************************
	 * OnexInto (sqlx) binding
	**********************************************************************/
	petx := &codegen.Pet{}
	err = NewStore(db).
		SQL("SELECT A.id, A.species ").
		SQL("FROM fake_benchmark.pet A ").
		SQL("WHERE A.id = ?").
		OnexInto(petx, 1)
	assert.Nil(t, err)
	assert.NotNil(t, petx)

	/**********************************************************************
	 * Store general binding
	**********************************************************************/
	var petbase struct {
		codegen.Pet
	}
	err = NewStore(db).
		SelectFields("A", PetQueryFields, codegen.Pet_ID, codegen.Pet_Species).
		SQL("FROM fake_benchmark.pet A ").
		SQL("WHERE A.id = ?").
		OneBind(&petbase, 1)
	assert.Nil(t, err)
	assert.NotNil(t, petbase)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	masterPet := &codegen.Pet{}
	masterPet.ID = 1
	masterPet.Species = "cat"

	assert.Equal(t, masterPet, pet, "petstore")
	assert.Equal(t, masterPet, petx, "store sqlx")
	assert.Equal(t, masterPet, &petbase.Pet, "store qrm")
}

// test helper
func testJoins(db *sql.DB, mock sqlmock.Sqlmock) *Store {
	mock.
		ExpectQuery("SELECT A.id, A.name , B.id, B.person_id, B.tag_id, B.species , C.id, C.name FROM fake_benchmark.person A INNER JOIN fake_benchmark.pet B ON (A.id = B.person_id) INNER JOIN fake_benchmark.tag C ON (B.tag_id = C.id) WHERE person.id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "id", "person_id", "tag_id", "species", "id", "name"}).
			AddRow(1, "unknown", 1, 1, 1, "cat", 1, "one").
			AddRow(2, "unknown", 2, 2, 2, "dog", 2, "two").
			AddRow(3, "unknown", 3, 3, 2, "weasel", 2, "two"))

	return NewStore(db).
		SelectFields("A", PersonQueryFields).
		SelectFields("B", PetQueryFields).
		SelectFields("C", TagQueryFields).
		SQL("FROM fake_benchmark.person A ").
		SQL("INNER JOIN fake_benchmark.pet B ON (A.id = B.person_id) ").
		SQL("INNER JOIN fake_benchmark.tag C ON (B.tag_id = C.id) ").
		SQL("WHERE person.id = ?")
}

// test helper
func execScanToStruct(dest interface{}, t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.
		ExpectQuery("SELECT A.id, A.name , B.id, B.person_id, B.tag_id, B.species , C.id, C.name FROM fake_benchmark.person A INNER JOIN fake_benchmark.pet B ON (A.id = B.person_id) INNER JOIN fake_benchmark.tag C ON (B.tag_id = C.id) WHERE person.id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "id", "person_id", "tag_id", "species", "id", "name"}).
			AddRow(1, "unknown", 1, 1, 1, "cat", 1, "one"))

	err = NewStore(db).
		SelectFields("A", PersonQueryFields).
		SelectFields("B", PetQueryFields).
		SelectFields("C", TagQueryFields).
		SQL("FROM fake_benchmark.person A ").
		SQL("INNER JOIN fake_benchmark.pet B ON (A.id = B.person_id) ").
		SQL("INNER JOIN fake_benchmark.tag C ON (B.tag_id = C.id) ").
		SQL("WHERE person.id = ?").
		OneBind(dest, 1)

	require.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestScanToStruct(t *testing.T) {
	tag1 := codegen.Tag{ID: 1, Name: "one"}
	person1 := codegen.Person{ID: 1, Name: "unknown"}
	pet1 := codegen.Pet{ID: 1, PersonID: 1, TagID: 1, Species: "cat"}

	t.Run("empy struct", func(t *testing.T) {
		var dest struct{}
		execScanToStruct(&dest, t)
	})
	t.Run("one struct", func(t *testing.T) {
		var dest struct {
			codegen.Person
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
	})
	t.Run("all embedded", func(t *testing.T) {
		var dest struct {
			codegen.Person
			codegen.Pet
			codegen.Tag
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
		assert.Equal(t, pet1, dest.Pet)
		assert.Equal(t, tag1, dest.Tag)
	})
	t.Run("ignore named structs", func(t *testing.T) {
		var dest struct {
			codegen.Person
			P2 codegen.Person
			codegen.Pet
			codegen.Tag
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
		assert.Equal(t, pet1, dest.Pet)
		assert.Equal(t, tag1, dest.Tag)
	})
	t.Run("nested structs 1", func(t *testing.T) {
		var dest struct {
			codegen.Person
			Pet struct {
				codegen.Pet
			}
			codegen.Tag
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
		assert.Equal(t, pet1, dest.Pet.Pet)
		assert.Equal(t, tag1, dest.Tag)
	})
	t.Run("nested structs 2", func(t *testing.T) {
		var dest struct {
			codegen.Person
			Pet struct {
				codegen.Pet
				codegen.Tag
			}
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
		assert.Equal(t, pet1, dest.Pet.Pet)
		assert.Equal(t, tag1, dest.Pet.Tag)
	})
	t.Run("nested structs 3", func(t *testing.T) {
		var dest struct {
			codegen.Person
			Pet struct {
				codegen.Pet
				Tag struct {
					codegen.Tag
				}
			}
		}
		execScanToStruct(&dest, t)

		assert.Equal(t, person1, dest.Person)
		assert.Equal(t, pet1, dest.Pet.Pet)
		assert.Equal(t, tag1, dest.Pet.Tag.Tag)
	})
}

func TestScanToSlice(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tag1 := codegen.Tag{ID: 1, Name: "one"}
	tag2 := codegen.Tag{ID: 2, Name: "two"}

	person1 := codegen.Person{ID: 1, Name: "unknown"}
	person2 := codegen.Person{ID: 2, Name: "unknown"}
	person3 := codegen.Person{ID: 3, Name: "unknown"}

	pet1 := codegen.Pet{ID: 1, PersonID: 1, TagID: 1, Species: "cat"}
	pet2 := codegen.Pet{ID: 2, PersonID: 2, TagID: 2, Species: "dog"}
	pet3 := codegen.Pet{ID: 3, PersonID: 3, TagID: 2, Species: "weasel"}

	t.Run("all embedded", func(t *testing.T) {
		var dest []struct {
			codegen.Person
			codegen.Pet
			codegen.Tag
		}

		store := testJoins(db, mock)
		err = store.QueryBind(&dest, 1)
		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())

		assert.Equal(t, 3, len(dest))

		assert.Equal(t, person1, dest[0].Person)
		assert.Equal(t, pet1, dest[0].Pet)
		assert.Equal(t, tag1, dest[0].Tag)

		assert.Equal(t, person2, dest[1].Person)
		assert.Equal(t, pet2, dest[1].Pet)
		assert.Equal(t, tag2, dest[1].Tag)

		assert.Equal(t, person3, dest[2].Person)
		assert.Equal(t, pet3, dest[2].Pet)
		assert.Equal(t, tag2, dest[2].Tag)
	})

	t.Run("nested structs", func(t *testing.T) {
		var dest []struct {
			codegen.Person

			Pet codegen.Pet
			Tag struct {
				codegen.Tag
			}
		}

		store := testJoins(db, mock)
		err = store.QueryBind(&dest, 1)
		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())

		assert.Equal(t, 3, len(dest))

		assert.Equal(t, person1, dest[0].Person)
		assert.Equal(t, pet1, dest[0].Pet)
		assert.Equal(t, tag1, dest[0].Tag.Tag)

		assert.Equal(t, person2, dest[1].Person)
		assert.Equal(t, pet2, dest[1].Pet)
		assert.Equal(t, tag2, dest[1].Tag.Tag)

		assert.Equal(t, person3, dest[2].Person)
		assert.Equal(t, pet3, dest[2].Pet)
		assert.Equal(t, tag2, dest[2].Tag.Tag)
	})

	t.Run("nested structs 2", func(t *testing.T) {
		var dest []struct {
			codegen.Person

			Pet struct {
				codegen.Pet
				codegen.Tag
			}
		}

		store := testJoins(db, mock)
		err = store.QueryBind(&dest, 1)
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, 3, len(dest))

		assert.Equal(t, person1, dest[0].Person)
		assert.Equal(t, pet1, dest[0].Pet.Pet)
		assert.Equal(t, tag1, dest[0].Pet.Tag)

		assert.Equal(t, person2, dest[1].Person)
		assert.Equal(t, pet2, dest[1].Pet.Pet)
		assert.Equal(t, tag2, dest[1].Pet.Tag)

		assert.Equal(t, person3, dest[2].Person)
		assert.Equal(t, pet3, dest[2].Pet.Pet)
		assert.Equal(t, tag2, dest[2].Pet.Tag)
	})

	t.Run("nested structs 3", func(t *testing.T) {
		var dest []struct {
			codegen.Person

			Pet struct {
				codegen.Pet

				Tag struct {
					codegen.Tag
				}
			}
		}

		store := testJoins(db, mock)
		err = store.QueryBind(&dest, 1)
		if err != nil {
			t.Fatalf("SQL error '%s'", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, 3, len(dest))

		assert.Equal(t, person1, dest[0].Person)
		assert.Equal(t, pet1, dest[0].Pet.Pet)
		assert.Equal(t, tag1, dest[0].Pet.Tag.Tag)

		assert.Equal(t, person2, dest[1].Person)
		assert.Equal(t, pet2, dest[1].Pet.Pet)
		assert.Equal(t, tag2, dest[1].Pet.Tag.Tag)

		assert.Equal(t, person3, dest[2].Person)
		assert.Equal(t, pet3, dest[2].Pet.Pet)
		assert.Equal(t, tag2, dest[2].Pet.Tag.Tag)
	})
}
