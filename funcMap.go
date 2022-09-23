package codegen

import (
	"text/template"

	"bitbucket.org/codegen/config"
	"github.com/Masterminds/sprig"
	"github.com/valyala/bytebufferpool"
)

func getFuncMap() template.FuncMap {
	fmap := sprig.TxtFuncMap()
	fmap["zerologFields"] = genZerologFields

	return fmap
}

func genZerologFields(fields []*config.Field, prefix string) string {
	bb := NewGenBuffer(bytebufferpool.Get())
	bb.Log(fields, prefix)
	s := string(bb.Bytes())
	bb.Free()
	return s
}
