package mysql

import (
	"database/sql/driver"
	"testing"

	codegen "bitbucket.org/codegen"

	"github.com/brianvoe/gofakeit/v5"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

func BenchmarkPersonInsert(b *testing.B) {
	b.ReportAllocs()

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := store.Insert(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPersonUpdate(b *testing.B) {
	b.ReportAllocs()

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Update(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
func BenchmarkPersonDelete(b *testing.B) {
	b.ReportAllocs()

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Delete(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPersonSelectAll(b *testing.B) {
	b.ReportAllocs()

	db, err := selectQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.Name})
	}

	store := NewPersonStore(db)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPersonSelectCols(b *testing.B) {
	b.ReportAllocs()

	db, err := selectQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.Name})
	}

	store := NewPersonStore(db).Columns(Person_ID, Person_Name)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^