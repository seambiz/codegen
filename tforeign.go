package codegen

import (
	"strings"
)

// TODO: non IsUnique
// TODO: automatisch JOIN, um alle abhängigkeiten mitzuladen. sehr spannend
// TODO: ggf. nur für 1:1 IsUnique

// TForeign template
func TForeign(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {
	for _, fk := range table.ForeignKeys {
		fkRefTable := strings.Title(fk.RefTable)
		fkTable := table
		fkSchema := conf.getSchema(fk.RefSchema)
		if t := fkSchema.getTable(fk.RefTable); t != nil {
			fkRefTable = t.title
			fkTable = t
		}
		if fk.CustomName == "" {
			fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
		}

		bb.Line("// Get", fk.CustomName, " fetches a record from referenced table '", fk.RefTable, "'.")
		bb.Func(table.receiver, "Get"+fk.CustomName)
		bb.FuncParams("db *sql.DB")
		bb.FuncReturn("error")
		bb.Line("if ", table.initials, ".", fk.CustomName, " == nil {")
		bb.Line("var err error")
		bb.S(table.initials, ".", fk.CustomName, ",err = New", fkRefTable, "Store(db).")
		var funcName string
		if fk.IsUnique {
			funcName = "OneBy"
		} else {
			funcName = "QueryBy"
		}
		for i := range fk.Fields {
			if i > 0 {
				funcName += "And"
			}
			funcName += fkTable.Fields[fkTable.fieldMapping[fk.RefFields[i]]].title
		}
		bb.S(funcName, "(")
		for i := range fk.Fields {
			if i > 0 {
				bb.S(",")
			}
			bb.S(table.initials, ".", table.Fields[table.fieldMapping[fk.Fields[i]]].title)
		}
		bb.Line(")")

		bb.Line("return err")
		bb.Line("}")
		bb.Line("return nil")
		bb.FuncEnd()
	}

	for _, fk := range table.ForeignKeys {
		if !fk.IsUnique {
			fkRefTable := strings.Title(fk.RefTable)
			fkTable := table
			fkSchema := conf.getSchema(fk.RefSchema)
			if t := fkSchema.getTable(fk.RefTable); t != nil {
				fkRefTable = t.title
				fkTable = t
			}
			if fk.CustomName == "" {
				fk.CustomName = table.title + strings.Replace(fk.Name, "fk", "", 1)
			}

			bb.Line("// EagerFetch", fk.CustomName, " eagerly fetches N records from referenced table '", fk.RefTable, "'.")
			bb.Func(table.storeReceiver, "EagerFetch"+fk.CustomName)
			bb.FuncParams("data []*" + table.title)
			bb.FuncReturn("error")

			if len(fk.RefFields) > 1 {
				panic("too many ref fields")
			}

			bb.Line(`stmt := NewSQLStatement()`)
			bb.Line(`stmt.Append("`, fk.RefFields[0], ` IN (")`)

			bb.Line(`for i, d := range data {`)
			bb.Line(`if i > 0 {`)
			bb.Line(`stmt.Append(",")`)
			bb.Line(`}`)
			bb.Line(`stmt.AppendInt(d.`, table.Fields[table.fieldMapping[fk.Fields[0]]].title, `)`)
			bb.Line(`}`)
			bb.Line(`stmt.Append(")")`)

			bb.Line(`details, err := New`, fkRefTable, `Store(`, table.initials, `.db).Where(stmt.Query()).OrderBy("A.`, fk.RefFields[0], " DESC, A.", fk.Fields[0], ` DESC").Query()`)
			bb.Line(`if err != nil {`)
			bb.Line(`log.Error().Err(err).Msg("fetch details")`)
			bb.Line(`return err`)
			bb.Line(`}`)

			bb.Line(`for i := range data {`)
			bb.Line(`for j := len(details) - 1; j >= 0; j-- {`)
			bb.Line(`if details[j].`, fkTable.Fields[fkTable.fieldMapping[fk.RefFields[0]]].title, ` == data[i].`, table.Fields[table.fieldMapping[fk.Fields[0]]].title, ` {`)
			bb.Line(`data[i].`, fk.CustomName, ` = append(data[i].`, fk.CustomName, `, details[j])`)
			bb.Line(`details = append(details[:j], details[j+1:]...)`)
			bb.Line(`}`)
			bb.Line(`}`)

			bb.Line("}")
			bb.Line("return nil")
			bb.FuncEnd()
		}
	}
}
