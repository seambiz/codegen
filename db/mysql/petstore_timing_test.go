package mysql

import (
	"database/sql/driver"
	"testing"

	codegen "github.com/seambiz/codegen"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/rs/zerolog"
)

// GENERATED BY CODEGEN.

func BenchmarkPetInsert(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Pet{}
	store := NewPetStore(ctx, db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := store.Insert(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPetUpdate(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Pet{}
	store := NewPetStore(ctx, db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Update(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPetDelete(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Pet{}
	store := NewPetStore(ctx, db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Delete(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPetSelectAll(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(petQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Pet{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.PersonID, data.TagID, data.Species})
	}
	store := NewPetStore(ctx, db)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPetSelectCols(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(petQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Pet{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.ID, data.Species})
	}
	store := NewPetStore(ctx, db).Columns(codegen.Pet_ID, codegen.Pet_Species)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ^^ END OF GENERATED BY CODEGEN. ^^
