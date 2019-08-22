package codegen

import (
	"encoding/json"
	"fmt"
	"strings"

	"bitbucket.org/codegen/models"
	"bitbucket.org/seambiz/sdb"
	"github.com/imdario/mergo"
	"github.com/jmoiron/sqlx"
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
		fmt.Println("new table", tableName)
		t = &Table{}
		t.Generate = true
		schema.Tables = append(schema.Tables, t)
		t.Name = tableName
	}

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
	dbTemp := sdb.OpenDatabaseDSN(conf.Database.DSN)
	conn := sqlx.NewDb(dbTemp, "mysql")

	for _, schemaName := range conf.Database.Schemas {
		schema := getSchema(conf, schemaName)

		tables, err := models.NewTablesStore(conn.DB).Where("UPPER(table_schema) = UPPER(?) AND table_type IN (UPPER('base table'), UPPER('system view'))").Query(schema.Name)
		if err != nil {
			panic(err)
		}

		for _, table := range tables {
			table := getTable(schema, table.TableName)

			sql := sdb.NewSQLStatement()
			sql.Append("SELECT")
			sql.Append("column_name AS name,")
			sql.Append("IF(column_type = 'tinyint(1)',column_type, IF(INSTR(data_type, 'int'), IF(RIGHT(column_type, 8) = 'unsigned', CONCAT(data_type, ' unsigned'), data_type), data_type)) AS dbtype,")
			//sql.Append("COALESCE(column_default, '') AS `default`,")
			sql.Append("IF(is_nullable = 'YES', TRUE, FALSE) AS isnullable,")
			sql.Append("IF(INSTR(extra, 'auto_increment'), TRUE, FALSE) as isautoincrement,")
			sql.Append("IF(column_key = 'PRI', TRUE, FALSE) AS isprimarykey")
			sql.Append("FROM")
			sql.Append("information_schema.columns c")
			sql.Append("WHERE")
			sql.Append("UPPER(table_schema) = UPPER(?)")
			sql.Append("AND UPPER(table_name) = UPPER(?)")
			sql.Append("AND UPPER(c.extra) not like '%VIRTUAL%'")
			sql.Append("ORDER BY ordinal_position")

			var fields []Field
			err := conn.Select(&fields, sql.Query(), schema.Name, table.Name)
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

			foreignKeys, err := models.NewKeyColumnUsageStore(conn.DB).Where("UPPER(table_schema) = UPPER(?) AND UPPER(referenced_table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?)").
				OrderBy("constraint_name, ordinal_position").
				Query(schema.Name, schema.Name, table.Name)
			if err != nil {
				panic(err)
			}

			for i := range foreignKeys {
				fk := getForeignKey(table, foreignKeys[i].ConstraintName)
				fk.Fields = fk.Fields[:0]
				fk.Fields = append(fk.Fields, foreignKeys[i].ColumnName)
				fk.RefFields = fk.RefFields[:0]
				fk.RefFields = append(fk.RefFields, foreignKeys[i].ReferencedColumnName)
				fk.RefTable = foreignKeys[i].ReferencedTableName
				fk.RefSchema = foreignKeys[i].ReferencedTableSchema
				fk.IsUnique = true // TODO is always true?
				fk.Name = foreignKeys[i].ConstraintName
			}

			indices, err := models.NewStatisticsStore(conn.DB).
				Columns(models.StatisticsIndexName).
				Where("UPPER(table_schema) = UPPER(?) AND UPPER(table_name) = UPPER(?)").
				GroupBy("index_name").
				Query(schema.Name, table.Name)
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
				sql.Append("WHERE UPPER(table_schema) = UPPER(?)")
				sql.Append("  AND UPPER(table_name) = UPPER(?)")
				sql.Append("  AND UPPER(index_name) = UPPER(?)")
				sql.Append("ORDER BY seq_in_index")

				var index []indexField
				err = conn.Select(&index, sql.Query(), schema.Name, table.Name, indexName.IndexName)
				if err != nil {
					panic(err)
				}

				tableIndex := getIndex(table, indexName.IndexName)
				tableIndex.IsUnique = index[0].IsUnique
				tableIndex.Fields = make([]string, 0)
				for _, field := range index {
					tableIndex.Fields = append(tableIndex.Fields, field.ColName)
				}
			}
		}

		/*
			TODO: currently dont do the inverse stuff. it pollutes the code. actually needed foreign keys for eager fetching have to be added manually
				for _, t := range tables {
					table := getTable(schema, t.TableName)

					for i := range table.ForeignKeys {
						fk := getForeignKey(table, table.ForeignKeys[i].Name)
						refTable := getTable(getSchema(conf, fk.RefSchema), fk.RefTable)

						fkInverse := getForeignKey(refTable, table.ForeignKeys[i].Name)
						fkInverse.Fields = fk.RefFields
						fkInverse.RefFields = fk.Fields
						fkInverse.RefSchema = schema.Name
						fkInverse.RefTable = table.Name
						fkInverse.Name = fk.Name

						fkInverse.IsUnique = false
						for _, index := range table.Indices {
							if index.IsUnique && reflect.DeepEqual(index.Fields, fkInverse.RefFields) {
								fkInverse.IsUnique = true
								break
							}
						}

					}
				}
		*/
	}

	jsonBytes, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return []byte(""), err
	}

	return jsonBytes, nil
}
