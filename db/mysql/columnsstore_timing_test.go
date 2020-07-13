package mysql

import (
	"database/sql/driver"
	"testing"

	"bitbucket.org/codegen"
	"github.com/brianvoe/gofakeit/v5"
)

// GENERATED BY CODEGEN. DO NOT EDIT.

func BenchmarkColumnsInsert(b *testing.B) {
	b.ReportAllocs()

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Columns{}
	store := NewColumnsStore(db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := store.Insert(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkColumnsUpdate(b *testing.B) {
	b.ReportAllocs()

	db, err := insertQuery()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Columns{}
	store := NewColumnsStore(db)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Update(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkColumnsSelectAll(b *testing.B) {
	b.ReportAllocs()

	db, err := selectQuery(columnsQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Columns{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.TableCatalog, data.TableSchema, data.TableName, data.ColumnName, data.OrdinalPosition, *data.ColumnDefault, data.IsNullable, data.DataType, *data.CharacterMaximumLength, *data.CharacterOctetLength, *data.NumericPrecision, *data.NumericScale, *data.DatetimePrecision, *data.CharacterSetName, *data.CollationName, data.ColumnType, data.ColumnKey, data.Extra, data.Privileges, data.ColumnComment, data.GenerationExpression})
	}
	store := NewColumnsStore(db)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColumnsSelectCols(b *testing.B) {
	b.ReportAllocs()

	db, err := selectQuery(columnsQueryFieldsAll)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	data := &codegen.Columns{}
	gofakeit.Struct(data)
	for i := 0; i < 100; i++ {
		addResultRowDSN("bench", []driver.Value{data.TableCatalog, data.GenerationExpression})
	}
	store := NewColumnsStore(db).Columns(codegen.Columns_TableCatalog, codegen.Columns_GenerationExpression)

	for i := 0; i < b.N; i++ {
		_, err = store.Query()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
