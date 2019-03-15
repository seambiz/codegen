package codegen

// TUpdate template
func TUpdate(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line(`// Update updates the `, table.Title, ` in the database.`)
	bb.Line("// nolint[gocyclo]")
	bb.Func(table.storeReceiver, "Update")
	bb.FuncParams("data *" + table.Title)
	bb.FuncReturn("int64", "error")
	bb.Line("sql := NewSQLStatement()")
	bb.Line("var prepend string")
	bb.Line("args := []interface{}{}")
	bb.Line(`sql.Append("UPDATE `, schema.Name, ".", table.Name, ` SET")`)
	for i, f := range table.otherFields {
		bb.Line(`if `, table.initials, `.colSet == nil || `, table.initials, `.colSet.Bit(`, table.Title+f.Title, `) == 1 {`)
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
	bb.S(table.initials, `.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
		return 0, err
	}
	return res.RowsAffected()
}
`)
}
