package codegen

// TUpdate template
func TUpdate(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line(`// Update updates the `, table.title, ` in the database.`)
	bb.Func(table.storeReceiver, "Update")
	bb.FuncParams("data *" + table.title)
	bb.FuncReturn("int64", "error")
	bb.Line("sql := sdb.NewSQLStatement()")
	bb.Line(`sql.Append("UPDATE `, schema.Name, ".", table.Name, ` SET")`)
	bb.S(`sql.Append("`)

	for i, f := range table.otherFields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(f.Name)
		bb.S(" = ?")
	}
	bb.S(`")
	sql.Append("WHERE `)
	for i, f := range table.pkFields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S(f.Name)
		bb.S(" = ?")
	}
	bb.S(`")
	
	if logging.LogDB.Check(zap.DebugLevel, "") != nil {
		logging.LogDB.Debug("`, schema.Name, ".", table.Name, ".Update\", ", `zap.String("stmt", sql.String()), `)
	bb.Log(table.Fields, "data")
	bb.S(`)
	}
	res, err := `)
	bb.S(table.initials, `.db.Exec(sql.Query(), `)
	for i, f := range table.otherFields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S("data.", f.title)
	}
	for _, f := range table.pkFields {
		bb.S(", data.", f.title)
	}
	bb.S(`)
	if err != nil {
		logging.SQLError(err)
		return 0, err
	}
	return res.RowsAffected()
}
`)
	/*
	   	bb.Line(`// UpdatePartial updates the `, table.title, ` in the database.`)
	   	bb.Func(table.storeReceiver, "UpdatePartial")
	   	bb.FuncParams("data *" + table.title)
	   	bb.FuncReturn("int64", "error")
	   	bb.Line("sql := sdb.NewSQLStatement()")
	   	bb.Line(`sql.Append("UPDATE `, schema.Name, ".", table.Name, ` SET")`)
	   	for i, f := range table.otherFields {
	   		bb.Line("if !", table.initials, ".null.Get(", strconv.Itoa(i+len(table.pkFields)), ") {")
	   		bb.S(`sql.Append("`, f.Name, ` = ?`)
	   		if i < len(table.otherFields) {
	   			bb.S(", ")
	   		}
	   		bb.Line(`")`)
	   	}
	   	bb.S(`sql.Append("WHERE `)
	   	for i, f := range table.pkFields {
	   		if i > 0 {
	   			bb.S(" AND ")
	   		}
	   		bb.S(f.Name)
	   		bb.S(" = ?")
	   	}
	   	bb.S(`")

	   	if logging.LogDB.Check(zap.DebugLevel, "") != nil {
	   		logging.LogDB.Debug("`, schema.Name, ".", table.Name, ".Update\", ", `zap.String("stmt", sql.String()), `)
	   	bb.Log(table.Fields, "data")
	   	bb.S(`)
	   	}
	   	res, err := `)
	   	bb.S(table.initials, `.db.Exec(sql.Query(), `)
	   	for i, f := range table.otherFields {
	   		if i > 0 {
	   			bb.S(", ")
	   		}
	   		bb.S("data.", f.title)
	   	}
	   	for _, f := range table.pkFields {
	   		bb.S(", data.", f.title)
	   	}
	   	bb.S(`)
	   	if err != nil {
	   		logging.SQLError(err)
	   		return 0, err
	   	}
	   	return res.RowsAffected()
	   }
	   `)*/
}
