package codegen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"bitbucket.org/codegen/static"
	"github.com/danverbraganza/varcaser/varcaser"
	"github.com/valyala/bytebufferpool"
)

var (
	codegenStart = []byte("// GENERATED BY CODEGEN. DO NOT EDIT.")
	codegenEnd   = []byte("// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^")
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
func (g *GenBuffer) Func(Receiver, funcName string) {
	g.S("func ")
	if Receiver != "" {
		g.S("(")
		g.S(Receiver)
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

// LogField generates zerolog logging instruction for single field
func (g *GenBuffer) LogField(f *Field, prefix string) {
	switch f.GoType {
	case "[]byte":
		g.S("Bytes")
	case "time.Time":
		g.S("Time")
	case "sql.NullString":
		g.S("Str")
	case "sql.NullInt64", "int64":
		g.S("Int64")
	case "sql.NullFloat64":
		g.S("Float64")
	case "string":
		g.S("Str")
	case "uint64":
		g.S("Uint64")
	default:
		g.S(strings.Title(f.GoType))
	}
	g.S(`("`)
	g.S(f.Title)
	g.S(`", `)
	if f.IsNullable {
		switch f.GoType {
		case "time.Time":
			g.S("logTime(")
		case "string":
			g.S("logString(")
		case "bool":
			g.S("logBool(")
		case "int":
			g.S("logInt(")
		case "uint":
			g.S("logUInt(")
		case "float32":
			g.S("logFloat32(")
		case "float64":
			g.S("logFloat64(")
		case "uint64":
			g.S("logUInt64(")
		case "int64":
			g.S("logInt64(")
		case "[]byte":
			g.S("logBytes(")
		default:
			panic("unsupported pointer type: " + f.GoType)
		}
	}
	if prefix != "" {
		g.S(prefix)
		g.S(".")
		g.S(f.Title)
	} else {
		g.S(f.ParamName)
	}

	switch f.GoType {
	case "sql.NullString":
		g.S(".Str")
	case "sql.NullInt64":
		g.S(".Int64")
	case "sql.NullFloat64":
		g.S(".Float64")
	}
	g.S(")")
	if f.IsNullable {
		g.S(")")
	}
}

// Log generates zerolog logging instruction for array of fields
func (g *GenBuffer) Log(fields []*Field, prefix string) {
	for i, f := range fields {
		if i > 0 {
			g.S(".")
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
    str - the string to get Initials from
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

/*
 Parse the config and add all filtered templates that should be processed to each schema and table.
 This way the generation code will be much simpler.
 And this function will be easier to test.
*/
func generateTemplatesConfig(conf *Config) {
	// not the most performant approach, as templates will be rebuilt for each table, but much more readable than the old code
	for _, schema := range conf.Schemas {

		schema.preparedTemplatefiles = make(map[string][]string)
		for _, t := range schema.Tables {
			t.preparedTemplatefiles = make(map[string][]string)
		}

		var fileNames []string

		if schema.TemplateFolder != "" {
			files, err := ioutil.ReadDir(schema.TemplateFolder)
			if err != nil {
				panic(err)
			}

			for _, f := range files {
				fileNames = append(fileNames, f.Name())
			}
		} else {
			// use embedded templates
			var err error
			fileNames, err = static.WalkDirs("", false)
			if err != nil {
				panic(err)
			}
			sort.Strings(fileNames)
		}

		// first check all templates, that do not need to be concatenated
		for _, fName := range fileNames {
			segments := strings.Split(fName, ".")
			if len(segments) != 4 && len(segments) != 5 {
				panic("filename segment length != 4,5")
			}

			if len(segments) == 4 {
				switch segments[0] {
				case "once":
					// 1. add all simple "once" templates to the schema
					mapKey := strings.Join([]string{segments[0], segments[1], segments[2]}, ".")
					schema.preparedTemplatefiles[mapKey] = append(schema.preparedTemplatefiles[mapKey], fName)

				case "table":
					// 2. add all simple "table" templates to the tables
					for _, t := range schema.Tables {
						if !t.Generate {
							continue
						}
						mapKey := strings.Join([]string{segments[0], segments[1], segments[2]}, ".")
						t.preparedTemplatefiles[mapKey] = append(t.preparedTemplatefiles[mapKey], fName)
					}
				}
			}
		}

		// remove all templates, that were processed above
		for i := 0; i < len(fileNames); i++ {
			segments := strings.Split(fileNames[i], ".")
			if len(segments) != 5 {
				fileNames = append(fileNames[:i], fileNames[i+1:]...)
				i--
			}
		}

		// add all "once" files to be concatenated
		for _, fName := range fileNames {
			segments := strings.Split(fName, ".")
			if len(segments) != 5 {
				panic("wrong []segement length")
			}

			currentMapKey := strings.Join([]string{segments[0], segments[1], segments[2]}, ".")
			if segments[0] == "once" {
				schema.preparedTemplatefiles[currentMapKey] = append(schema.preparedTemplatefiles[currentMapKey], fName)
			}
		}

		// generate general template files list for table
		templateFiles := []string{}
		for _, fName := range fileNames {
			segments := strings.Split(fName, ".")
			if len(segments) != 5 {
				panic("wrong []segement length")
			}

			if segments[0] == "table" {
				templateFiles = append(templateFiles, fName)
			}
		}

		// add specific or global templates to each table
		for _, t := range schema.Tables {
			if !t.Generate {
				continue
			}

			// choose files to process
			tempFiles := templateFiles
			if len(t.TemplateFiles) > 0 {
				tempFiles = t.TemplateFiles
			}

			for _, tFile := range tempFiles {
				segments := strings.Split(tFile, ".")
				if len(segments) != 5 {
					panic("wrong []segement length")
				}

				currentMapKey := strings.Join([]string{segments[0], segments[1], segments[2]}, ".")
				t.preparedTemplatefiles[currentMapKey] = append(t.preparedTemplatefiles[currentMapKey], tFile)
			}
		}
	}
}

func generateFile(conf *Config, schema *Schema, table *Table, fileprefix string, templateFiles map[string][]string) {
	for key, fileNames := range templateFiles {
		var buf bytes.Buffer
		templateContents := []byte("")

		for _, tFile := range fileNames {
			var err error
			var contents []byte
			if schema.TemplateFolder != "" {
				contents, err = ioutil.ReadFile(schema.TemplateFolder + tFile)
				if err != nil {
					panic(err)
				}
			} else {
				contents, err = static.ReadFile(tFile)
				if err != nil {
					panic(err)
				}
			}
			templateContents = append(templateContents, contents...)
		}

		tmpl, err := template.New(key).Funcs(getFuncMap()).Parse(string(templateContents))
		if err != nil {
			panic(err)
		}

		type GenData struct {
			Conf   *Config
			Schema *Schema
			Table  *Table
		}

		d := GenData{}
		d.Conf = conf
		d.Schema = schema
		d.Table = table

		err = tmpl.Execute(&buf, d)
		if err != nil {
			panic(err)
		}

		segments := strings.Split(key, ".")
		prefix := ""
		if segments[0] == "table" {
			prefix = strings.ReplaceAll(fileprefix, "_", "")
		}
		switch segments[1] {
		case "subpackage":
			writeToCodgenFile(buf, conf, prefix+strings.Title(segments[2]), conf.SubPackage)
		case "package":
			writeToCodgenFile(buf, conf, prefix+strings.Title(segments[2]), filepath.Join(conf.SubPackage, conf.Package))
		case "root":
			writeToCodgenFile(buf, conf, prefix+strings.Title(segments[2]), "")
		}
	}
}

func generateCode(conf *Config) {
	for _, schema := range conf.Schemas {
		generateFile(conf, schema, nil, schema.Prefix, schema.preparedTemplatefiles)

		for _, table := range schema.Tables {
			generateFile(conf, schema, table, schema.Prefix+table.Name, table.preparedTemplatefiles)
		}
	}

	if conf.GoFmtCmd != "" {
		execCommand(conf.GoFmtCmd + " " + conf.DirOut)
	}
	if conf.LintPackage != "" {
		execCommand("go vet " + conf.LintPackage)
	}
	if conf.MetaLinter != "" {
		execCommand(conf.MetaLinter)
	}
}

// Generate main call to start generation
func Generate(conf *Config) {
	prepareSchemaConfig(conf)
	generateTemplatesConfig(conf)
	generateCode(conf)
}

func fnTableJoinFields(baseTable *Table, conf *Config, table *Table, refAlias rune, tableAlias *rune, fks []*ForeignKey) {
	for _, fk := range fks {
		*tableAlias++
		if fk.IsUnique {
			fkRefTableTitle := strings.Title(fk.RefTable)
			fkRefTable := strings.Title(fk.RefTable)
			fkSchema := conf.getSchema(fk.RefSchema)
			pRefTable := fkSchema.getTable(fk.RefTable)
			if pRefTable != nil {
				fkRefTableTitle = pRefTable.Title
				fkRefTable = pRefTable.Name
			}
			t := Join{
				Alias:    string(*tableAlias),
				Name:     fkRefTable,
				Title:    fkRefTableTitle,
				Initials: table.Initials,
				Schema:   fkSchema.Name,
				Table:    pRefTable,
			}
			for i := range fk.Fields {
				t.Fields = append(t.Fields, JoinField{
					Alias:    string(refAlias),
					Name:     fk.Fields[i],
					RefAlias: string(*tableAlias),
					RefName:  fk.RefFields[i],
				})
			}
			baseTable.Joins = append(baseTable.Joins, t)

			if len(fk.ForeignKeys) > 0 {
				fnTableJoinFields(baseTable, conf, table, *tableAlias, tableAlias, fk.ForeignKeys)
			}
		}
	}
}

func prepareSchemaConfig(conf *Config) {
	tableNr := 0
	for _, schema := range conf.Schemas {
		prefix := conf.Prefix
		if schema.Prefix != "" {
			prefix = schema.Prefix
		}

		// which templates should be used
		if schema.TemplateFolder == "" {
			schema.TemplateFolder = conf.TemplateFolder
		}

		tableNames := make([]string, len(schema.Tables))
		for i := range schema.Tables {
			tableNames[i] = schema.Tables[i].Name
		}
		tablesCase, err := varcaser.Detect(tableNames)
		if err != nil {
			tablesCase = varcaser.LowerSnakeCase
		}
		if schema.Title == "" {
			schema.Title = varcaser.Caser{From: tablesCase, To: varcaser.UpperCamelCase}.String(schema.Name)
		}

		for _, table := range schema.Tables {
			if !table.Generate {
				continue
			}
			table.id = tableNr
			tableNr++

			// generate helper variables
			// parts := strings.Split(table.Name, "_")
			// for i := range parts {
			// 	if strings.ToLower(parts[i]) == "id" {
			// 		parts[i] = "ID"
			// 	} else {
			// 		parts[i] = strings.Title(parts[i])
			// 	}
			// }
			table.Title = prefix + varcaser.Caser{From: tablesCase, To: varcaser.UpperCamelCase}.String(table.Name)
			table.lower = lowerFirst(table.Title)
			table.Initials = Initials(table.Name)
			table.Initials += Initials(table.Name[1:])
			table.Initials = "s"
			table.Receiver = table.Initials + " *" + table.Title
			table.store = table.Title + "Store"
			table.StoreReceiver = "s *" + table.store

			fieldNames := make([]string, len(table.Fields))
			for i := range table.Fields {
				fieldNames[i] = table.Fields[i].Name
			}
			fieldsCase, err := varcaser.Detect(fieldNames)
			if err != nil {
				fieldsCase = varcaser.LowerSnakeCase
			}

			// fill mapping for easy access to field properties
			table.FieldMapping = make(map[string]int)
			for i := range table.Fields {
				// if table.Fields[i].Name == "id" {
				// 	table.Fields[i].Title = "ID"
				// } else {
				// 	parts := strings.Split(table.Fields[i].Name, "_")
				// 	for i := range parts {
				// 		if strings.ToLower(parts[i]) == "id" {
				// 			parts[i] = "ID"
				// 		} else {
				// 			parts[i] = strings.Title(parts[i])
				// 		}
				// 	}
				// 	table.Fields[i].Title = strings.Join(parts, "")
				// }

				table.Fields[i].Title = varcaser.Caser{From: fieldsCase, To: varcaser.UpperCamelCase}.String(table.Fields[i].Name)
				table.Fields[i].Title = strings.ReplaceAll(table.Fields[i].Title, " ", "")

				// uppercase abbreviations
				for _, substring := range commonInitialisms {
					if strings.HasSuffix(strings.ToUpper(table.Fields[i].Title), substring) ||
						strings.HasPrefix(strings.ToUpper(table.Fields[i].Title), substring) {
						index := strings.Index(strings.ToUpper(table.Fields[i].Title), substring)
						stemp := table.Fields[i].Title[index : index+len(substring)]
						table.Fields[i].Title = strings.Replace(table.Fields[i].Title, stemp, substring, 1)
						break
					}
				}

				table.Fields[i].ParamName = strings.ToLower(table.Fields[i].Title)
				table.FieldMapping[table.Fields[i].Name] = i
				if table.Fields[i].IsPrimaryKey {
					table.PkFields = append(table.PkFields, table.Fields[i])
				} else {
					table.OtherFields = append(table.OtherFields, table.Fields[i])
				}
				typename, ok := GoTypeMapping[table.Fields[i].DBType]
				if !ok {
					panic(table.Fields[i].Name)
				}
				table.Fields[i].GoType = typename
				zero, ok := goZeroMapping[typename]
				if !ok {
					panic(typename)
				}
				table.Fields[i].GoZero = zero

				jsonFunc, ok := goJSONMapping[typename]
				if !ok {
					panic(typename)
				}
				table.Fields[i].JsonFunc = jsonFunc

				MappingFunc, ok := goDbMappingFunc[typename]
				if !ok {
					panic(typename)
				}
				table.Fields[i].MappingFunc = MappingFunc

				table.numFields = len(table.Fields)
			}
		}

		for _, table := range schema.Tables {
			if !table.Generate {
				continue
			}

			fieldNames := make([]string, len(table.Fields))
			for i := range table.Fields {
				fieldNames[i] = table.Fields[i].Name
			}
			fieldsCase, err := varcaser.Detect(fieldNames)
			if err != nil {
				fieldsCase = varcaser.LowerSnakeCase
			}

			for k, fk := range table.ForeignKeys {
				if fk.IsUnique {
					table.NumUniqueFKs++
				}
				fk.GenTableName = strings.Title(fk.RefTable)

				fkSchema := conf.getSchema(fk.RefSchema)
				if t := fkSchema.getTable(fk.RefTable); t != nil {
					fk.GenTableName = t.Title
					fk.GenTable = t
				}

				if fk.IsUnique {
					fk.GenFuncName = "OneBy"
				} else {
					fk.GenFuncName = "QueryBy"
				}
				for i := range fk.Fields {
					if i > 0 {
						fk.GenFuncName += "And"
					}
					fk.GenFuncName += fk.GenTable.Fields[fk.GenTable.FieldMapping[fk.RefFields[i]]].Title
				}

				// table.ForeignKeys[k].RefTable = varcaser.Caser{From: tablesCase, To: varcaser.UpperCamelCase}.String(fkRefTable)
				table.ForeignKeys[k].RefTableTitle = fk.GenTableName
				if fk.CustomName == "" {
					table.ForeignKeys[k].CustomName = varcaser.Caser{From: fieldsCase, To: varcaser.UpperCamelCase}.String(strings.Replace(fk.Name, "fk_", "", 1))
				}

				if !fk.IsUnique && len(fk.RefFields) > 1 {
					panic("FK: too many ref fields")
				}
			}

			for _, fk := range table.ForeignKeys {
				if fk.CustomName != "" {
					fk.GenName = fk.CustomName
				} else {
					fk.GenName = fk.GenTable.Title + strings.Replace(fk.Name, "fk", "", 1)
				}
			}

			table.Alias = "A"
			tableAlias := 'A'
			fnTableJoinFields(table, conf, table, tableAlias, &tableAlias, table.ForeignKeys)

			// Index func name
			for _, index := range table.Indices {
				funcName := ""
				if index.IsUnique {
					funcName = "OneBy"
				} else {
					funcName = "QueryBy"
				}
				for i, f := range index.Fields {
					if i > 0 {
						funcName += "And"
					}
					funcName += table.Fields[table.FieldMapping[f]].Title
				}
				index.FuncName = funcName
			}
		}
	}
}

func writeBufferToCodgenFile(bb *GenBuffer, conf *Config, filename string) {
	fileName := filepath.Join(conf.DirOut, fmt.Sprintf(conf.FilePattern, strings.ToLower(strings.Replace(filename, "_", "", -1))))
	fmt.Println("Writing to", fileName)

	// check if file exists and if it already has codegen comments
	// if not, just write everything to the file
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(fileName, bb.Bytes(), os.ModePerm)
		bb.Free()
		if err != nil {
			panic(err)
		}
	} else {
		fileContents, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		if bytes.Contains(fileContents, codegenStart) && bytes.Contains(fileContents, codegenEnd) {
			start := bytes.Index(fileContents, codegenStart)
			if start == -1 {
				panic("start == -1")
			}
			end := bytes.LastIndex(fileContents, codegenEnd)
			if end == -1 {
				panic("end == -1")
			}

			newStart := bytes.Index(bb.Bytes(), codegenStart)
			if newStart == -1 {
				panic("newStart == -1")
			}
			newEnd := bytes.LastIndex(bb.Bytes(), codegenEnd)
			if newEnd == -1 {
				panic("newEnd == -1")
			}

			var newContent []byte
			newContent = append(newContent, fileContents[:start]...)
			newContent = append(newContent, bb.Bytes()[newStart:newEnd]...)
			newContent = append(newContent, fileContents[end:]...)

			err := ioutil.WriteFile(fileName, newContent, os.ModePerm)
			bb.Free()
			if err != nil {
				panic(err)
			}

		} else {
			fmt.Println("ERR: existing file (" + fileName + ") does not contain codegen comment.")
			fmt.Println("ERR: exiting now, so content does not get overwritten.")
			os.Exit(1)
		}
	}
}

func writeToCodgenFile(buf bytes.Buffer, conf *Config, filename string, subfolder string) {
	fileName := filepath.Join(conf.DirOut, subfolder, fmt.Sprintf(conf.FilePattern, strings.ToLower(filename)))
	fmt.Println("Writing to", fileName)

	// check if file exists and if it already has codegen comments
	// if not, just write everything to the file
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println(err)
		err := ioutil.WriteFile(fileName, buf.Bytes(), os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		fileContents, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		if bytes.Contains(fileContents, codegenStart) && bytes.Contains(fileContents, codegenEnd) {
			start := bytes.Index(fileContents, codegenStart)
			if start == -1 {
				panic("start == -1")
			}
			end := bytes.LastIndex(fileContents, codegenEnd)
			if end == -1 {
				panic("end == -1")
			}

			newStart := bytes.Index(buf.Bytes(), codegenStart)
			if newStart == -1 {
				panic("newStart == -1")
			}
			newEnd := bytes.LastIndex(buf.Bytes(), codegenEnd)
			if newEnd == -1 {
				panic("newEnd == -1")
			}

			var newContent []byte
			newContent = append(newContent, fileContents[:start]...)
			newContent = append(newContent, buf.Bytes()[newStart:newEnd]...)
			newContent = append(newContent, fileContents[end:]...)

			err := ioutil.WriteFile(fileName, newContent, os.ModePerm)
			if err != nil {
				panic(err)
			}

		} else {
			fmt.Println("ERR: existing file (" + fileName + ") does not contain codegen comment.")
			fmt.Println("ERR: exiting now, so content does not get overwritten.")
			os.Exit(1)
		}
	}
}
