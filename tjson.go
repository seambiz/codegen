package codegen

import "strings"

// TJSON template
func TJSON(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {

	bb.Line("// ToJSON writes a single object to the buffer.")
	bb.Line("// nolint[gocylco]")
	bb.Func(table.StoreReceiver, "ToJSON")
	bb.FuncParams("t *buffer.TemplateBuffer", "data *"+table.Title)
	bb.FuncReturn("")
	bb.Line(`prepend := "{"`)
	lenFields := len(table.Fields) - 1
	for i, f := range table.Fields {
		bb.Line("if ", table.Initials, ".colSet == nil || ", table.Initials, ".colSet.Bit(", table.Title+f.Title, ") == 1 {")
		bb.Line("t.", f.jsonFunc, `(prepend, "`, strings.ToLower(f.Name), `", data.`, f.Title, ")")
		if i != lenFields {
			bb.Line(`prepend = ","`)
		}
		bb.Line("}")
	}
	bb.Line("t.S(`}`)")
	bb.FuncEnd()

	bb.Line("// ToJSONArray writes a slice to the named array.")
	bb.Func(table.StoreReceiver, "ToJSONArray")
	bb.FuncParams("w io.Writer", "data []*"+table.Title, "name string")
	bb.FuncReturn("")
	bb.Line(`t := buffer.NewTemplateBuffer()`)
	bb.Line("t.SS(`{\"`, name, `\":[`)")
	bb.S(`for i := range data {
		if i > 0 {
			t.S(",")
		}
		`, table.Initials, `.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}`)
	bb.NewLine()

	bb.FuncEnd()
}
