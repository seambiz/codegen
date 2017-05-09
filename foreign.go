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
		bb.FuncParams("db *sqlx.DB")
		bb.FuncReturn("error")
		bb.Line("if ", table.initials, ".", fk.CustomName, " == nil {")
		bb.Line("var err error")
		bb.S(table.initials, ".", fk.CustomName, ",err = New", fkRefTable, "Store(db).")
		if fk.IsUnique {
			bb.S("One")
		} else {
			bb.S("Query")
		}
		bb.Line("By", fkTable.Fields[fkTable.fieldMapping[fk.RefFields[0]]].title, "(", table.initials, ".", table.Fields[table.fieldMapping[fk.Fields[0]]].title, ")")
		bb.Line("return err")
		bb.Line("}")
		bb.Line("return nil")
		bb.FuncEnd()
	}
}
