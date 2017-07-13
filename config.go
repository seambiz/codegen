package codegen

// Config base template input
type Config struct {
	MetaLinter  string   `toml:",omitempty"`
	LintPackage string   `toml:",omitempty"`
	Package     string   `toml:",omitempty"`
	DirOut      string   `toml:",omitempty"`
	FilePattern string   `toml:",omitempty"`
	Templates   []string `toml:",omitempty"`

	Schemas  []*Schema `toml:",omitempty"`
	Database db
}

type db struct {
	DSN     string
	Schemas []string
}

// Schema single schema type
type Schema struct {
	Name   string   `toml:",omitempty"`
	Tables []*Table `toml:",omitempty"`
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
	Name      string   `toml:",omitempty"`
	Templates []string `toml:",omitempty"`

	Fields      []*Field      `toml:",omitempty"`
	Indices     []*Index      `toml:",omitempty"`
	Ignores     IgnoreFields  `toml:",omitempty"`
	ForeignKeys []*ForeignKey `toml:",omitempty"`

	// generated Contents
	title         string
	lower         string
	receiver      string
	initials      string
	store         string
	storeReceiver string
	fieldMapping  map[string]int
	pkFields      []*Field
	otherFields   []*Field
}

// ForeignKey type
type ForeignKey struct {
	Name       string   `toml:",omitempty"`
	Fields     []string `toml:",omitempty"`
	RefSchema  string   `toml:",omitempty"`
	RefTable   string   `toml:",omitempty"`
	RefFields  []string `toml:",omitempty"`
	IsUnique   bool     `toml:",omitempty"`
	CustomName string   `toml:",omitempty"`

	ForeignKeys []*ForeignKey `toml:",omitempty"`
}

// Field type
type Field struct {
	Name            string `toml:",omitempty"`
	DBType          string `toml:",omitempty"`
	IsNullable      bool   `toml:",omitempty"`
	IsAutoincrement bool   `toml:",omitempty"`
	IsPrimaryKey    bool   `toml:",omitempty"`

	// generated Contents
	title       string
	goType      string
	goZero      string
	paramName   string
	mappingFunc string
}

// IgnoreFields is used to ignore fields for specific statements
type IgnoreFields struct {
	Upsert []string `toml:",omitempty"`
}

// Index type
type Index struct {
	Name     string   `toml:",omitempty"`
	Fields   []string `toml:",omitempty"`
	IsUnique bool     `toml:",omitempty"`
}

type indexField struct {
	IndexName string
	ColName   string
	IsUnique  bool
}
