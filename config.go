package codegen

// Config base template input
type Config struct {
	MetaLinter     string   `json:",omitempty"`
	LintPackage    string   `json:",omitempty"`
	Package        string   `json:",omitempty"`
	RootPackage    string   `json:",omitempty"`
	DirOut         string   `json:",omitempty"`
	FilePattern    string   `json:",omitempty"`
	Templates      []string `json:",omitempty"`
	TemplateFolder string   `json:",omitempty"`

	Schemas  []*Schema `json:",omitempty"`
	Database db
}

type db struct {
	DSN     string
	Schemas []string
}

// Schema single schema type
type Schema struct {
	Name   string   `json:",omitempty"`
	Tables []*Table `json:",omitempty"`
}

func (c *Config) getSchema(schema string) *Schema {
	for i := range c.Schemas {
		if c.Schemas[i].Name == schema {
			return c.Schemas[i]
		}
	}
	return nil
}
func (s *Schema) getTable(table string) *Table {
	for i := range s.Tables {
		if s.Tables[i].Name == table {
			return s.Tables[i]
		}
	}
	return nil
}

// Table type
type Table struct {
	Name      string   `json:",omitempty"`
	Templates []string `json:",omitempty"`

	Fields      []*Field      `json:",omitempty"`
	Indices     []*Index      `json:",omitempty"`
	Ignores     IgnoreFields  `json:",omitempty"`
	ForeignKeys []*ForeignKey `json:",omitempty"`
	Generate    bool

	// generated Contents
	Title         string         `json:"-"`
	lower         string         `json:"-"`
	receiver      string         `json:"-"`
	initials      string         `json:"-"`
	store         string         `json:"-"`
	storeReceiver string         `json:"-"`
	FieldMapping  map[string]int `json:"-"`
	pkFields      []*Field       `json:"-"`
	otherFields   []*Field       `json:"-"`
	id            int            `json:"-"`
	numFields     int            `json:"-"`
	numUniqueFKs  int            `json:"-"`
}

// ForeignKey type
type ForeignKey struct {
	Name       string   `json:",omitempty"`
	Fields     []string `json:",omitempty"`
	RefSchema  string   `json:",omitempty"`
	RefTable   string   `json:",omitempty"`
	RefFields  []string `json:",omitempty"`
	IsUnique   bool     `json:",omitempty"`
	CustomName string   `json:",omitempty"`

	ForeignKeys []*ForeignKey `json:",omitempty"`
}

// Field type
type Field struct {
	Name            string `json:",omitempty"`
	DBType          string `json:",omitempty"`
	IsNullable      bool   `json:",omitempty"`
	IsAutoincrement bool   `json:",omitempty"`
	IsPrimaryKey    bool   `json:",omitempty"`

	// generated Contents
	Title       string `json:"-"`
	GoType      string `json:"-"`
	goZero      string `json:"-"`
	ParamName   string `json:"-"`
	mappingFunc string `json:"-"`
	jsonFunc    string `json:"-"`
}

// IgnoreFields is used to ignore fields for specific statements
type IgnoreFields struct {
	Upsert []string `json:",omitempty"`
}

// Index type
type Index struct {
	Name     string   `json:",omitempty"`
	Fields   []string `json:",omitempty"`
	IsUnique bool     `json:",omitempty"`
	Generate bool
}

type indexField struct {
	IndexName string
	ColName   string
	IsUnique  bool
}
