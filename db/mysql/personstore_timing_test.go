package mysql

import (
	"database/sql/driver"
	"testing"

	codegen "bitbucket.org/codegen"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

func BenchmarkPersonInsert(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(ctx, db)

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
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(ctx, db)

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
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	store := NewPersonStore(ctx, db)

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
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(personQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.Name})
	}
	store := NewPersonStore(ctx, db)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPersonSelectCols(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(personQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Person{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.Name})
	}
	store := NewPersonStore(ctx, db).Columns(codegen.Person_ID, codegen.Person_Name)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ^^ END OF GENERATED BY CODEGEN. ^^
