package codegen

var goTypeMapping = map[string]string{
	// numeric types
	"tinyint":    "int",
	"tinyint(1)": "bool", // special case
	"smallint":   "int",
	"mediumint":  "int",
	"int":        "int",
	"integer":    "int",
	"bigint":     "int64",
	// unsigned only for interger types
	"tinyint unsigned":   "uint",
	"smallint unsigned":  "uint",
	"mediumint unsigned": "uint",
	"int unsigned":       "uint",
	"integer unsigned":   "uint",
	"bigint unsigned":    "uint64",
	// floating point types
	"float":   "float32",
	"double":  "float64",
	"decimal": "float64", // maybe github.com/shopspring/decimal
	"numeric": "float64", // maybe github.com/shopspring/decimal
	"bit":     "bool",    // could be used differntly

	// date types
	"year":      "int",
	"date":      "time.Time",
	"time":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",

	// string types
	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",

	// binary types
	"binary":     "[]byte",
	"varbinary":  "[]byte",
	"tinyblob":   "[]byte",
	"blob":       "[]byte",
	"mediumblob": "[]byte",
	"longblob":   "[]byte",
}

var goIsNullMapping = map[string]string{
	// numeric types
	"int":    "sql.NullInt64",
	"uint":   "sql.NullInt64",
	"int64":  "sql.NullInt64",
	"uint64": "sql.NullInt64",
	"bool":   "sql.NullBool",

	"float32":   "sql.NullFloat64",
	"float64":   "sql.NullFloat64",
	"time.Time": "time.Time{}",

	"string": `""`,
}
var goZeroMapping = map[string]string{
	// numeric types
	"int":    "0",
	"uint":   "0",
	"int64":  "0",
	"uint64": "0",
	"bool":   "false",
	"[]byte": "nil",

	"float32":   "0",
	"float64":   "0",
	"time.Time": "*time.Time",

	"string": "sql.NullString",
}
var goDbMappingFunc = map[string]string{
	// numeric types
	"int":    "sdb.ToInt",
	"uint":   "sdb.ToUInt",
	"int64":  "sdb.ToInt64",
	"uint64": "sdb.ToUInt64",
	"bool":   "sdb.ToBool",
	"[]byte": "", //nothing to be done

	"float32":   "sdb.ToFloat32",
	"float64":   "sdb.ToFloat64",
	"time.Time": "sdb.ToTime",

	"string": "string",
}

var goSliceScanMapping = map[string]string{
	// numeric types
	"int":    "int64",
	"uint":   "int64",
	"int64":  "int64",
	"uint64": "uint64",

	"float32": "float64",
	"float64": "float64",

	"string": "[]uint8",
	"[]byte": "[]byte",
	// sepcial
	"bool":      "",
	"time.Time": "",
}

// goReservedNames is a map of of go reserved names to "safe" names.
var goReservedNames = map[string]string{
	"break":       "brk",
	"case":        "cs",
	"chan":        "chn",
	"const":       "cnst",
	"continue":    "cnt",
	"default":     "def",
	"defer":       "dfr",
	"else":        "els",
	"fallthrough": "flthrough",
	"for":         "fr",
	"func":        "fn",
	"go":          "goVal",
	"goto":        "gt",
	"if":          "ifVal",
	"import":      "imp",
	"interface":   "iface",
	"map":         "mp",
	"package":     "pkg",
	"range":       "rnge",
	"return":      "ret",
	"select":      "slct",
	"struct":      "strct",
	"switch":      "swtch",
	"type":        "typ",
	"var":         "vr",

	// go types
	"error":      "e",
	"bool":       "b",
	"string":     "str",
	"byte":       "byt",
	"rune":       "r",
	"uintptr":    "uptr",
	"int":        "i",
	"int8":       "i8",
	"int16":      "i16",
	"int32":      "i32",
	"int64":      "i64",
	"uint":       "u",
	"uint8":      "u8",
	"uint16":     "u16",
	"uint32":     "u32",
	"uint64":     "u64",
	"float32":    "z",
	"float64":    "f",
	"complex64":  "c",
	"complex128": "c128",
}
