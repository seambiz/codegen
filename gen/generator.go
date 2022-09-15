package gen

import (
	"strings"
)

var newLine = byte('\n')

type Generator struct {
	buf strings.Builder
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Go(ss ...string) {
	for i := range ss {
		g.buf.WriteString(ss[i])
	}
	g.buf.WriteByte(newLine)
}

func (g *Generator) Package(s string) {
	g.buf.WriteString("package ")
	g.buf.WriteString(s)
	g.buf.WriteByte(newLine)
}

func (g *Generator) Comment(ss ...string) {
	g.buf.WriteString("// ")
	for i := range ss {
		g.buf.WriteString(ss[i])
	}
	g.buf.WriteByte(newLine)
}

func (g *Generator) Line() {
	g.buf.WriteByte(newLine)
}

func (g *Generator) String() string {
	return g.buf.String()
}
