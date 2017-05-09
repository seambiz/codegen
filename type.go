package codegen

import (
	"strconv"
	"strings"
)

// TType template
func TType(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {

	bb.S("var " + strings.ToLower(table.title) + "QueryFields = []string{")
	for i, f := range table.Fields {
		if i > 0 {
			bb.S(",")
		}
		bb.S(`"` + strings.ToLower(f.Name) + `"`)
	}
	bb.Line(`}`)
	bb.NewLine()

	bb.Line("// ", table.title, " represents a row from '", schema.Name, ".", table.Name, "'.")

	bb.Struct(table.title)
	for _, f := range table.Fields {
		bb.Line(f.title, " ", f.goType, " `json:\"", f.Name, `" db:"`, strings.ToLower(f.Name), "\"`")
	}
	if len(table.ForeignKeys) > 0 {
		bb.NewLine()
	}
	for _, fk := range table.ForeignKeys {
		fkRefTable := strings.Title(fk.RefTable)
		fkSchema := conf.getSchema(fk.RefSchema)
		if t := fkSchema.getTable(fk.RefTable); t != nil {
			fkRefTable = t.title
		}
		if fk.CustomName == "" {
			fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
		}

		bb.S(fk.CustomName)
		bb.S(" ")
		if !fk.IsUnique {
			bb.S("[]")
		}
		bb.S("*")
		bb.Line(fkRefTable)
	}
	bb.StructEnd()

	bb.Line("// IsEmpty checks if primary key fields are zero.")
	bb.Func(table.receiver, "IsEmpty")
	bb.FuncParams()
	bb.FuncReturn("bool")
	if len(table.pkFields) == 1 {
		bb.Line("return ", table.initials, ".", table.pkFields[0].title, " == ", table.pkFields[0].goZero)
	} else {
		for _, f := range table.pkFields {
			bb.Line("if ", table.initials, ".", f.title, " != ", f.goZero, " {")
			bb.Line("return false")
			bb.Line("}")
		}
		bb.Line("return true")
	}
	bb.FuncEnd()

	if len(table.ForeignKeys) > 0 {
		bb.Line("// checkJoinFields checks if join was successful and if not resets pointers to nil.")
		bb.Func(table.receiver, "checkJoinFields")
		bb.FuncParams()
		bb.FuncReturn()
		for _, fk := range table.ForeignKeys {
			if fk.IsUnique {
				if fk.CustomName == "" {
					fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
				}

				bb.Line("if ", table.initials, ".", fk.CustomName, ".IsEmpty() {")
				bb.Line(table.initials, ".", fk.CustomName, " = nil")
				bb.Line("}")
			}
		}
		bb.FuncEnd()
	}

	bb.Line("// ", table.store, " is used to query for '", table.title, "' records.")
	bb.Struct(table.store)
	bb.Line("db *sqlx.DB")
	bb.Line("withJoin bool")
	bb.Line("joinType string")
	bb.Line("where string")
	bb.Line("orderby string")
	bb.StructEnd()

	bb.Line("// New", table.title, "Store return DAO Store for ", table.title)
	bb.Func("", "New"+table.title+"Store")
	bb.FuncParams("conn *sqlx.DB")
	bb.FuncReturn("*" + table.store)
	bb.Line(table.initials, " := &", table.store, "{}")
	bb.Line(table.initials, ".db = conn")
	bb.Line(table.initials, ".withJoin = true")
	bb.Line(table.initials, `.joinType = LEFT`)
	bb.Line("return ", table.initials)
	bb.FuncEnd()

	bb.Line("// WithoutJoins won't execute JOIN when querying for records.")
	bb.Func(table.storeReceiver, "WithoutJoins")
	bb.FuncParams()
	bb.FuncReturn("*" + table.store)
	bb.Line(table.initials, ".withJoin = false")
	bb.Line("return ", table.initials)
	bb.FuncEnd()

	bb.Line("// Where sets local sql, that will be appended to SELECT.")
	bb.Func(table.storeReceiver, "Where")
	bb.FuncParams("sql string")
	bb.FuncReturn("*" + table.store)
	bb.Line(table.initials, ".where = sql")
	bb.Line("return ", table.initials)
	bb.FuncEnd()

	bb.Line("// OrderBy sets local sql, that will be appended to SELECT.")
	bb.Func(table.storeReceiver, "OrderBy")
	bb.FuncParams("sql string")
	bb.FuncReturn("*" + table.store)
	bb.Line(table.initials, ".orderby = sql")
	bb.Line("return ", table.initials)
	bb.FuncEnd()

	bb.Line("// JoinType sets join statement type (Default: INNER | LEFT | RIGHT | OUTER).")
	bb.Func(table.storeReceiver, "JoinType")
	bb.FuncParams("t string")
	bb.FuncReturn("*" + table.store)
	bb.Line(table.initials, ".joinType = t")
	bb.Line("return ", table.initials)
	bb.FuncEnd()

	/*
	   wird jetzt mittels bind gemacht
	   	// func fields
	   	bb.Func(table.receiver, "scanFields")
	   	bb.FuncParams("withJoin bool")
	   	bb.FuncReturn("[]interface{}")

	   	bb.Line("f := []interface{}{}")
	   	for _, f := range table.Fields {
	   		bb.Line("f = append(f, &", table.initials, ".", f.title, ")")
	   	}
	   	if len(table.ForeignKeys) > 0 {
	   		bb.Line("if withJoin {")
	   		for _, fk := range table.ForeignKeys {
	   			if fk.IsUnique {
	   				var fkRefTable *Table
	   				fkSchema := conf.getSchema(fk.RefSchema)
	   				if t := fkSchema.getTable(fk.RefTable); t != nil {
	   					fkRefTable = t
	   				}
	   				if fk.CustomName == "" {
	   					fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
	   				}

	   				for _, f := range fkRefTable.Fields {
	   					bb.Line("f = append(f, &", table.initials, ".", fk.CustomName, ".", f.title, ")")
	   				}

	   			}
	   		}
	   	}
	   	bb.Line("}")
	   	bb.Line("return f")
	   	bb.FuncEnd()
	*/

	bind(bb, conf, schema, table)
	selectSQL(bb, conf, schema, table)
	oneSelect(bb, conf, schema, table)
	querySelect(bb, conf, schema, table)
	queryCustom(bb, conf, schema, table)
}

func bind(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	// func fields
	bb.Func(table.receiver, "bind")
	bb.FuncParams("row mysql.Row", "withJoin bool")
	bb.FuncReturn()

	for i, f := range table.Fields {
		bb.Line(table.initials, ".", f.title, " = row.", f.mappingFunc, "(", strconv.Itoa(i), ")")
	}
	if len(table.ForeignKeys) > 0 {
		numFields := len(table.Fields)
		bb.Line("if withJoin {")
		for _, fk := range table.ForeignKeys {
			if fk.IsUnique {
				var fkRefTable *Table
				fkSchema := conf.getSchema(fk.RefSchema)
				if t := fkSchema.getTable(fk.RefTable); t != nil {
					fkRefTable = t
				}
				if fk.CustomName == "" {
					fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
				}

				bb.Line(table.initials, ".", fk.CustomName, "= &", fkRefTable.title, "{}")
				for i, f := range fkRefTable.Fields {
					bb.Line(table.initials, ".", fk.CustomName, ".", f.title, " = row.", f.mappingFunc, "(", strconv.Itoa(i+numFields), ")")
				}
				numFields += len(fkRefTable.Fields)
			}
		}
		bb.Line("}")
	}
	bb.FuncEnd()
}

// selectSQL generates general SELECT Statement with optional JOINs based on foreign key definitions
func selectSQL(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Func(table.storeReceiver, "selectStatement")
	bb.FuncParams()
	bb.FuncReturn("*sdb.SQLStatement")

	tableAlias := 'A'
	bb.Line("sql := sdb.NewSQLStatement()")
	bb.Line(`sql.Append("SELECT")`)
	bb.Line(`sql.Fields("","`, string(tableAlias), `", `, strings.ToLower(table.title), `QueryFields)`)
	bb.Line("if ", table.initials, ".withJoin {")
	{
		for _, fk := range table.ForeignKeys {
			tableAlias++
			if fk.IsUnique {
				fkRefTable := strings.Title(fk.RefTable)
				fkSchema := conf.getSchema(fk.RefSchema)
				if t := fkSchema.getTable(fk.RefTable); t != nil {
					fkRefTable = t.title
				}
				bb.Line(`sql.Fields(",","`, string(tableAlias), `",`, strings.ToLower(fkRefTable), `QueryFields)`)
			}
		}
		bb.Line(`sql.Append("FROM `, schema.Name, ".", table.Name, ` A")`)
		tableAlias := 'A'
		for _, fk := range table.ForeignKeys {
			tableAlias++
			if fk.IsUnique {
				fkSchema := conf.getSchema(fk.RefSchema)
				// TODO Join type
				bb.S(`sql.Append(`, table.initials, ".joinType", `," JOIN `, fkSchema.Name, ".", fk.RefTable, " ", string(tableAlias), " ON (")
				for i, f := range fk.Fields {
					if i > 0 {
						bb.S(" AND ")
					}
					bb.S("A.", f, " = ", string(tableAlias), ".", fk.RefFields[i])
				}
				bb.Line(`)")`)
			}
		}
	}
	bb.Line("} else {")
	{
		bb.Line(`sql.Append("FROM `, schema.Name, ".", table.Name, ` A")`)
	}
	bb.Line("}")

	bb.Line("if ", table.initials, `.where != "" {`)
	bb.Line(`sql.Append("WHERE", `, table.initials, ".where)")
	bb.Line("}")

	bb.Line("if ", table.initials, `.orderby != "" {`)
	bb.Line(`sql.Append("ORDER BY", `, table.initials, ".orderby)")
	bb.Line("}")
	bb.Line("return sql")
	bb.FuncEnd()
}

func oneSelect(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// One retrieves a row from '", schema.Name, ".", table.Name, "' as a ", table.title, " with possible joined data.")
	bb.Func(table.storeReceiver, "One")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("*"+table.title, "error")

	bb.Line("var err error")
	bb.Line("data := ", table.title, "{}")
	bb.NewLine()

	bb.Line("sql := ", table.initials, ".selectStatement()")
	bb.Line(`if logging.LogDB.Check(zap.DebugLevel, "") != nil {`)
	bb.Line(`logging.LogDB.Debug("`, table.store, `.One", zap.String("stmt", sql.String()), zap.Any("args", args))`)
	bb.Line("}")

	bb.Line("stmt, err := Conn.Prepare(sql.Query())")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")

	bb.Line("row, _, err := stmt.ExecFirst(args...)")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")
	bb.Line("if len(row) > 0 {")
	{
		bb.Line("data.bind(row, ", table.initials, ".withJoin)")

		if len(table.ForeignKeys) > 0 {
			bb.Line("if ", table.initials, ".withJoin {")
			bb.Line("data.checkJoinFields()")
			bb.Line("}")
		}
	}
	bb.Line("}")

	bb.Line("return &data, err")
	bb.FuncEnd()
}

func querySelect(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// Query retrieves many rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.title, " with possible joined data.")
	bb.Func(table.storeReceiver, "Query")
	bb.FuncParams("args ...interface{}")
	bb.FuncReturn("[]*"+table.title, "error")

	bb.Line("var err error")
	bb.Line("res := []*", table.title, "{}")
	bb.NewLine()

	bb.Line("sql := ", table.initials, ".selectStatement()")
	bb.Line(`if logging.LogDB.Check(zap.DebugLevel, "") != nil {`)
	bb.Line(`logging.LogDB.Debug("`, table.store, `.Query", zap.String("stmt", sql.String()), zap.Any("args", args))`)
	bb.Line("}")

	bb.Line("stmt, err := Conn.Prepare(sql.Query())")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")

	bb.Line("rows, _, err := stmt.Exec(args...)")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")
	bb.Line("for _, row := range rows {")
	bb.Line("data := ", table.title, "{}")
	bb.Line("data.bind(row, ", table.initials, ".withJoin)")

	if len(table.ForeignKeys) > 0 {
		bb.Line("if ", table.initials, ".withJoin {")
		bb.Line("data.checkJoinFields()")
		bb.Line("}")
	}
	bb.Line("res = append(res, &data)")
	bb.Line("}")

	bb.Line("return res, err")
	bb.FuncEnd()
}

func queryCustom(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	bb.Line("// QueryCustom retrieves many rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.title, " with possible joined data.")
	bb.Func(table.storeReceiver, "QueryCustom")
	bb.FuncParams("sql string", "args ...interface{}")
	bb.FuncReturn("[]*"+table.title, "error")

	bb.Line("var err error")
	bb.Line("res := []*", table.title, "{}")
	bb.NewLine()

	bb.Line(`if logging.LogDB.Check(zap.DebugLevel, "") != nil {`)
	bb.Line(`logging.LogDB.Debug("`, table.store, `.Query", zap.String("stmt", sql), zap.Any("args", args))`)
	bb.Line("}")

	bb.Line("stmt, err := Conn.Prepare(sql)")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")

	bb.Line("rows, _, err := stmt.Exec(args...)")
	bb.Line("if err != nil {")
	bb.Line("return nil, err")
	bb.Line("}")
	bb.Line("for _, row := range rows {")
	bb.Line("data := ", table.title, "{}")
	bb.Line("data.bind(row, ", table.initials, ".withJoin)")

	if len(table.ForeignKeys) > 0 {
		bb.Line("if ", table.initials, ".withJoin {")
		bb.Line("data.checkJoinFields()")
		bb.Line("}")
	}
	bb.Line("res = append(res, &data)")
	bb.Line("}")

	bb.Line("return res, err")
	bb.FuncEnd()
}