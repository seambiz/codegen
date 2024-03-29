package updater

import (
	"sort"
	"strings"

	"github.com/seambiz/codegen"
	"github.com/seambiz/codegen/config"
	"golang.org/x/exp/slices"

	"github.com/imdario/mergo"
)

type MysqlUpdate struct {
	repoTable   codegen.TablesRepository
	repoColumns codegen.ColumnsRepository
	repoKeyCol  codegen.KeyColumnUsageRepository
	repoStats   codegen.StatisticsRepository
	ctx         *codegen.Context
}

func NewMysqlUpdate(ctx *codegen.Context, table codegen.TablesRepository, columns codegen.ColumnsRepository, keycol codegen.KeyColumnUsageRepository, stats codegen.StatisticsRepository) *MysqlUpdate {
	return &MysqlUpdate{
		repoTable:   table,
		repoColumns: columns,
		repoKeyCol:  keycol,
		repoStats:   stats,
		ctx:         ctx,
	}
}

func (MysqlUpdate) getSchema(conf *config.Config, name string) *config.Schema {
	var s *config.Schema
	for i := range conf.Schemas {
		if name == conf.Schemas[i].Name {
			s = conf.Schemas[i]
			break
		}
	}
	if s == nil {
		s = &config.Schema{}
		conf.Schemas = append(conf.Schemas, s)
	}
	s.Name = name

	return s
}

func (MysqlUpdate) getTable(schema *config.Schema, tableName string) *config.Table {
	var t *config.Table
	for i := range schema.Tables {
		if schema.Tables[i].Name == tableName {
			t = schema.Tables[i]
			break
		}
	}
	if t == nil {
		t = &config.Table{}
		t.Generate = true
		schema.Tables = append(schema.Tables, t)
		t.Name = tableName
	}

	return t
}

func (MysqlUpdate) getField(table *config.Table, fieldName string) *config.Field {
	var f *config.Field

	for i := range table.Fields {
		if table.Fields[i].Name == fieldName {
			f = table.Fields[i]
			break
		}
	}
	if f == nil {
		f = &config.Field{}
		table.Fields = append(table.Fields, f)
		f.Name = fieldName
	}

	return f
}

func (MysqlUpdate) getIndex(table *config.Table, indexName string) *config.Index {
	var ind *config.Index

	indexName = strings.ToLower(indexName)

	for i := range table.Indices {
		if strings.ToLower(table.Indices[i].Name) == indexName {
			ind = table.Indices[i]
			break
		}
	}
	if ind == nil {
		ind = &config.Index{}
		ind.Generate = true
		table.Indices = append(table.Indices, ind)
		ind.Name = indexName
	}

	return ind
}

func (MysqlUpdate) getForeignKey(table *config.Table, fkName string) *config.ForeignKey {
	var fk *config.ForeignKey

	fkName = strings.ToLower(fkName)

	for i := range table.ForeignKeys {
		if strings.ToLower(table.ForeignKeys[i].Name) == fkName {
			fk = table.ForeignKeys[i]
			break
		}
	}
	if fk == nil {
		fk = &config.ForeignKey{}
		fk.IsUnique = true
		table.ForeignKeys = append(table.ForeignKeys, fk)
		fk.Name = fkName
	}

	return fk
}

func getEnumValues(s string) []string {
	s = strings.Replace(s, "enum(", "", 1)
	s = s[1 : len(s)-2]

	return strings.Split(s, "','")
}

// Update command
func (u MysqlUpdate) Update(conf *config.Config) (config.Config, error) {
	for _, schemaName := range conf.Database.Schemas {
		schema := u.getSchema(conf, schemaName)

		tables, err := u.repoTable.QueryBySchema(u.ctx, schema.Name)
		if err != nil {
			panic(err)
		}

		for _, table := range tables {
			if len(schema.TableNames) > 0 {
				if !slices.Contains(schema.TableNames, table.TableName) {
					continue
				}
			}

			genTable := u.getTable(schema, table.TableName)

			if len(schema.IgnoreTableNames) > 0 {
				if slices.Contains(schema.IgnoreTableNames, table.TableName) {
					// do not skip. add a no generate
					genTable.Generate = false
				}
			}

			cols, err := u.repoColumns.QueryBySchemaAndTable(u.ctx, schema.Name, genTable.Name)
			if err != nil {
				panic(err)
			}

			for i := range cols {
				fRef := u.getField(genTable, *cols[i].ColumnName)
				fRef.IsAutoincrement = false
				fRef.IsNullable = false
				fRef.IsPrimaryKey = false

				c := cols[i]
				fNew := config.Field{}
				fNew.Name = *c.ColumnName
				if c.ColumnType == "tinyint(1)" {
					fNew.DBType = c.ColumnType
				} else {
					if strings.Contains(*c.DataType, "int") && strings.Contains(c.ColumnType, " unsigned") {
						fNew.DBType = *c.DataType + " unsigned"
					} else {
						fNew.DBType = *c.DataType
					}
				}
				fNew.IsNullable = c.IsNullable == "YES"
				fNew.IsAutoincrement = strings.Contains(*c.Extra, "auto_increment")
				fNew.IsPrimaryKey = c.ColumnKey == "PRI"

				if fNew.DBType == "enum" {
					fNew.EnumValues = getEnumValues(c.ColumnType)
				}

				mergo.MergeWithOverwrite(fRef, fNew)
			}

			foreignKeys, err := u.repoKeyCol.QueryBySchemaAndRefSchemaAndTable(u.ctx, schema.Name, schema.Name, genTable.Name)
			if err != nil {
				panic(err)
			}

			for i := range foreignKeys {
				fk := u.getForeignKey(genTable, *foreignKeys[i].ConstraintName)
				// fk.Fields = fk.Fields[:0]
				fk.Fields = append(fk.Fields, *foreignKeys[i].ColumnName)
				// fk.RefFields = fk.RefFields[:0]
				fk.RefFields = append(fk.RefFields, *foreignKeys[i].ReferencedColumnName)
				fk.RefTable = *foreignKeys[i].ReferencedTableName
				fk.RefSchema = *foreignKeys[i].ReferencedTableSchema
				fk.IsUnique = true
				fk.Name = *foreignKeys[i].ConstraintName
			}

			indices, err := u.repoStats.IndexNameBySchemaAndTable(u.ctx, schema.Name, genTable.Name)
			if err != nil {
				panic(err)
			}

			for _, indexName := range indices {
				indexFields, err := u.repoStats.QueryBySchemaAndTableAndIndex(u.ctx, schema.Name, genTable.Name, *indexName.IndexName)
				if err != nil {
					panic(err)
				}

				tableIndex := u.getIndex(genTable, *indexName.IndexName)
				tableIndex.IsUnique = indexFields[0].NonUnique == 0
				tableIndex.Fields = make([]string, 0)
				for _, field := range indexFields {
					tableIndex.Fields = append(tableIndex.Fields, *field.ColumnName)
				}
			}
			sort.Slice(genTable.Indices, func(i, j int) bool { return genTable.Indices[i].Name < genTable.Indices[j].Name })
		}
		sort.Slice(schema.Tables, func(i, j int) bool { return schema.Tables[i].Name < schema.Tables[j].Name })

		/* Attention: currently dont do the inverse stuff. it pollutes the code. actually needed foreign keys for eager fetching have to be added manually
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
	return *conf, nil
}
