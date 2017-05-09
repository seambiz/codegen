package codegen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/valyala/bytebufferpool"
)

// GenBuffer type alias for shorter template functions
type GenBuffer struct {
	buf *bytebufferpool.ByteBuffer
}

// NewGenBuffer factory
func NewGenBuffer(bb *bytebufferpool.ByteBuffer) *GenBuffer {
	return &GenBuffer{buf: bb}
}

// NewLine write newline char to buffer
func (g *GenBuffer) NewLine() {
	_, err := g.buf.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}
}

// S shortcut for writing string to buffer and check error
func (g *GenBuffer) S(ss ...string) {
	for _, s := range ss {
		_, err := g.buf.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
}

// Func shortcut for writing string to buffer and check error
func (g *GenBuffer) Func(receiver, funcName string) {
	g.S("func ")
	if receiver != "" {
		g.S("(")
		g.S(receiver)
		g.S(")")
	}
	g.S(" ")
	g.S(funcName)
}

// FuncParams shortcut for writing string to buffer and check error
func (g *GenBuffer) FuncParams(ss ...string) {
	g.S("(")
	for i, s := range ss {
		if i >= 1 {
			g.S(",")
		}
		g.S(s)
	}
	g.S(")")
}

// FuncEnd shortcut for writing string to buffer and check error
func (g *GenBuffer) FuncEnd() {
	g.S("}")
	g.NewLine()
}

// FuncReturn shortcut for writing string to buffer and check error
func (g *GenBuffer) FuncReturn(ss ...string) {
	if len(ss) == 0 {
		g.S("{")
		return
	}
	if len(ss) > 1 {
		g.S("(")
	}
	for i, s := range ss {
		if i >= 1 {
			g.S(",")
		}
		g.S(s)
	}
	if len(ss) > 1 {
		g.S(")")
	}
	g.S("{")
	g.NewLine()
}

// Struct shortcut for writing string to buffer and check error
func (g *GenBuffer) Struct(name string) {
	g.Line("type ", name, " struct {")
}

// StructEnd shortcut for writing string to buffer and check error
func (g *GenBuffer) StructEnd() {
	g.Line("}")
}

// Line shortcut for writing string to buffer and check error
func (g *GenBuffer) Line(ss ...string) {
	for _, s := range ss {
		g.S(s)
	}
	g.NewLine()
}

// LogField generates zap logging instruction for single field
func (g *GenBuffer) LogField(f *Field, prefix string) {
	g.S("zap.")
	switch f.goType {
	case "[]byte":
		g.S("ByteString")
		break
	case "time.Time":
		g.S("Time")
		break
	case "sql.NullString":
		g.S("String")
		break
	case "sql.NullInt64":
		g.S("Int64")
		break
	case "sql.NullFloat64":
		g.S("Float64")
		break
	default:
		g.S(strings.Title(f.goType))
	}
	g.S(`("`)
	g.S(f.title)
	g.S(`", `)
	if prefix != "" {
		g.S(prefix)
		g.S(".")
		g.S(f.title)
	} else {
		g.S(f.paramName)
	}
	switch f.goType {
	case "sql.NullString":
		g.S(".String")
		break
	case "sql.NullInt64":
		g.S(".Int64")
		break
	case "sql.NullFloat64":
		g.S(".Float64")
		break
	}
	g.S(")")
}

// Log generates zap logging instruction for array of fields
func (g *GenBuffer) Log(fields []*Field, prefix string) {

	for i, f := range fields {
		if i > 0 {
			g.S(", ")
		}
		g.LogField(f, prefix)
	}
}

// Bytes returns buffer contents
func (g *GenBuffer) Bytes() []byte {
	return g.buf.Bytes()
}

// Free returns buffer to pool
func (g *GenBuffer) Free() {
	bytebufferpool.Put(g.buf)
}

/*
Initials extracts the initial letters from each word in the string. The first letter of the string and all first
letters after the defined delimiters are returned as a new string. Their case is not changed. If the delimiters
parameter is excluded, then Whitespace is used. Whitespace is defined by unicode.IsSpacea(char). An empty delimiter array returns an empty string.
Parameters:
    str - the string to get initials from
    delimiters - set of characters to determine words, exclusion of this parameter means whitespace would be delimeter
Returns:
    string of initial letters
*/
func Initials(str string, delimiters ...rune) string {
	if str == "" {
		return str
	}
	if delimiters != nil && len(delimiters) == 0 {
		return ""
	}
	strLen := len(str)
	var buf bytes.Buffer
	lastWasGap := true
	for i := 0; i < strLen; i++ {
		ch := rune(str[i])

		if isDelimiter(ch, delimiters...) {
			lastWasGap = true
		} else if lastWasGap {
			buf.WriteRune(ch)
			lastWasGap = false
		}
	}
	return buf.String()
}

// private function (lower case func name)
func isDelimiter(ch rune, delimiters ...rune) bool {
	if delimiters == nil {
		return unicode.IsSpace(ch)
	}
	for _, delimiter := range delimiters {
		if ch == delimiter {
			return true
		}
	}
	return false
}

// execCommand wraps exec.Command
func execCommand(command string) {
	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		return
	}

	fmt.Println(parts)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

var commonInitialisms = []string{
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF8",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
}

// Generate main call to start generation
func Generate(conf *Config) {
	for _, schema := range conf.Schemas {
		for _, table := range schema.Tables {

			// generate helper variables
			parts := strings.Split(table.Name, "_")
			for i := range parts {
				if strings.ToLower(parts[i]) == "id" {
					parts[i] = "ID"
				} else {
					parts[i] = strings.Title(parts[i])
				}
			}
			table.title = strings.Join(parts, "")
			table.lower = lowerFirst(table.title)
			table.initials = Initials(table.Name)
			table.receiver = table.initials + " *" + table.title
			table.store = table.title + "Store"
			table.storeReceiver = table.initials + " *" + table.store

			// fill mapping for easy access to field properties
			table.fieldMapping = make(map[string]int)
			for i := range table.Fields {
				if table.Fields[i].Name == "id" {
					table.Fields[i].title = "ID"
				} else {
					parts := strings.Split(table.Fields[i].Name, "_")
					for i := range parts {
						if strings.ToLower(parts[i]) == "id" {
							parts[i] = "ID"
						} else {
							parts[i] = strings.Title(parts[i])
						}
					}
					table.Fields[i].title = strings.Join(parts, "")
				}
				// uppercase abbreviations
				for _, substring := range commonInitialisms {
					if strings.HasSuffix(strings.ToUpper(table.Fields[i].title), substring) ||
						strings.HasPrefix(strings.ToUpper(table.Fields[i].title), substring) {
						index := strings.Index(strings.ToUpper(table.Fields[i].title), substring)
						stemp := table.Fields[i].title[index : index+len(substring)]
						table.Fields[i].title = strings.Replace(table.Fields[i].title, stemp, substring, 1)
						break
					}
				}

				table.Fields[i].paramName = strings.ToLower(table.Fields[i].title)
				table.fieldMapping[table.Fields[i].Name] = i
				if table.Fields[i].IsPrimaryKey {
					table.pkFields = append(table.pkFields, table.Fields[i])
				} else {
					table.otherFields = append(table.otherFields, table.Fields[i])
				}
				typename, ok := goTypeMapping[table.Fields[i].DBType]
				if !ok {
					panic(table.Fields[i].Name)
				}
				/*
					if table.Fields[i].IsNullable {
						//fmt.Println("NullType", table.Name, table.Fields[i].title)
						nulltype, ok := goIsNullMapping[typename]
						if !ok {
							panic(table.Fields[i].Name)
						}
						typename = nulltype
					}
				*/
				table.Fields[i].goType = typename
				zero, ok := goZeroMapping[typename]
				if !ok {
					panic(typename)
				}
				table.Fields[i].goZero = zero

				mappingFunc, ok := goDbMappingFunc[typename]
				if !ok {
					panic(typename)
				}
				table.Fields[i].mappingFunc = mappingFunc
			}
		}
	}

	for _, schema := range conf.Schemas {
		for _, table := range schema.Tables {
			bb := NewGenBuffer(bytebufferpool.Get())

			var templateFiles = table.Templates
			if len(templateFiles) == 0 {
				templateFiles = conf.Templates
			}

			for _, templateFile := range templateFiles {
				switch templateFile {
				case "header":
					THeader(bb, conf, schema, table)
				case "delete":
					TDelete(bb, conf, schema, table)
				case "insert":
					TInsert(bb, conf, schema, table)
				case "truncate":
					TTruncate(bb, conf, schema, table)
				case "update":
					TUpdate(bb, conf, schema, table)
				case "upsert":
					TUpsert(bb, conf, schema, table)
				case "type":
					TType(bb, conf, schema, table)
				case "foreign":
					TForeign(bb, conf, schema, table)
				case "index":
					type set map[string]struct{}
					indexSet := set{}

					for i := range table.Indices {
						if len(table.Indices[i].Fields) > 1 {
							if _, ok := indexSet[table.Indices[i].Fields[0]]; ok {
								continue
							}
							index := *table.Indices[i]
							index.IsUnique = false
							index.Fields = []string{table.Indices[i].Fields[0]}
							table.Indices = append(table.Indices, &index)
						}
						indexSet[table.Indices[i].Fields[0]] = struct{}{}

					}

					for _, index := range table.Indices {
						TIndex(bb, conf, schema, table, index)
					}
				}
			}

			fileName := filepath.Join(conf.DirOut, fmt.Sprintf(conf.FilePattern, strings.ToLower(strings.Replace(table.Name, "_", "", -1))))
			fmt.Println("Writing to", fileName)

			err := ioutil.WriteFile(fileName, bb.Bytes(), os.ModePerm)
			bb.Free()
			if err != nil {
				panic(err)
			}
		}
	}

	execCommand("goimports -w " + conf.DirOut)
	if conf.LintPackage != "" {
		execCommand("go vet -v " + conf.LintPackage)
	}
	if conf.MetaLinter != "" {
		execCommand(conf.MetaLinter)
	}
}
