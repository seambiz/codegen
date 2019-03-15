package codegen

// TInsert template
func TInsert(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Insert inserts the ", table.Title, ` to the database.`)
	bb.Func(table.storeReceiver, "Insert")
	bb.FuncParams("data *" + table.Title)
	bb.FuncReturn("error")
	bb.Line("var err error")

	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("INSERT INTO `, schema.Name, ".", table.Name, ` (")`)
	bb.Line("fields := ", table.Title, `QueryFields(`, table.initials, `.colSet)`)
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
	bb.S(" ", table.initials, `.db.Exec(sql.Query(),`)
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

}
