package codegen

// TDelete template
func TDelete(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Delete deletes the ", table.title, ` from the database.`)
	bb.Func(table.storeReceiver, "Delete")
	bb.FuncParams("data *" + table.title)
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

	bb.Line("_, err = ", table.initials, `.db.Exec(sql.Query(), `)
	for i, f := range table.pkFields {
		if i > 0 {
			bb.S(".")
		}
		bb.S("data.")
		bb.S(f.title)
	}
	bb.S(`)
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)

	// DeleteSlice only with single int primary keys
	if len(table.pkFields) == 1 && table.pkFields[0].goType == "int" {
		bb.Line("// DeleteSlice delets all slice element from the database.")
		bb.Func(table.storeReceiver, "DeleteSlice")
		bb.FuncParams("data []*" + table.title)
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
		bb.Line(`sql.AppendInt(data[i].`, table.pkFields[0].title, ")")
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

		bb.Line("_, err = ", table.initials, `.db.Exec(sql.Query())
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)
	}

	bb.Line("// DeleteByQuery uses a where condition to delete entries.")
	bb.Func(table.storeReceiver, "DeleteByQuery")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("error")
	bb.Line(`var err error`)
	bb.Line(`sql := NewSQLStatement()`)
	bb.Line(`sql.Append("DELETE FROM `, schema.Name, ".", table.Name, `")`)
	bb.Line(`if `, table.initials, `.where == "" {`)
	bb.Line(`return errors.New("no where condition set")`)
	bb.Line(`}`)

	bb.Line(`sql.Append("WHERE", `, table.initials, `.where)`)
	bb.S(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {
		log.Debug().Str("fn", "`)
	bb.S(schema.Name)
	bb.S(".")
	bb.S(table.Name)
	bb.Line(`.DeleteByQuery").Str("stmt", sql.String()). Interface("args", args).Msg("sql")`)
	bb.Line(`}`)
	bb.NewLine()

	bb.Line("_, err = ", table.initials, `.db.Exec(sql.Query(), args...)
	if err != nil {
		log.Error().Err(err).Msg("exec")
	}

	return err
}
`)
}
