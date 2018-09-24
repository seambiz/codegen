package codegen

import "strings"

// TJSON template
func TJSON(bb *GenBuffer, conf *Config, schema *Schema, table *Table) {

	bb.Line("// ToJSON writes a single object to the buffer.")
	bb.Func(table.storeReceiver, "ToJSON")
	bb.FuncParams("t *buffer.TemplateBuffer", "data *"+table.title)
	bb.FuncReturn("")
	bb.Line(`prepend := "{"`)
	for _, f := range table.Fields {
		bb.Line("if ", table.initials, ".colSet == nil || ", table.initials, ".colSet.Test(", table.title+f.title, ") {")
		bb.Line("t.", f.jsonFunc, `(prepend, "`, strings.ToLower(f.Name), `", data.`, f.title, ")")
		bb.Line(`prepend = ","`)
		bb.Line("}")
	}
	bb.Line("t.S(`}`)")
	bb.FuncEnd()

	bb.Line("// ToJSONArray writes a slice to the named array.")
	bb.Func(table.storeReceiver, "ToJSONArray")
	bb.FuncParams("w io.Writer", "data []*"+table.title, "name string")
	bb.FuncReturn("")
	bb.Line(`t := buffer.NewTemplateBuffer()`)
	bb.Line("t.SS(`{\"`, name, `\":[`)")
	bb.S(`for i := range data {
		if i > 0 {
			t.S(",")
		}
		`, table.initials, `.ToJSON(t, data[i])
	}

	t.S("]}")
	_, err := w.Write(t.Bytes())
	if err != nil {
		panic(err)
	}`)
	bb.NewLine()

	bb.FuncEnd()
}