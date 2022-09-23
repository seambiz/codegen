package gen

import (
	"strings"

	"bitbucket.org/codegen/config"
	"github.com/elliotchance/orderedmap/v2"
)

var newLine = byte('\n')

type Codegen interface {
	Generate(conf *config.Config) string
}

type Generator struct {
	buf strings.Builder
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Line(ss ...string) {
	g.Lit(ss...)
	g.NewLine()
}

func (g *Generator) Go(ss ...string) {
	g.Lit(ss...)
	g.NewLine()
}

func (g *Generator) Id(ss ...string) string {
	s := ""
	if len(ss) == 1 {
		s = ss[0]
	} else if len(ss) == 2 {
		s = ss[0] + "." + ss[1]
	} else {
		panic("RetSlice must receive 1 or 2 parameters")
	}

	return s
}

func Sprintf(fmt string, ss ...string) string {
	g := NewGenerator()

	g.Lit("fmt.Sprintf(", `"`, fmt, `",`)
	g.Lit(ss...)
	g.Lit(")")

	return g.String()
}

func Ptr(ss ...string) string {
	g := NewGenerator()
	g.Lit("*")
	g.Lit(ss...)

	return g.String()
}

func PtrIf(b bool, ss ...string) string {
	g := NewGenerator()
	if b {
		g.Lit("*")
	}
	g.Lit(ss...)

	return g.String()
}

func (g *Generator) S(ss ...string) string {
	return strings.Join(ss, "")
}

func (g *Generator) Append(name string, value string) {
	g.Lit(name, " = append(", name, ", ", value, ")")
	g.NewLine()
}

func (g *Generator) Return(s string) {
	g.Lit("return ", s)
	g.NewLine()
}

func (g *Generator) Panic(s string) {
	g.Lit("panic (", s, ")")
	g.NewLine()
}

func NewStruct(name string, kv *KV) string {
	g := NewGenerator()

	g.Lit(name, "{")
	g.NewLine()

	for key, value := range *kv {
		g.Lit(key, ":", value, ",")
		g.NewLine()
	}
	g.Lit(name, "}")
	g.NewLine()

	return g.String()
}

func (g *Generator) Import(ss ...string) {
	g.Lit("import (")
	g.NewLine()

	for i := range ss {
		g.Lit(`"`, ss[i], `"`)
		g.NewLine()
	}
	g.Lit(")")
	g.NewLine()
	g.NewLine()
}

func (g *Generator) Package(s string) {
	g.Lit("package ", s)
	g.NewLine()
	g.NewLine()
}

func (g *Generator) Comment(ss ...string) {
	g.Lit("// ")
	g.Lit(ss...)
	g.NewLine()
}

func (g *Generator) ConstFn(fn func()) {
	g.Lit("const (")
	g.NewLine()

	fn()

	g.Lit(")")
	g.NewLine()
}

type KV map[string]string

func (g *Generator) Struct(name string, fn func()) {
	g.Lit("type ", name, " struct {")
	g.NewLine()

	fn()

	g.Lit("}")
	g.NewLine()
}

type Field struct {
	name        string
	isPtr       bool
	isSlice     bool
	typePointer bool
	typeName    string
	tags        *orderedmap.OrderedMap[string, string]

	g *Generator
}

type Method struct {
	name     string
	receiver string
	params   *orderedmap.OrderedMap[string, string]
	returns  *orderedmap.OrderedMap[string, string]

	g *Generator
}

type Func struct {
	name    string
	params  *orderedmap.OrderedMap[string, string]
	returns *orderedmap.OrderedMap[string, string]

	g *Generator
}

type IfStmt struct {
	g *Generator
}

func (g *Generator) If(ss ...string) *IfStmt {
	g.Lit("if (")
	g.Lit(ss...)
	g.Lit(") {")
	g.NewLine()

	ifstmt := &IfStmt{}
	ifstmt.g = g
	return ifstmt
}

func (i *IfStmt) Body(fn func()) {
	fn()

	i.g.Lit("}")
}

func (g *Generator) Method(ss ...string) *Method {
	m := &Method{}
	m.name = strings.Join(ss, "")
	m.g = g

	m.params = orderedmap.NewOrderedMap[string, string]()
	m.returns = orderedmap.NewOrderedMap[string, string]()

	return m
}

func (m *Method) Receiver(ss ...string) *Method {
	for i, s := range ss {
		if i > 0 {
			m.receiver += " "
		}
		m.receiver += s
	}

	return m
}

func (g *Generator) Func(ss ...string) *Func {
	f := &Func{}
	f.name = strings.Join(ss, "")
	f.g = g

	f.params = orderedmap.NewOrderedMap[string, string]()
	f.returns = orderedmap.NewOrderedMap[string, string]()

	return f
}

func (f *Func) Param(name string, sType string) *Func {
	f.params.Set(name, sType)

	return f
}

func (f *Func) Returns(ss ...string) *Func {
	if len(ss) == 1 {
		f.returns.Set(ss[0], "")
	} else if len(ss) == 2 {
		f.returns.Set(ss[0], ss[1])
	} else {
		panic("Returns must receive 1 or 2 parameters")
	}

	return f
}

func (f *Func) Body(fn func()) {
	f.g.Lit("func ")

	f.g.Lit(f.name, "(")
	i := 0
	for el := f.params.Front(); el != nil; el = el.Next() {
		if i > 0 {
			f.g.Lit(",")
		}
		f.g.Lit(el.Key, " ", el.Value)
	}
	f.g.Lit(") ")

	lenReturns := f.returns.Len()
	if lenReturns > 0 {
		if lenReturns > 1 {
			f.g.Lit("(")
		}
		i = 0
		for el := f.returns.Front(); el != nil; el = el.Next() {
			if i > 0 {
				f.g.Lit(",")
			}
			f.g.Lit(el.Key, " ", el.Value)
		}
		if lenReturns > 1 {
			f.g.Lit(")")
		}
	}

	f.g.Lit("{")
	f.g.NewLine()

	fn()

	f.g.Lit("}")
	f.g.NewLine()
	f.g.NewLine()
}

func (m *Method) Param(name string, sType string) *Method {
	m.params.Set(name, sType)

	return m
}

// func (m *Method) RetSlice(ss ...string) *Method {
// 	if len(ss) == 1 {
// 		m.returns["[]"+ss[0]] = ""
// 	} else if len(ss) == 2 {
// 		m.returns["[]"+ss[0]] = ss[1]
// 	} else {
// 		panic("RetSlice must receive 1 or 2 parameters")
// 	}

// 	return m
// }

func (m *Method) Returns(ss ...string) *Method {
	if len(ss) == 1 {
		m.returns.Set(ss[0], "")
	} else if len(ss) == 2 {
		m.returns.Set(ss[0], ss[1])
	} else {
		panic("Returns must receive 1 or 2 parameters")
	}

	return m
}

func (m *Method) Body(fn func()) {
	m.g.Lit("func ")

	if m.receiver != "" {
		m.g.Lit("(", m.receiver, ") ")
	}

	m.g.Lit(m.name, "(")
	i := 0
	for el := m.params.Front(); el != nil; el = el.Next() {
		if i > 0 {
			m.g.Lit(",")
		}
		m.g.Lit(el.Key, " ", el.Value)
	}
	m.g.Lit(") ")

	lenReturns := m.returns.Len()
	if lenReturns > 0 {
		if lenReturns > 1 {
			m.g.Lit("(")
		}
		i = 0
		for el := m.returns.Front(); el != nil; el = el.Next() {
			if i > 0 {
				m.g.Lit(",")
			}
			m.g.Lit(el.Key, " ", el.Value)
		}
		if lenReturns > 1 {
			m.g.Lit(")")
		}
	}

	m.g.Lit("{")
	m.g.NewLine()

	fn()

	m.g.Lit("}")
	m.g.NewLine()
	m.g.NewLine()
}

func (f *Field) Go() {
	f.g.Lit(f.name)

	if f.isPtr {
		f.g.Lit(" *")
	} else {
		f.g.Lit(" ")
	}

	if f.isSlice {
		f.g.Lit("[]")
	}

	if f.typePointer {
		f.g.Lit("*")
	}

	f.g.Lit(f.typeName)

	if f.tags.Len() > 0 {
		f.g.Lit(" `")

		i := 0
		for el := f.tags.Front(); el != nil; el = el.Next() {

			if i > 0 {
				f.g.Lit(" ")
			}
			f.g.Lit(el.Key, `:"`, el.Value, `"`)
			i++
		}

		f.g.Lit("`")
	}
	f.g.NewLine()
}

func (f *Field) Tag(key string, value string) *Field {
	f.tags.Set(key, value)

	return f
}

func (f *Field) Ptr() *Field {
	f.isPtr = true

	return f
}

func (f *Field) Slice(b bool) *Field {
	f.isSlice = b

	return f
}

func (f *Field) Type(name string) *Field {
	f.typeName = name

	return f
}

func (f *Field) PointerType(name string) *Field {
	f.typePointer = true
	f.typeName = name

	return f
}

func (g *Generator) Field(name string) *Field {
	f := &Field{}
	f.name = name
	f.g = g
	f.tags = orderedmap.NewOrderedMap[string, string]()

	return f
}

func (g *Generator) Lit(ss ...string) {
	for i := range ss {
		g.buf.WriteString(ss[i])
	}
}

func (g *Generator) NewLine() {
	g.buf.WriteByte(newLine)
}

func (g *Generator) String() string {
	return g.buf.String()
}
