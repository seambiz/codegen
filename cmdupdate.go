package codegen

import (
	"encoding/json"
	"strings"

	"bitbucket.org/seambiz/seambiz/sdb"
	"github.com/imdario/mergo"
)

func getSchema(conf *Config, name string) *Schema {
	var s *Schema
	for i := range conf.Schemas {
		if name == conf.Schemas[i].Name {
			s = conf.Schemas[i]
			break
		}
	}
	if s == nil {
		s = &Schema{}
		conf.Schemas = append(conf.Schemas, s)
	}
	s.Name = name

	return s
}

func getTable(schema *Schema, tableName string) *Table {
	var t *Table
	for i := range schema.Tables {
		if schema.Tables[i].Name == tableName {
			t = schema.Tables[i]
			break
		}
	}
	if t == nil {
		t = &Table{}
		t.Generate = true
		schema.Tables = append(schema.Tables, t)
	}
	t.Name = tableName

	return t
}

func getField(table *Table, fieldName string) *Field {
	var f *Field

	for i := range table.Fields {
		if table.Fields[i].Name == fieldName {
			f = table.Fields[i]
			break
		}
	}
	if f == nil {
		f = &Field{}
		table.Fields = append(table.Fields, f)
		f.Name = fieldName
	}

	return f
}

func getIndex(table *Table, indexName string) *Index {
	var ind *Index

	indexName = strings.ToLower(indexName)

	for i := range table.Indices {
		if strings.ToLower(table.Indices[i].Name) == indexName {
			ind = table.Indices[i]
			break
		}
	}
	if ind == nil {
		ind = &Index{}
		ind.Generate = true
		table.Indices = append(table.Indices, ind)
		ind.Name = indexName
	}

	return ind
}

func getForeignKey(table *Table, fkName string) *ForeignKey {
	var fk *ForeignKey

	fkName = strings.ToLower(fkName)

	for i := range table.ForeignKeys {
		if strings.ToLower(table.ForeignKeys[i].Name) == fkName {
			fk = table.ForeignKeys[i]
			break
		}
	}
	if fk == nil {
		fk = &ForeignKey{}
		fk.IsUnique = true
		table.ForeignKeys = append(table.ForeignKeys, fk)
		fk.Name = fkName
	}

	return fk
}

// Update command
func Update(conf *Config) ([]byte, error) {
	db := sdb.OpenDatabaseDSN(conf.Database.DSN)

	for _, schemaName := range conf.Database.Schemas {
		schema := getSchema(conf, schemaName)

		sql := sdb.NewSQLStatement()
		sql.Append("SELECT table_name AS tables")
		sql.Append("FROM information_schema.tables")
		sql.Append("WHERE table_schema = ?")
		sql.Append(" AND table_type IN ('base table', 'system view')")

		var tables []string
		db.Select(&tables, sql.Query(), schema.Name)

		for _, tableName := range tables {
			table := getTable(schema, tableName)

			sql = sdb.NewSQLStatement()
			sql.Append("SELECT")
			sql.Append("column_name AS name,")
			sql.Append("IF(column_type = 'tinyint(1)',column_type, IF(INSTR(data_type, 'int'), IF(RIGHT(column_type, 8) = 'unsigned', CONCAT(data_type, ' unsigned'), data_type), data_type)) AS dbtype,")
			//sql.Append("COALESCE(column_default, '') AS `default`,")
			sql.Append("IF(is_nullable = 'YES', TRUE, FALSE) AS isnullable,")
			sql.Append("IF(INSTR(extra, 'auto_increment'), TRUE, FALSE) as isautoincrement,")
			sql.Append("IF(column_key = 'PRI', TRUE, FALSE) AS isprimarykey")
			sql.Append("FROM")
			sql.Append("information_schema.columns")
			sql.Append("WHERE")
			sql.Append("table_schema = ?")
			sql.Append("AND table_name = ?")
			sql.Append("ORDER BY ordinal_position")

			var fields []Field
			err := db.Select(&fields, sql.Query(), schema.Name, table.Name)
			if err != nil {
				panic(err)
			}

			for i := range fields {
				f := getField(table, fields[i].Name)
				f.IsAutoincrement = false
				f.IsNullable = false
				f.IsPrimaryKey = false
				mergo.MergeWithOverwrite(f, fields[i])
			}

			sql = sdb.NewSQLStatement()
			sql.Append("SELECT")
			sql.Append(" CONSTRAINT_NAME AS name, REFERENCED_TABLE_SCHEMA AS refschema, REFERENCED_TABLE_NAME AS reftable")
			sql.Append("FROM")
			sql.Append("INFORMATION_SCHEMA.KEY_COLUMN_USAGE")
			sql.Append("WHERE")
			sql.Append("TABLE_SCHEMA = ? and REFERENCED_TABLE_SCHEMA = ? AND TABLE_NAME = ?")
			sql.Append("ORDER BY CONSTRAINT_NAME, ORDINAL_POSITION")
			var fks []ForeignKey
			err = db.Select(&fks, sql.Query(), schema.Name, schema.Name, table.Name)
			if err != nil {
				panic(err)
			}
			for i := range fks {
				fk := getForeignKey(table, fks[i].Name)
				mergo.MergeWithOverwrite(fk, fks[i])
			}

			var indices []string
			sql = sdb.NewSQLStatement()
			sql.Append("SELECT")
			sql.Append("  DISTINCT(index_name) AS indexname")
			sql.Append("FROM information_schema.statistics")
			sql.Append("WHERE table_schema = ?")
			sql.Append("  AND table_name = ?")
			sql.Append("ORDER BY  index_name")
			err = db.Select(&indices, sql.Query(), schema.Name, table.Name)
			if err != nil {
				panic(err)
			}

			for _, indexName := range indices {
				sql = sdb.NewSQLStatement()
				sql.Append("SELECT")
				sql.Append("  index_name AS indexname,")
				sql.Append("  column_name AS colname,")
				sql.Append("  IF (non_unique = 0, TRUE, FALSE) as isunique")
				sql.Append("FROM information_schema.statistics")
				sql.Append("WHERE table_schema = ?")
				sql.Append("  AND table_name = ?")
				sql.Append("  AND index_name = ?")
				sql.Append("ORDER BY seq_in_index")

				var index []indexField
				err = db.Select(&index, sql.Query(), schema.Name, table.Name, indexName)
				if err != nil {
					panic(err)
				}

				tableIndex := getIndex(table, indexName)
				tableIndex.IsUnique = index[0].IsUnique
				tableIndex.Fields = make([]string, 0)
				for _, field := range index {
					tableIndex.Fields = append(tableIndex.Fields, field.ColName)
				}
			}
		}
	}

	jsonBytes, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return []byte(""), err
	}

	return jsonBytes, nil
}
