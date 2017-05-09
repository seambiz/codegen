package codegen

import "strings"

// TInsert template
func TInsert(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Insert inserts the ", table.title, ` to the database.`)
	bb.Func(table.storeReceiver, "Insert")
	bb.FuncParams("data *" + table.title)
	bb.FuncReturn("error")
	bb.Line("var err error")

	bb.Line("sql := sdb.NewSQLStatement()")
	bb.Line(`sql.Append("INSERT INTO `, schema.Name, ".", table.Name, ` (")`)
	bb.Line(`sql.Fields("","", `, strings.ToLower(table.title), `QueryFields)`)
	bb.S(`sql.Append(") VALUES (`)
	for i := range table.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("?")
	}
	bb.Line(`)")`)
	bb.NewLine()

	bb.Line(`if logging.LogDB.Check(zap.DebugLevel, "") != nil {`)
	bb.S(`logging.LogDB.Debug("`, schema.Name, ".", table.Name, `.Insert", zap.String("stmt", sql.String()), `)
	bb.Log(table.Fields, "data")
	bb.Line(`)`)
	bb.Line("}")

	if table.Fields[table.fieldMapping["id"]].IsPrimaryKey && table.Fields[table.fieldMapping["id"]].IsAutoincrement {
		bb.S("res, err :=")
	} else {
		bb.S("_, err =")
	}
	bb.S(" ", table.initials, `.db.Exec(sql.Query(),`)
	for i, f := range table.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.", f.title)
	}
	bb.S(`)	
	if err != nil {
		logging.SQLError(err)
		return err
	}`)

	// only do this on AutoIncrement fields
	if table.Fields[table.fieldMapping["id"]].IsPrimaryKey && table.Fields[table.fieldMapping["id"]].IsAutoincrement {
		bb.S(`
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		logging.SQLError(err)
		return err
	}

	// set primary key and existence
	`)
		bb.S("data.ID = ")
		if table.Fields[table.fieldMapping["id"]].goType == "int64" {
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

}
