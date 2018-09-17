package codegen

// TTruncate template
func TTruncate(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Truncate deletes all rows from ", table.title, `.`)
	bb.Func(table.storeReceiver, "Truncate")
	bb.FuncParams()
	bb.FuncReturn("error")

	bb.Line("sql := sdb.NewSQLStatement()")
	bb.Line(`sql.Append("TRUNCATE `, schema.Name, ".", table.Name, `")`)
	bb.Line(`if  zerolog.GlobalLevel() ==  zerolog.DebugLevel {`)
	bb.Line(`log.Debug().Str("fn", "`, schema.Name, ".", table.Name, `.Truncate").Str("stmt", sql.String()).Msg("sql")`)
	bb.Line("}")
	bb.Line("_, err := ", table.initials, ".db.Exec(sql.Query())")
	bb.Line("if err != nil {")
	bb.Line("log.Error().Err(err).Msg(\"exec\")")
	bb.Line("}")
	bb.Line("return err")
	bb.FuncEnd()
}
