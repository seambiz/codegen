package codegen

import (
	"bitbucket.com/seambiz/sdb"
	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
	"github.com/valyala/bytebufferpool"
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
		schema.Tables = append(schema.Tables, t)
	}
	t.Name = tableName

	return t
}

func getIndex(table *Table, i int) *Index {
	var ind *Index

	if i > len(table.Indices)-1 {
		ind = &Index{}
		table.Indices = append(table.Indices, ind)
	} else {
		ind = table.Indices[i]
	}
	return ind
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
		sql.Append(" AND table_type = 'base table'")

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
				if len(table.Fields) < i+1 {
					table.Fields = append(table.Fields, &fields[i])
				} else {
					// reset bools
					table.Fields[i].IsAutoincrement = false
					table.Fields[i].IsNullable = false
					table.Fields[i].IsPrimaryKey = false
					mergo.MergeWithOverwrite(table.Fields[i], fields[i])
				}
			}

			if len(table.Fields) > len(fields) {
				table.Fields = table.Fields[:len(fields)]
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

			for i, indexName := range indices {
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

				tableIndex := getIndex(table, i)
				tableIndex.Name = indexName
				tableIndex.IsUnique = index[0].IsUnique
				tableIndex.Fields = make([]string, 0)
				for _, field := range index {
					tableIndex.Fields = append(tableIndex.Fields, field.ColName)
				}
			}
		}
	}

	bb := bytebufferpool.Get()
	defer bytebufferpool.Put(bb)
	enc := toml.NewEncoder(bb)
	err := enc.Encode(conf)
	if err != nil {
		return []byte(""), err
	}

	return bb.Bytes(), nil
}
