package codegen

// TDelete template
func TDelete(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Delete deletes the ", table.title, ` from the database.`)
	bb.Func(table.storeReceiver, "Delete")
	bb.FuncParams("data *" + table.title)
	bb.FuncReturn("error")
	bb.S(`var err error

	sql := sdb.NewSQLStatement()
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
	
	if logging.LogDB.Check(zap.DebugLevel, "") != nil {
		logging.LogDB.Debug("`)
	bb.S(schema.Name)
	bb.S(".")
	bb.S(table.Name)
	bb.S(`.Delete", zap.String("stmt", sql.String()), `)
	bb.Log(table.pkFields, "data")
	bb.S(`)
}`)
	bb.NewLine()

	bb.Line("_, err = ", table.initials, `.db.Exec(sql.Query(), `)
	for i, f := range table.pkFields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.")
		bb.S(f.title)
	}
	bb.S(`)
	if err != nil {
		logging.SQLError(err)
	}

	return err
}
`)
}
