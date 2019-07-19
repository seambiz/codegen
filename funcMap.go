package codegen

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/valyala/bytebufferpool"
)

func getFuncMap() template.FuncMap {
	fmap := sprig.TxtFuncMap()
	fmap["genSelectSQL"] = genSelectSQL
	fmap["genSelectSQLMultiTenant"] = genSelectSQLMultiTenant
	fmap["genOneSelect"] = genOneSelect
	fmap["genQuerySelect"] = genQuerySelect
	fmap["genQueryCustom"] = genQueryCustom
	fmap["genUpsert"] = genUpsert
	fmap["genUpdate"] = genUpdate
	fmap["genUpdateMultiTenant"] = genUpdateMultiTenant
	fmap["genInsert"] = genInsert
	fmap["genInsertMultiTenant"] = genInsertMultiTenant
	fmap["genDelete"] = genDelete
	fmap["genForeign"] = genForeign
	fmap["genIndex"] = genIndex
	fmap["genTruncate"] = genTruncate
	fmap["genJSON"] = genJSON

	return fmap
}

// selectSQL generates general SELECT Statement with optional JOINs based on foreign key definitions
func genSelectSQL(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Func(table.StoreReceiver, "selectStatement")
	bb.FuncParams()
	bb.FuncReturn("*SQLStatement")

	tableAlias := 'A'
	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("SELECT")`)
	bb.Line(`sql.Fields("","`, string(tableAlias), `", `, table.Title, `QueryFields(`, table.Initials, `.colSet))`)
	if table.NumUniqueFKs > 0 {
		bb.Line("if ", table.Initials, ".withJoin {")
		{
			selectJoinFields(bb, conf, table, &tableAlias, table.ForeignKeys)
			bb.Line(`sql.Append("FROM `, schema.Name, ".", table.Name, ` A")`)
			tableAlias := 'A'
			selectJoinTable(bb, conf, table, 'A', &tableAlias, table.ForeignKeys)
		}
		bb.Line("} else {")
	}
	{
		bb.Line(`sql.Append("FROM `, schema.Name, ".", table.Name, ` A")`)
	}
	if table.NumUniqueFKs > 0 {
		bb.Line("}")
	}

	bb.Line("if ", table.Initials, `.where != "" {`)
	bb.Line(`sql.Append("WHERE", `, table.Initials, ".where)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.groupBy != "" {`)
	bb.Line(`sql.Append("GROUP BY", `, table.Initials, ".groupBy)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.orderBy != "" {`)
	bb.Line(`sql.Append("ORDER BY", `, table.Initials, ".orderBy)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.limit > 0 {`)
	bb.Line(`sql.AppendRaw("LIMIT ", `, table.Initials, ".limit)")
	bb.Line("if ", table.Initials, `.offset > 0 {`)
	bb.Line(`sql.AppendRaw(",", `, table.Initials, ".offset)")
	bb.Line("}")
	bb.Line("}")
	bb.Line("return sql")
	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func selectJoinTableMultiTenant(bb *GenBuffer, conf *Config, table *Table, refAlias rune, tableAlias *rune, fks []*ForeignKey) {
	for _, fk := range fks {
		*tableAlias++
		if fk.IsUnique {
			fkSchema := conf.getSchema(fk.RefSchema)
			if fkSchema.Name == "now00002" {
				bb.S(`sql.AppendRaw(`, table.Initials, ".joinType", `," JOIN ", fmt.Sprintf("now%05d", `, table.Initials, `.companyID), "`, ".", fk.RefTable, " ", string(*tableAlias), " ON (")
			} else {
				bb.S(`sql.Append(`, table.Initials, ".joinType", `," JOIN `, fkSchema.Name, ".", fk.RefTable, " ", string(*tableAlias), " ON (")
			}
			for i, f := range fk.Fields {
				if i > 0 {
					bb.S(" AND ")
				}
				bb.S(string(refAlias), ".", f, " = ", string(*tableAlias), ".", fk.RefFields[i])
			}
			bb.Line(`)")`)

			if len(fk.ForeignKeys) > 0 {
				selectJoinTableMultiTenant(bb, conf, table, *tableAlias, tableAlias, fk.ForeignKeys)
			}
		}
	}
}

// selectSQL generates general SELECT Statement with optional JOINs based on foreign key definitions
func genSelectSQLMultiTenant(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Func(table.StoreReceiver, "selectStatement")
	bb.FuncParams()
	bb.FuncReturn("*SQLStatement")

	tableAlias := 'A'
	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("SELECT")`)
	bb.Line(`sql.Fields("","`, string(tableAlias), `", `, table.Title, `QueryFields(`, table.Initials, `.colSet))`)
	if table.NumUniqueFKs > 0 {
		bb.Line("if ", table.Initials, ".withJoin {")
		{
			selectJoinFields(bb, conf, table, &tableAlias, table.ForeignKeys)
			bb.Line(`sql.AppendRaw("FROM ", fmt.Sprintf("now%05d", `, table.Initials, `.companyID), ".`, table.Name, ` A ")`)
			tableAlias := 'A'
			selectJoinTableMultiTenant(bb, conf, table, 'A', &tableAlias, table.ForeignKeys)
		}
		bb.Line("} else {")
	}
	{
		bb.Line(`sql.AppendRaw("FROM ", fmt.Sprintf("now%05d", `, table.Initials, `.companyID), ".`, table.Name, ` A ")`)
	}
	if table.NumUniqueFKs > 0 {
		bb.Line("}")
	}

	bb.Line("if ", table.Initials, `.where != "" {`)
	bb.Line(`sql.Append("WHERE", `, table.Initials, ".where)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.groupBy != "" {`)
	bb.Line(`sql.Append("GROUP BY", `, table.Initials, ".groupBy)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.orderBy != "" {`)
	bb.Line(`sql.Append("ORDER BY", `, table.Initials, ".orderBy)")
	bb.Line("}")
	bb.Line("if ", table.Initials, `.limit > 0 {`)
	bb.Line(`sql.AppendRaw("LIMIT ", `, table.Initials, ".limit)")
	bb.Line("if ", table.Initials, `.offset > 0 {`)
	bb.Line(`sql.AppendRaw(",", `, table.Initials, ".offset)")
	bb.Line("}")
	bb.Line("}")
	bb.Line("return sql")
	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genJSON(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())

	bb.Line("// ToJSON writes a single object to the buffer.")
	bb.Line("// nolint[gocylco]")
	bb.Func(table.StoreReceiver, "ToJSON")
	bb.FuncParams("t *buffer.TemplateBuffer", "data *"+table.Title)
	bb.FuncReturn("")
	bb.Line(`prepend := "{"`)
	lenFields := len(table.Fields) - 1
	for i, f := range table.Fields {
		bb.Line("if ", table.Initials, ".colSet == nil || ", table.Initials, ".colSet.Bit(", table.Title+"_"+f.Title, ") == 1 {")
		bb.Line("t.", f.jsonFunc, `(prepend, "`, strings.ToLower(f.Name), `", data.`, f.Title, ")")
		if i != lenFields {
			bb.Line(`prepend = ","`)
		}
		bb.Line("}")
	}
	bb.Line("t.S(`}`)")
	bb.FuncEnd()

	bb.Line("// ToJSONArray writes a slice to the named array.")
	bb.Func(table.StoreReceiver, "ToJSONArray")
	bb.FuncParams("w io.Writer", "data []*"+table.Title, "name string")
	bb.FuncReturn("")
	bb.Line(`t := buffer.NewTemplateBuffer()`)
	bb.Line("t.SS(`{\"`, name, `\":[`)")
	bb.S(`for i := range data {
		if i > 0 {
			t.S(",")
		}
		`, table.Initials, `.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}`)
	bb.NewLine()

	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genIndex(conf *Config, schema *Schema, table *Table, index *Index) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	var funcName string
	arrayType := ""
	if index.IsUnique {
		funcName = "OneBy"
		for i, f := range index.Fields {
			if i > 0 {
				funcName += "And"
			}
			funcName += table.Fields[table.FieldMapping[f]].Title
		}
		bb.Line("// ", funcName, " retrieves a row from '", schema.Name, ".", table.Name, "' as a ", table.Title, ".")
	} else {
		// Query for slice
		arrayType = "[]"
		funcName = "QueryBy"
		for i, f := range index.Fields {
			if i > 0 {
				funcName += "And"
			}
			funcName += table.Fields[table.FieldMapping[f]].Title
		}
		bb.Line("// ", funcName, " retrieves multiple rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.Title, `.`)
	}
	bb.Line("//")
	bb.Line("// Generated from index '", index.Name, "'.")
	bb.Line("// nolint[goconst]")
	bb.Func(table.StoreReceiver, funcName)
	bb.S("(")
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
		bb.S(" ")
		bb.S(table.Fields[table.FieldMapping[f]].GoType)
	}
	bb.S(") ")
	bb.FuncReturn(arrayType+"*"+conf.RootPackage+"."+table.Title, "error")
	bb.S(table.Initials, `.where = "`)
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S("A.", table.Fields[table.FieldMapping[f]].Name)
		bb.S(" = ?")
	}
	bb.Line(`"`)
	if index.IsUnique {
		bb.S("return ", table.Initials, ".One(")
	} else {
		bb.S("return ", table.Initials, ".Query(")
	}
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
	}
	bb.Line(")")

	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genTruncate(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// Truncate deletes all rows from ", table.Title, `.`)
	bb.Func(table.StoreReceiver, "Truncate")
	bb.FuncParams()
	bb.FuncReturn("error")

	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("TRUNCATE `, schema.Name, ".", table.Name, `")`)
	bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {`)
	bb.Line(`log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Truncate").Str("stmt", sql.String()).Msg("sql")`)
	bb.Line("}")
	bb.Line("_, err := ", table.Initials, ".db.Exec(sql.Query())")
	bb.Line("if err != nil {")
	bb.Line("log.Error().Err(err).Msg(\"exec\")")
	bb.Line("}")
	bb.Line("return err")
	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genOneSelect(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// One retrieves a row from '", schema.Name, ".", table.Name, "' as a ", table.Title, " with possible joined data.")
	bb.Func(table.StoreReceiver, "One")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("*"+conf.RootPackage+"."+table.Title, "error")

	bb.Line("data := &", table.Title, "{}")
	bb.NewLine()

	bb.Line("err := ", table.Initials, ".one(data, ", table.Initials, ".selectStatement(), args...)")
	bb.Line("if err != nil {")
	bb.Line("log.Error().Err(err).Msg(\"query one\")")
	bb.Line("return nil, err")
	bb.Line("}")
	bb.Line("return &data.", table.Title, ", nil")
	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genQuerySelect(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// Query retrieves many rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.Title, " with possible joined data.")
	bb.Func(table.StoreReceiver, "Query")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("[]*"+conf.RootPackage+"."+table.Title, "error")

	bb.Line("stmt := ", table.Initials, ".selectStatement()")
	bb.Line("return ", table.Initials, ".QueryCustom(stmt.Query(), args...)")
	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genQueryCustom(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// QueryCustom retrieves many rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.Title, " with possible joined data.")
	bb.Func(table.StoreReceiver, "QueryCustom")
	bb.FuncParams("stmt string", "args ...interface{}")
	bb.FuncReturn("[]*"+conf.RootPackage+"."+table.Title, "error")

	bb.Line("dto := &", table.Title, "{}")
	bb.Line("data := &", table.Title, "Slice{}")

	bb.Line("err := ", table.Initials, ".queryCustom(data, dto, stmt, args...)")
	bb.Line("if err != nil {")
	bb.Line(`log.Error().Err(err).Msg("querycustom")`)
	bb.Line("return nil, err")
	bb.Line("}")

	bb.Line(`retValues := make([]*`, conf.RootPackage, ".", table.Title, ", len(data.data))")
	bb.Line("for i := range data.data {")
	bb.Line("	retValues[i] = &data.data[i].", table.Title)
	bb.Line("}")
	bb.Line("return retValues, nil")

	bb.FuncEnd()
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genForeign(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	for _, fk := range table.ForeignKeys {
		fkRefTable := strings.Title(fk.RefTable)
		fkTable := table
		fkSchema := conf.getSchema(fk.RefSchema)
		if t := fkSchema.getTable(fk.RefTable); t != nil {
			fkRefTable = t.Title
			fkTable = t
		}
		if fk.CustomName == "" {
			fk.CustomName = table.Title + strings.Replace(fk.Name, "fk", "", 1)
		}

		bb.Line("// Get", fk.CustomName, " fetches a record from referenced table '", fk.RefTable, "'.")
		bb.Func(table.Receiver, "Get"+fk.CustomName)
		bb.FuncParams("db Execer")
		bb.FuncReturn("error")
		bb.Line("if ", table.Initials, ".", fk.CustomName, " == nil {")
		bb.Line("var err error")
		bb.S(table.Initials, ".", fk.CustomName, ",err = New", fkRefTable, "Store(db).")
		var funcName string
		if fk.IsUnique {
			funcName = "OneBy"
		} else {
			funcName = "QueryBy"
		}
		for i := range fk.Fields {
			if i > 0 {
				funcName += "And"
			}
			funcName += fkTable.Fields[fkTable.FieldMapping[fk.RefFields[i]]].Title
		}
		bb.S(funcName, "(")
		for i := range fk.Fields {
			if i > 0 {
				bb.S(",")
			}
			bb.S(table.Initials, ".", table.Fields[table.FieldMapping[fk.Fields[i]]].Title)
		}
		bb.Line(")")

		bb.Line("return err")
		bb.Line("}")
		bb.Line("return nil")
		bb.FuncEnd()
	}

	for _, fk := range table.ForeignKeys {
		if !fk.IsUnique {
			fkRefTable := strings.Title(fk.RefTable)
			fkTable := table
			fkSchema := conf.getSchema(fk.RefSchema)
			if t := fkSchema.getTable(fk.RefTable); t != nil {
				fkRefTable = t.Title
				fkTable = t
			}
			if fk.CustomName == "" {
				fk.CustomName = table.Title + strings.Replace(fk.Name, "fk", "", 1)
			}

			bb.Line("// EagerFetch", fk.CustomName, " eagerly fetches N records from referenced table '", fk.RefTable, "'.")
			bb.Func(table.StoreReceiver, "EagerFetch"+fk.CustomName)
			bb.FuncParams("data []*" + table.Title)
			bb.FuncReturn("error")

			if len(fk.RefFields) > 1 {
				panic("too many ref fields")
			}

			bb.Line(`stmt := NewSQLStatement()`)
			bb.Line(`stmt.Append("`, fk.RefFields[0], ` IN (")`)

			bb.Line(`for i, d := range data {`)
			bb.Line(`if i > 0 {`)
			bb.Line(`stmt.Append(",")`)
			bb.Line(`}`)
			bb.Line(`stmt.AppendInt(d.`, table.Fields[table.FieldMapping[fk.Fields[0]]].Title, `)`)
			bb.Line(`}`)
			bb.Line(`stmt.Append(")")`)

			bb.Line(`details, err := New`, fkRefTable, `Store(`, table.Initials, `.db).Where(stmt.Query()).OrderBy("A.`, fk.RefFields[0], " DESC, A.", fk.Fields[0], ` DESC").Query()`)
			bb.Line(`if err != nil {`)
			bb.Line(`log.Error().Err(err).Msg("fetch details")`)
			bb.Line(`return err`)
			bb.Line(`}`)

			bb.Line(`for i := range data {`)
			bb.Line(`for j := len(details) - 1; j >= 0; j-- {`)
			bb.Line(`if details[j].`, fkTable.Fields[fkTable.FieldMapping[fk.RefFields[0]]].Title, ` == data[i].`, table.Fields[table.FieldMapping[fk.Fields[0]]].Title, ` {`)
			bb.Line(`data[i].`, fk.CustomName, ` = append(data[i].`, fk.CustomName, `, details[j])`)
			bb.Line(`details = append(details[:j], details[j+1:]...)`)
			bb.Line(`}`)
			bb.Line(`}`)

			bb.Line("}")
			bb.Line("return nil")
			bb.FuncEnd()
		}
	}
	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genUpsert(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())

	// generate upsert statement
	bb.Line("// ", table.lower, "UpsertStmt helper for generating Upserts general statement")
	bb.Line("// nolint[gocyclo]")
	bb.Func(table.StoreReceiver, table.lower+"UpsertStmt")
	bb.FuncParams()
	bb.FuncReturn("*sdb.UpsertStatement")

	bb.Line("upsert := []string{}")
	{
		for _, f := range table.otherFields {
			if !contains(table.Ignores.Upsert, f.Name) {
				bb.Line(`if `, table.Initials, `.colSet == nil || `, table.Initials, `.colSet.Bit(`, table.Title+"_"+f.Title, `) == 1 {`)
				bb.Line(`upsert = append(upsert, "`, f.Name, " = VALUES(", f.Name, `)")`)
				bb.Line("}")
			}
		}
	}

	bb.Line("sql := &sdb.UpsertStatement{}")
	bb.Line(`sql.InsertInto("`, schema.Name, ".", table.Name, `")`)
	bb.S(`sql.Columns(`)
	{
		for _, f := range table.Fields {
			bb.S(`"`)
			bb.S(strings.ToLower(f.Name))
			bb.S(`",`)
		}
	}
	bb.Line(`)`)
	bb.Line(`sql.OnDuplicateKeyUpdate(upsert)`)
	bb.Line(`return sql`)
	bb.Line(`}`)

	// Upsert for a single record
	bb.Line("// UpsertOne inserts the ", table.Title, " to the database.")
	bb.Func(table.StoreReceiver, "UpsertOne")
	bb.FuncParams("data *" + table.Title)
	bb.FuncReturn("int64", "error")
	bb.Line("return ", table.Initials, ".Upsert([]*", table.Title, "{data})")
	bb.Line("}")

	// upsert for data array
	bb.Line("// Upsert executes upsert for array of ", table.Title)
	bb.Func(table.StoreReceiver, "Upsert")
	bb.FuncParams("data []*" + table.Title)
	bb.FuncReturn("int64", "error")

	bb.Line(`sql := `, table.Initials, ".", table.lower, `UpsertStmt()
	
	for _, d := range data {
		sql.Record(d)
	}

	if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`, table.Title, `Upsert").Str("stmt", sql.String()).Msg("sql")
	}
	res, err := `, table.Initials, `.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return -1, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("rowsaffected")
		return -1, err
	}

	return affected, nil
}
`)
	s := string(bb.Bytes())
	bb.Free()
	return s

}

func genInsert(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// Insert inserts the ", table.Title, ` to the database.`)
	bb.Func(table.StoreReceiver, "Insert")
	bb.FuncParams("data *" + conf.RootPackage + "." + table.Title)
	bb.FuncReturn("error")
	bb.Line("var err error")

	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("INSERT INTO `, schema.Name, ".", table.Name, ` (")`)
	bb.Line("fields := ", table.Title, `QueryFields(`, table.Initials, `.colSet)`)
	bb.Line(`sql.Fields("","", fields)`)
	bb.Line(`sql.Append(") VALUES (")`)
	bb.Line("for i := range fields {")
	bb.Line("if i > 0 {")
	bb.Line(`sql.Append(",")`)
	bb.Line(`}`)
	bb.Line(`sql.Append("?")`)
	bb.Line(`}`)
	bb.Line(`sql.Append(")")`)
	bb.NewLine()

	bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {`)
	bb.S(`log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Insert").Str("stmt", sql.String()).`)
	bb.Log(table.Fields, "data")
	bb.Line(".Msg(\"sql\") }")

	if table.Fields[table.FieldMapping["id"]].IsPrimaryKey && table.Fields[table.FieldMapping["id"]].IsAutoincrement {
		bb.S("res, err :=")
	} else {
		bb.S("_, err =")
	}
	bb.S(" ", table.Initials, `.db.Exec(sql.Query(),`)
	for i, f := range table.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.", f.Title)
	}
	bb.S(`)	
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}`)

	// only do this on AutoIncrement fields
	if table.Fields[table.FieldMapping["id"]].IsPrimaryKey && table.Fields[table.FieldMapping["id"]].IsAutoincrement {
		bb.S(`
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("lastinsertid")
		return err
	}

	// set primary key and existence
	`)
		bb.S("data.ID = ")
		if table.Fields[table.FieldMapping["id"]].GoType == "int64" {
			bb.S("id")
		} else {
			bb.S("int(id)")
		}
		bb.S(`
	`)
	}
	bb.S(`
	return nil
}
`)

	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genInsertMultiTenant(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// Insert inserts the ", table.Title, ` to the database.`)
	bb.Func(table.StoreReceiver, "Insert")
	bb.FuncParams("data *" + conf.RootPackage + "." + table.Title)
	bb.FuncReturn("error")
	bb.Line("var err error")

	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.AppendRaw("INSERT INTO ", fmt.Sprintf("now%05d", `, table.Initials, `.companyID),".`, table.Name, ` (")`)
	bb.Line("fields := ", table.Title, `QueryFields(`, table.Initials, `.colSet)`)
	bb.Line(`sql.Fields("","", fields)`)
	bb.Line(`sql.Append(") VALUES (")`)
	bb.Line("for i := range fields {")
	bb.Line("if i > 0 {")
	bb.Line(`sql.Append(",")`)
	bb.Line(`}`)
	bb.Line(`sql.Append("?")`)
	bb.Line(`}`)
	bb.Line(`sql.Append(")")`)
	bb.NewLine()

	bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {`)
	bb.S(`log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Insert").Str("stmt", sql.String()).`)
	bb.Log(table.Fields, "data")
	bb.Line(".Msg(\"sql\") }")

	if table.Fields[table.FieldMapping["id"]].IsPrimaryKey && table.Fields[table.FieldMapping["id"]].IsAutoincrement {
		bb.S("res, err :=")
	} else {
		bb.S("_, err =")
	}
	bb.S(" ", table.Initials, `.db.Exec(sql.Query(),`)
	for i, f := range table.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.", f.Title)
	}
	bb.S(`)	
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return err
	}`)

	// only do this on AutoIncrement fields
	if table.Fields[table.FieldMapping["id"]].IsPrimaryKey && table.Fields[table.FieldMapping["id"]].IsAutoincrement {
		bb.S(`
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("lastinsertid")
		return err
	}

	// set primary key and existence
	`)
		bb.S("data.ID = ")
		if table.Fields[table.FieldMapping["id"]].GoType == "int64" {
			bb.S("id")
		} else {
			bb.S("int(id)")
		}
		bb.S(`
	`)
	}
	bb.S(`
	return nil
}
`)

	s := string(bb.Bytes())
	bb.Free()
	return s
}

func genUpdate(conf *Config, schema *Schema, table *Table) string {
	if len(table.otherFields) > 0 {
		bb := NewGenBuffer(bytebufferpool.Get())
		bb.Line(`// Update updates the `, table.Title, ` in the database.`)
		bb.Line("// nolint[gocyclo]")
		bb.Func(table.StoreReceiver, "Update")
		bb.FuncParams("data *" + conf.RootPackage + "." + table.Title)
		bb.FuncReturn("int64", "error")
		bb.Line("sql := NewSQLStatement()")
		bb.Line("var prepend string")
		bb.Line("args := []interface{}{}")
		bb.Line(`sql.Append("UPDATE `, schema.Name, ".", table.Name, ` SET")`)
		for i, f := range table.otherFields {
			bb.Line(`if `, table.Initials, `.colSet == nil || `, table.Initials, `.colSet.Bit(`, table.Title+"_"+f.Title, `) == 1 {`)
			bb.Line(`sql.AppendRaw(prepend, "`, f.Name, ` = ?")`)
			if i+1 != len(table.otherFields) {
				bb.Line(`prepend = ","`)
			}
			bb.Line(`args = append(args, data.`, f.Title, `)`)
			bb.Line("}")
		}
		bb.S(`sql.Append(" WHERE `)
		for i, f := range table.pkFields {
			if i > 0 {
				bb.S(" AND ")
			}
			bb.S(f.Name)
			bb.S(" = ?")
		}
		bb.S(`")`)
		bb.NewLine()

		for _, f := range table.pkFields {
			bb.Line(`args = append(args, data.`, f.Title, `)`)
		}

		bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Update").Str("stmt", sql.String()). Interface("args", args).Msg("sql")
	}
	res, err := `)
		bb.S(table.Initials, `.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}
`)
		s := string(bb.Bytes())
		bb.Free()
		return s
	}
	return ""
}

func genUpdateMultiTenant(conf *Config, schema *Schema, table *Table) string {
	if len(table.otherFields) > 0 {
		bb := NewGenBuffer(bytebufferpool.Get())
		bb.Line(`// Update updates the `, table.Title, ` in the database.`)
		bb.Line("// nolint[gocyclo]")
		bb.Func(table.StoreReceiver, "Update")
		bb.FuncParams("data *" + conf.RootPackage + "." + table.Title)
		bb.FuncReturn("int64", "error")
		bb.Line("sql := NewSQLStatement()")
		bb.Line("var prepend string")
		bb.Line("args := []interface{}{}")
		bb.Line(`sql.AppendRaw("UPDATE ", fmt.Sprintf("now%05d", `, table.Initials, `.companyID),".`, table.Name, ` SET ")`)
		for i, f := range table.otherFields {
			bb.Line(`if `, table.Initials, `.colSet == nil || `, table.Initials, `.colSet.Bit(`, table.Title+"_"+f.Title, `) == 1 {`)
			bb.Line(`sql.AppendRaw(prepend, "`, f.Name, ` = ?")`)
			if i+1 != len(table.otherFields) {
				bb.Line(`prepend = ","`)
			}
			bb.Line(`args = append(args, data.`, f.Title, `)`)
			bb.Line("}")
		}
		bb.S(`sql.Append(" WHERE `)
		for i, f := range table.pkFields {
			if i > 0 {
				bb.S(" AND ")
			}
			bb.S(f.Name)
			bb.S(" = ?")
		}
		bb.S(`")`)
		bb.NewLine()

		for _, f := range table.pkFields {
			bb.Line(`args = append(args, data.`, f.Title, `)`)
		}

		bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Update").Str("stmt", sql.String()). Interface("args", args).Msg("sql")
	}
	res, err := `)
		bb.S(table.Initials, `.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}
`)
		s := string(bb.Bytes())
		bb.Free()
		return s
	}
	return ""
}

func genDelete(conf *Config, schema *Schema, table *Table) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Line("// Delete deletes the ", table.Title, ` from the database.`)
	bb.Func(table.StoreReceiver, "Delete")
	bb.FuncParams("data *" + conf.RootPackage + "." + table.Title)
	bb.FuncReturn("error")
	bb.S(`var err error

	sql := NewSQLStatement()
	sql.Append("DELETE FROM `)
	bb.S(schema.Name)
	bb.S(".")
	bb.S(table.Name)
	bb.S(` WHERE")
	sql.Append("`)
	for i, f := range table.pkFields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S(f.Name)
		bb.S(" = ?")
	}
	bb.S(`")
	
	if  zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "`)
	bb.S(schema.Name)
	bb.S(".")
	bb.S(table.Name)
	bb.S(`.Delete").Str("stmt", sql.String()).`)
	bb.Log(table.pkFields, "data")
	bb.S(`.Msg("sql")
}`)
	bb.NewLine()

	bb.Line("_, err = ", table.Initials, `.db.Exec(sql.Query(), `)
	for i, f := range table.pkFields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.")
		bb.S(f.Title)
	}
	bb.S(`)
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)

	// DeleteSlice only with single int primary keys
	if len(table.pkFields) == 1 && table.pkFields[0].GoType == "int" {
		bb.Line("// DeleteSlice delets all slice element from the database.")
		bb.Func(table.StoreReceiver, "DeleteSlice")
		bb.FuncParams("data []*" + table.Title)
		bb.FuncReturn("error")
		bb.S(`var err error

	sql := NewSQLStatement()
	sql.Append("DELETE FROM `)
		bb.S(schema.Name)
		bb.S(".")
		bb.S(table.Name)
		bb.S(` WHERE")`)
		bb.NewLine()

		bb.Line(`sql.AppendRaw("`, table.pkFields[0].Name, ` IN (")`)
		bb.Line(`for i := range data {`)
		bb.Line(`if i > 0 {`)
		bb.Line(`sql.AppendRaw(",")`)
		bb.Line(`}`)
		bb.Line(`sql.AppendInt(data[i].`, table.pkFields[0].Title, ")")
		bb.Line(`}`)
		bb.Line(`sql.Append(")")`)

		bb.S(`if  zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Debug().Str("fn", "`)
		bb.S(schema.Name)
		bb.S(".")
		bb.S(table.Name)
		bb.S(`.DeleteSlice").Str("stmt", sql.String()).Msg("sql")
	}`)
		bb.NewLine()

		bb.Line("_, err = ", table.Initials, `.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)
	}

	bb.Line("// DeleteByQuery uses a where condition to delete entries.")
	bb.Func(table.StoreReceiver, "DeleteByQuery")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("error")
	bb.Line(`var err error`)
	bb.Line(`sql := NewSQLStatement()`)
	bb.Line(`sql.Append("DELETE FROM `, schema.Name, ".", table.Name, `")`)
	bb.Line(`if `, table.Initials, `.where == "" {`)
	bb.Line(`return errors.New("no where condition set")`)
	bb.Line(`}`)

	bb.Line(`sql.Append("WHERE", `, table.Initials, `.where)`)
	bb.S(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`)
	bb.S(schema.Name)
	bb.S(".")
	bb.S(table.Name)
	bb.Line(`.DeleteByQuery").Str("stmt", sql.String()). Interface("args", args).Msg("sql")`)
	bb.Line(`}`)
	bb.NewLine()

	bb.Line("_, err = ", table.Initials, `.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)
	s := string(bb.Bytes())
	bb.Free()
	return s
}
