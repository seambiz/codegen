package config

// Config base template input
type Config struct {
	MetaLinter         string   `json:",omitempty"`
	LintPackage        string   `json:",omitempty"`
	GoFmtCmd           string   `json:",omitempty"`
	Package            string   `json:",omitempty"`
	RootPackage        string   `json:",omitempty"`
	RootPackageInclude string   `json:",omitempty"`
	DirOut             string   `json:",omitempty"`
	SubPackage         string   `json:",omitempty"`
	FilePattern        string   `json:",omitempty"`
	Templates          []string `json:",omitempty"`
	TemplateFolder     string   `json:",omitempty"`
	Prefix             string   `json:",omitempty"`
	GenTests           bool     `json:",omitempty"`
	CommentStart       string   `json:",omitempty"`
	CommentEnd         string   `json:",omitempty"`

	Schemas  []*Schema `json:",omitempty"`
	Database db
}

type db struct {
	DSN     string
	Schemas []string
}

// Schema single schema type
type Schema struct {
	Name           string   `json:",omitempty"`
	Title          string   `json:",omitempty"`
	TemplateFolder string   `json:",omitempty"`
	Prefix         string   `json:",omitempty"`
	IsMultiTenant  bool     `json:",omitempty"`
	MTVarName      string   `json:",omitempty"`
	MTVarType      string   `json:",omitempty"`
	MTSchemaFmt    string   `json:",omitempty"`
	NoAudit        []string `json:",omitempty"`
	TableNames     []string `json:",omitempty"`
	Tables         []*Table `json:",omitempty"`
	NoTests        bool     `json:",omitempty"`

	PreparedTemplatefiles map[string][]string `json:"-"`
}

func (c *Config) GetSchema(schema string) *Schema {
	for i := range c.Schemas {
		if c.Schemas[i].Name == schema {
			return c.Schemas[i]
		}
	}
	return nil
}

func (s *Schema) GetTable(table string) *Table {
	for i := range s.Tables {
		if s.Tables[i].Name == table {
			return s.Tables[i]
		}
	}
	return nil
}

type Join struct {
	Alias    string
	Name     string
	Title    string
	Initials string
	Schema   string

	Fields []JoinField
	Table  *Table
}

type JoinField struct {
	Alias    string
	Name     string
	RefAlias string
	RefName  string
}

// Table type
type Table struct {
	Name      string   `json:",omitempty"`
	Templates []string `json:",omitempty"`

	Fields        []*Field      `json:",omitempty"`
	Indices       []*Index      `json:",omitempty"`
	Ignores       IgnoreFields  `json:",omitempty"`
	ForeignKeys   []*ForeignKey `json:",omitempty"`
	Generate      bool
	NoCreate      bool     `json:",omitempty"`
	NoUpdate      bool     `json:",omitempty"`
	NoDelete      bool     `json:",omitempty"`
	NoUpsert      bool     `json:",omitempty"`
	Audit         bool     `json:",omitempty"`
	TemplateFiles []string `json:",omitempty"`

	// generated Contents
	Alias                 string              `json:"-"`
	Schema                string              `json:"-"`
	Title                 string              `json:"-"`
	Joins                 []Join              `json:"-"`
	Lower                 string              `json:",omitempty"`
	Receiver              string              `json:"-"`
	Initials              string              `json:"-"`
	Store                 string              `json:",omitempty"`
	StoreReceiver         string              `json:"-"`
	FieldMapping          map[string]int      `json:"-"`
	PkFields              []*Field            `json:"-"`
	OtherFields           []*Field            `json:"-"`
	ID                    int                 `json:",omitempty"`
	NumFields             int                 `json:",omitempty"`
	NumUniqueFKs          int                 `json:"-"`
	PreparedTemplatefiles map[string][]string `json:"-"`
}

// ForeignKey type
type ForeignKey struct {
	Name          string   `json:",omitempty"`
	Fields        []string `json:",omitempty"`
	RefSchema     string   `json:",omitempty"`
	RefTable      string   `json:",omitempty"`
	RefTableTitle string   `json:",omitempty"`
	RefFields     []string `json:",omitempty"`
	IsUnique      bool     `json:",omitempty"`
	CustomName    string   `json:",omitempty"`

	ForeignKeys []*ForeignKey `json:",omitempty"`

	// generated Contents
	GenTableName string `json:"-"`
	GenTable     *Table `json:"-"`
	GenName      string `json:"-"`
	GenFuncName  string `json:"-"`
}

// Field type
type Field struct {
	Name            string   `json:",omitempty"`
	DBType          string   `json:",omitempty"`
	IsNullable      bool     `json:",omitempty"`
	IsAutoincrement bool     `json:",omitempty"`
	IsPrimaryKey    bool     `json:",omitempty"`
	EnumValues      []string `json:",omitempty"`

	// generated Contents
	Alias       string `json:"-"`
	Title       string `json:"-"`
	GoType      string `json:"-"`
	GoZero      string `json:"-"`
	ParamName   string `json:"-"`
	MappingFunc string `json:"-"`
	JsonFunc    string `json:"-"`
	Default     string `json:"-"`
	NoAudit     bool   `json:"-"`
}

// IgnoreFields is used to ignore fields for specific statements
type IgnoreFields struct {
	Upsert []string `json:",omitempty"`
}

// Index type
type Index struct {
	FuncName   string   `json:"-"`
	CustomName string   `json:",omitempty"`
	Name       string   `json:",omitempty"`
	Fields     []string `json:",omitempty"`
	IsUnique   bool     `json:",omitempty"`
	Generate   bool
}

type IndexField struct {
	IndexName string
	ColName   string
	IsUnique  bool
}
