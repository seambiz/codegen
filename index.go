package codegen

import (
	"strings"
)

func indexUnique(bb *GenBuffer, conf *Config, schema *Schema, table *Table, index *Index) {
	funcName := "OneBy"
	for i, f := range index.Fields {
		if i > 0 {
			funcName += "And"
		}
		funcName += table.Fields[table.FieldMapping[f]].Title
	}
	bb.Line("// ", funcName, " retrieves a row from '", schema.Name, ".", table.Name, "' as a ", table.Title, ".")
	bb.Line("//")
	bb.Line("// Generated from index '", index.Name, "'.")
	bb.Line("// nolint[goconst]")
	bb.Func(table.storeReceiver, funcName)
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
	bb.FuncReturn("*"+table.Title, "error")
	bb.S(table.initials, `.where = "`)
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].Name)
		bb.S(" = ?")
	}
	bb.Line(`"`)
	bb.S("return ", table.initials, ".One(")
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
	}
	bb.Line(")")

	bb.FuncEnd()
}

func indexSlice(bb *GenBuffer, conf *Config, schema *Schema, table *Table, index *Index) {
	funcName := "QueryBy"
	for i, f := range index.Fields {
		if i > 0 {
			funcName += "And"
		}
		funcName += table.Fields[table.FieldMapping[f]].Title
	}
	bb.Line("// ", funcName, " retrieves multiple rows from '", schema.Name, ".", table.Name, "' as a slice of ", table.Title, `.`)
	bb.Line("//")
	bb.Line("// Generated from index '", index.Name, "'.")
	bb.Func(table.storeReceiver, funcName)
	bb.S("(")
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
		bb.S(" ")
		bb.S(table.Fields[table.FieldMapping[f]].GoType)
	}
	bb.S(`)`)
	bb.FuncReturn("[]*"+table.Title, "error")
	bb.Line("var err error")
	bb.Line("res := []*", table.Title, "{}")
	bb.NewLine()

	bb.Line("sql := NewSQLStatement()")
	bb.Line(`sql.Append("SELECT")`)
	bb.Line(`sql.Fields("","", `, strings.ToLower(table.Title), `QueryFields)`)
	bb.Line(`sql.Append("FROM `, schema.Name, ".", table.Name, `")`)
	bb.S(`sql.Append("WHERE `)
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].Name)
		bb.S(" = ?")
	}
	bb.S(`")`)
	bb.NewLine()
	bb.Line(`if  zerolog.GlobalLevel() == zerolog.DebugLevel {`)
	bb.S(`log.Debug().Str("fn, "`)
	bb.S(funcName)
	bb.S(`").Str("stmt", sql.String()).Msg("sql")`)

	for i, f := range index.Fields {
		if i > 0 {
			bb.S(".")
		}
		bb.LogField(table.Fields[table.FieldMapping[f]], "")
	}
	bb.S(`)
	}
	`)
	bb.Line("q, err := ", table.initials, `.db.Query(sql.Query(), `)
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
	}
	bb.S(`)
	if err != nil {
		log.Error().Err(err).Msg("query")
		return nil, err
	}
	
	for q.Next() {`)
	bb.Line("data := ", table.Title, "{}")

	bb.S(`err = q.Scan(`)
	bb.Line("data.scanFields(", table.initials, ".withJoin)...)")
	bb.S(`if err != nil {
			log.Error().Err(err).Msg("scanFields")
			return nil, err
		}

		res = append(res, &data)
	}
	err = q.Close()
	return res, err
}
`)
}

// TIndex template
func TIndex(bb *GenBuffer, conf *Config, schema *Schema, table *Table, index *Index) {
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
	bb.Func(table.storeReceiver, funcName)
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
	bb.FuncReturn(arrayType+"*"+table.Title, "error")
	bb.S(table.initials, `.where = "`)
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(" AND ")
		}
		bb.S("A.", table.Fields[table.FieldMapping[f]].Name)
		bb.S(" = ?")
	}
	bb.Line(`"`)
	if index.IsUnique {
		bb.S("return ", table.initials, ".One(")
	} else {
		bb.S("return ", table.initials, ".Query(")
	}
	for i, f := range index.Fields {
		if i > 0 {
			bb.S(", ")
		}
		bb.S(table.Fields[table.FieldMapping[f]].ParamName)
	}
	bb.Line(")")

	bb.FuncEnd()
}
