package codegen

// TTruncate template
func TTruncate(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Truncate deletes all rows from ", table.title, `.`)
	bb.Func(table.storeReceiver, "Truncate")
	bb.FuncParams()
	bb.FuncReturn("error")

	bb.Line("sql := sdb.NewSQLStatement()")
	bb.Line(`sql.Append("TRUNCATE `, schema.Name, ".", table.Name, `")`)
	bb.Line(`if logging.LogDB.Check(zap.DebugLevel, "") != nil {`)
	bb.Line(`logging.LogDB.Debug("`, schema.Name, ".", table.Name, `.Truncate", zap.String("stmt", sql.String()))`)
	bb.Line("}")
	bb.Line("_, err := ", table.initials, ".db.Exec(sql.Query())")
	bb.Line("if err != nil {")
	bb.Line("logging.SQLError(err)")
	bb.Line("}")
	bb.Line("return err")
	bb.FuncEnd()
}
