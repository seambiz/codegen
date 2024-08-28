package mysql

import (
	"database/sql/driver"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/rs/zerolog"
	codegen "github.com/seambiz/codegen"
)

// GENERATED BY CODEGEN.

func BenchmarkTablesInsert(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Tables{}
	store := NewTablesStore(ctx, db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := store.Insert(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkTablesUpdate(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Tables{}
	store := NewTablesStore(ctx, db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Update(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkTablesSelectAll(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(tablesQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Tables{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.TableCatalog, data.TableSchema, data.TableName, data.TableType, *data.Engine, *data.Version, *data.RowFormat, *data.TableRows, *data.AvgRowLength, *data.DataLength, *data.MaxDataLength, *data.IndexLength, *data.DataFree, *data.AutoIncrement, data.CreateTime, *data.UpdateTime, *data.CheckTime, *data.TableCollation, *data.Checksum, *data.CreateOptions, *data.TableComment})
	}
	store := NewTablesStore(ctx, db)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTablesSelectCols(b *testing.B) {
	b.ReportAllocs()
	ctx := &codegen.Context{Log: &zerolog.Logger{}}

	db, err := selectQuery(tablesQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Tables{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.TableCatalog, *data.TableComment})
	}
	store := NewTablesStore(ctx, db).Columns(codegen.Tables_TableCatalog, codegen.Tables_TableComment)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ^^ END OF GENERATED BY CODEGEN. ^^
