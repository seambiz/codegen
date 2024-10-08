package codegen

import (
	"math/big"
	"time"
)

// GENERATED BY CODEGEN.

/* Columns represents a row from COLUMNS. */
type Columns struct {
	TableCatalog           string  `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema            string  `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName              string  `json:"TABLE_NAME" db:"table_name"`
	ColumnName             *string `json:"COLUMN_NAME" db:"column_name"`
	OrdinalPosition        uint    `json:"ORDINAL_POSITION" db:"ordinal_position"`
	ColumnDefault          *string `json:"COLUMN_DEFAULT" db:"column_default"`
	IsNullable             string  `json:"IS_NULLABLE" db:"is_nullable"`
	DataType               *string `json:"DATA_TYPE" db:"data_type"`
	CharacterMaximumLength *int64  `json:"CHARACTER_MAXIMUM_LENGTH" db:"character_maximum_length"`
	CharacterOctetLength   *int64  `json:"CHARACTER_OCTET_LENGTH" db:"character_octet_length"`
	NumericPrecision       *uint64 `json:"NUMERIC_PRECISION" db:"numeric_precision"`
	NumericScale           *uint64 `json:"NUMERIC_SCALE" db:"numeric_scale"`
	DatetimePrecision      *uint   `json:"DATETIME_PRECISION" db:"datetime_precision"`
	CharacterSetName       *string `json:"CHARACTER_SET_NAME" db:"character_set_name"`
	CollationName          *string `json:"COLLATION_NAME" db:"collation_name"`
	ColumnType             string  `json:"COLUMN_TYPE" db:"column_type"`
	ColumnKey              string  `json:"COLUMN_KEY" db:"column_key"`
	Extra                  *string `json:"EXTRA" db:"extra"`
	Privileges             *string `json:"PRIVILEGES" db:"privileges"`
	ColumnComment          string  `json:"COLUMN_COMMENT" db:"column_comment"`
	GenerationExpression   string  `json:"GENERATION_EXPRESSION" db:"generation_expression"`
}

/* KeyColumnUsage represents a row from KEY_COLUMN_USAGE. */
type KeyColumnUsage struct {
	ConstraintCatalog          string  `json:"CONSTRAINT_CATALOG" db:"constraint_catalog"`
	ConstraintSchema           string  `json:"CONSTRAINT_SCHEMA" db:"constraint_schema"`
	ConstraintName             *string `json:"CONSTRAINT_NAME" db:"constraint_name"`
	TableCatalog               string  `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema                string  `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName                  string  `json:"TABLE_NAME" db:"table_name"`
	ColumnName                 *string `json:"COLUMN_NAME" db:"column_name"`
	OrdinalPosition            uint    `json:"ORDINAL_POSITION" db:"ordinal_position"`
	PositionInUniqueConstraint *uint   `json:"POSITION_IN_UNIQUE_CONSTRAINT" db:"position_in_unique_constraint"`
	ReferencedTableSchema      *string `json:"REFERENCED_TABLE_SCHEMA" db:"referenced_table_schema"`
	ReferencedTableName        *string `json:"REFERENCED_TABLE_NAME" db:"referenced_table_name"`
	ReferencedColumnName       *string `json:"REFERENCED_COLUMN_NAME" db:"referenced_column_name"`
}

/* Statistics represents a row from STATISTICS. */
type Statistics struct {
	TableCatalog string  `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema  string  `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName    string  `json:"TABLE_NAME" db:"table_name"`
	NonUnique    int     `json:"NON_UNIQUE" db:"non_unique"`
	IndexSchema  string  `json:"INDEX_SCHEMA" db:"index_schema"`
	IndexName    *string `json:"INDEX_NAME" db:"index_name"`
	SeqInIndex   uint    `json:"SEQ_IN_INDEX" db:"seq_in_index"`
	ColumnName   *string `json:"COLUMN_NAME" db:"column_name"`
	Collation    *string `json:"COLLATION" db:"collation"`
	Cardinality  *int64  `json:"CARDINALITY" db:"cardinality"`
	SubPart      *int64  `json:"SUB_PART" db:"sub_part"`
	Packed       *[]byte `json:"PACKED" db:"packed"`
	Nullable     string  `json:"NULLABLE" db:"nullable"`
	IndexType    string  `json:"INDEX_TYPE" db:"index_type"`
	Comment      string  `json:"COMMENT" db:"comment"`
	IndexComment string  `json:"INDEX_COMMENT" db:"index_comment"`
	IsVisible    string  `json:"IS_VISIBLE" db:"is_visible"`
	Expression   *string `json:"EXPRESSION" db:"expression"`
}

/* Tables represents a row from TABLES. */
type Tables struct {
	TableCatalog   string     `json:"TABLE_CATALOG" db:"table_catalog"`
	TableSchema    string     `json:"TABLE_SCHEMA" db:"table_schema"`
	TableName      string     `json:"TABLE_NAME" db:"table_name"`
	TableType      string     `json:"TABLE_TYPE" db:"table_type"`
	Engine         *string    `json:"ENGINE" db:"engine"`
	Version        *int       `json:"VERSION" db:"version"`
	RowFormat      *string    `json:"ROW_FORMAT" db:"row_format"`
	TableRows      *uint64    `json:"TABLE_ROWS" db:"table_rows"`
	AvgRowLength   *uint64    `json:"AVG_ROW_LENGTH" db:"avg_row_length"`
	DataLength     *uint64    `json:"DATA_LENGTH" db:"data_length"`
	MaxDataLength  *uint64    `json:"MAX_DATA_LENGTH" db:"max_data_length"`
	IndexLength    *uint64    `json:"INDEX_LENGTH" db:"index_length"`
	DataFree       *uint64    `json:"DATA_FREE" db:"data_free"`
	AutoIncrement  *uint64    `json:"AUTO_INCREMENT" db:"auto_increment"`
	CreateTime     time.Time  `json:"CREATE_TIME" db:"create_time"`
	UpdateTime     *time.Time `json:"UPDATE_TIME" db:"update_time"`
	CheckTime      *time.Time `json:"CHECK_TIME" db:"check_time"`
	TableCollation *string    `json:"TABLE_COLLATION" db:"table_collation"`
	Checksum       *int64     `json:"CHECKSUM" db:"checksum"`
	CreateOptions  *string    `json:"CREATE_OPTIONS" db:"create_options"`
	TableComment   *string    `json:"TABLE_COMMENT" db:"table_comment"`
}

/* Person represents a row from person. */
type Person struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Pets []*Pet `json:"pets,omitempty"`
}

/* Tag represents a row from tag. */
type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/* Pet represents a row from pet. */
type Pet struct {
	ID       int    `json:"id" db:"id"`
	PersonID int    `json:"person_id" db:"person_id"`
	TagID    int    `json:"tag_id" db:"tag_id"`
	Species  string `json:"species" db:"species"`

	BelongsTo *Person `json:"belongsto,omitempty"`
	HasTag    *Tag    `json:"hastag,omitempty"`
}

/* Extensive represents a row from extensive. */
type Extensive struct {
	ID               int        `json:"id" db:"id"`
	Tinyint          int        `json:"tinyint" db:"tinyint"`
	TinyintUnsigned  uint       `json:"tinyint_unsigned" db:"tinyint_unsigned"`
	TinyintNull      *int       `json:"tinyint_null" db:"tinyint_null"`
	Smallint         int        `json:"smallint" db:"smallint"`
	SmallintUnsigned uint       `json:"smallint_unsigned" db:"smallint_unsigned"`
	SmallintNull     *int       `json:"smallint_null" db:"smallint_null"`
	Int              int        `json:"int" db:"int"`
	IntNull          *int       `json:"int_null" db:"int_null"`
	IntUnsigned      uint       `json:"int_unsigned" db:"int_unsigned"`
	Bigint           int64      `json:"bigint" db:"bigint"`
	BigintNull       *int64     `json:"bigint_null" db:"bigint_null"`
	BigintUnsigned   uint64     `json:"bigint_unsigned" db:"bigint_unsigned"`
	Varchar          string     `json:"varchar" db:"varchar"`
	VarcharNull      *string    `json:"varchar_null" db:"varchar_null"`
	Float            float32    `json:"float" db:"float"`
	FloatNull        *float32   `json:"float_null" db:"float_null"`
	Double           float64    `json:"double" db:"double"`
	DoubleNull       *float64   `json:"double_null" db:"double_null"`
	Decimal          float64    `json:"decimal" db:"decimal"`
	DecimalNull      *float64   `json:"decimal_null" db:"decimal_null"`
	Numeric          float64    `json:"numeric" db:"numeric"`
	NumericNull      *float64   `json:"numeric_null" db:"numeric_null"`
	CreatedAt        uint       `json:"created_at" db:"created_at"`
	UpdatedAt        uint       `json:"updated_at" db:"updated_at"`
	Tinyint1         bool       `json:"tinyint1" db:"tinyint1"`
	Tinyint1Null     *bool      `json:"tinyint1_null" db:"tinyint1_null"`
	Year             int        `json:"year" db:"year"`
	YearNull         *int       `json:"year_null" db:"year_null"`
	Date             time.Time  `json:"date" db:"date"`
	DateNull         *time.Time `json:"date_null" db:"date_null"`
	Time             time.Time  `json:"time" db:"time"`
	TimeNull         *time.Time `json:"time_null" db:"time_null"`
	Datetime         time.Time  `json:"datetime" db:"datetime"`
	DatetimeNull     *time.Time `json:"datetime_null" db:"datetime_null"`
	Timestamp        time.Time  `json:"timestamp" db:"timestamp"`
	TimestampNull    *time.Time `json:"timestamp_null" db:"timestamp_null"`
	Char             string     `json:"char" db:"char"`
	CharNull         *string    `json:"char_null" db:"char_null"`
	Tinytext         string     `json:"tinytext" db:"tinytext"`
	TinytextNull     *string    `json:"tinytext_null" db:"tinytext_null"`
	Text             string     `json:"text" db:"text"`
	TextNull         *string    `json:"text_null" db:"text_null"`
	Mediumtext       string     `json:"mediumtext" db:"mediumtext"`
	MediumtextNull   *string    `json:"mediumtext_null" db:"mediumtext_null"`
	Longtext         string     `json:"longtext" db:"longtext"`
	LongtextNull     *string    `json:"longtext_null" db:"longtext_null"`
	Binary           []byte     `json:"binary" db:"binary"`
	BinaryNull       *[]byte    `json:"binary_null" db:"binary_null"`
	Varbinary        []byte     `json:"varbinary" db:"varbinary"`
	VarbinaryNull    *[]byte    `json:"varbinary_null" db:"varbinary_null"`
	Tinyblob         []byte     `json:"tinyblob" db:"tinyblob"`
	TinyblobNull     *[]byte    `json:"tinyblob_null" db:"tinyblob_null"`
	Blob             []byte     `json:"blob" db:"blob"`
	BlobNull         *[]byte    `json:"blob_null" db:"blob_null"`
	Mediumblob       []byte     `json:"mediumblob" db:"mediumblob"`
	MediumblobNull   *[]byte    `json:"mediumblob_null" db:"mediumblob_null"`
	Longblob         []byte     `json:"longblob" db:"longblob"`
	LongblobNull     *[]byte    `json:"longblob_null" db:"longblob_null"`
}

/* ColumnsPartial is used for updating specific columns from COLUMNS. */
type ColumnsPartial struct {
	Columns
	Touched big.Int
}

func (p *ColumnsPartial) SetTableCatalog(param string) {
	p.TableCatalog = param
	p.Touched.SetBit(&p.Touched, Columns_TableCatalog, 1)
}

func (p *ColumnsPartial) SetTableSchema(param string) {
	p.TableSchema = param
	p.Touched.SetBit(&p.Touched, Columns_TableSchema, 1)
}

func (p *ColumnsPartial) SetTableName(param string) {
	p.TableName = param
	p.Touched.SetBit(&p.Touched, Columns_TableName, 1)
}

func (p *ColumnsPartial) SetColumnName(param *string) {
	p.ColumnName = param
	p.Touched.SetBit(&p.Touched, Columns_ColumnName, 1)
}

func (p *ColumnsPartial) SetOrdinalPosition(param uint) {
	p.OrdinalPosition = param
	p.Touched.SetBit(&p.Touched, Columns_OrdinalPosition, 1)
}

func (p *ColumnsPartial) SetColumnDefault(param *string) {
	p.ColumnDefault = param
	p.Touched.SetBit(&p.Touched, Columns_ColumnDefault, 1)
}

func (p *ColumnsPartial) SetIsNullable(param string) {
	p.IsNullable = param
	p.Touched.SetBit(&p.Touched, Columns_IsNullable, 1)
}

func (p *ColumnsPartial) SetDataType(param *string) {
	p.DataType = param
	p.Touched.SetBit(&p.Touched, Columns_DataType, 1)
}

func (p *ColumnsPartial) SetCharacterMaximumLength(param *int64) {
	p.CharacterMaximumLength = param
	p.Touched.SetBit(&p.Touched, Columns_CharacterMaximumLength, 1)
}

func (p *ColumnsPartial) SetCharacterOctetLength(param *int64) {
	p.CharacterOctetLength = param
	p.Touched.SetBit(&p.Touched, Columns_CharacterOctetLength, 1)
}

func (p *ColumnsPartial) SetNumericPrecision(param *uint64) {
	p.NumericPrecision = param
	p.Touched.SetBit(&p.Touched, Columns_NumericPrecision, 1)
}

func (p *ColumnsPartial) SetNumericScale(param *uint64) {
	p.NumericScale = param
	p.Touched.SetBit(&p.Touched, Columns_NumericScale, 1)
}

func (p *ColumnsPartial) SetDatetimePrecision(param *uint) {
	p.DatetimePrecision = param
	p.Touched.SetBit(&p.Touched, Columns_DatetimePrecision, 1)
}

func (p *ColumnsPartial) SetCharacterSetName(param *string) {
	p.CharacterSetName = param
	p.Touched.SetBit(&p.Touched, Columns_CharacterSetName, 1)
}

func (p *ColumnsPartial) SetCollationName(param *string) {
	p.CollationName = param
	p.Touched.SetBit(&p.Touched, Columns_CollationName, 1)
}

func (p *ColumnsPartial) SetColumnType(param string) {
	p.ColumnType = param
	p.Touched.SetBit(&p.Touched, Columns_ColumnType, 1)
}

func (p *ColumnsPartial) SetColumnKey(param string) {
	p.ColumnKey = param
	p.Touched.SetBit(&p.Touched, Columns_ColumnKey, 1)
}

func (p *ColumnsPartial) SetExtra(param *string) {
	p.Extra = param
	p.Touched.SetBit(&p.Touched, Columns_Extra, 1)
}

func (p *ColumnsPartial) SetPrivileges(param *string) {
	p.Privileges = param
	p.Touched.SetBit(&p.Touched, Columns_Privileges, 1)
}

func (p *ColumnsPartial) SetColumnComment(param string) {
	p.ColumnComment = param
	p.Touched.SetBit(&p.Touched, Columns_ColumnComment, 1)
}

func (p *ColumnsPartial) SetGenerationExpression(param string) {
	p.GenerationExpression = param
	p.Touched.SetBit(&p.Touched, Columns_GenerationExpression, 1)
}

/* KeyColumnUsagePartial is used for updating specific columns from KEY_COLUMN_USAGE. */
type KeyColumnUsagePartial struct {
	KeyColumnUsage
	Touched big.Int
}

func (p *KeyColumnUsagePartial) SetConstraintCatalog(param string) {
	p.ConstraintCatalog = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ConstraintCatalog, 1)
}

func (p *KeyColumnUsagePartial) SetConstraintSchema(param string) {
	p.ConstraintSchema = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ConstraintSchema, 1)
}

func (p *KeyColumnUsagePartial) SetConstraintName(param *string) {
	p.ConstraintName = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ConstraintName, 1)
}

func (p *KeyColumnUsagePartial) SetTableCatalog(param string) {
	p.TableCatalog = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_TableCatalog, 1)
}

func (p *KeyColumnUsagePartial) SetTableSchema(param string) {
	p.TableSchema = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_TableSchema, 1)
}

func (p *KeyColumnUsagePartial) SetTableName(param string) {
	p.TableName = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_TableName, 1)
}

func (p *KeyColumnUsagePartial) SetColumnName(param *string) {
	p.ColumnName = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ColumnName, 1)
}

func (p *KeyColumnUsagePartial) SetOrdinalPosition(param uint) {
	p.OrdinalPosition = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_OrdinalPosition, 1)
}

func (p *KeyColumnUsagePartial) SetPositionInUniqueConstraint(param *uint) {
	p.PositionInUniqueConstraint = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_PositionInUniqueConstraint, 1)
}

func (p *KeyColumnUsagePartial) SetReferencedTableSchema(param *string) {
	p.ReferencedTableSchema = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ReferencedTableSchema, 1)
}

func (p *KeyColumnUsagePartial) SetReferencedTableName(param *string) {
	p.ReferencedTableName = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ReferencedTableName, 1)
}

func (p *KeyColumnUsagePartial) SetReferencedColumnName(param *string) {
	p.ReferencedColumnName = param
	p.Touched.SetBit(&p.Touched, KeyColumnUsage_ReferencedColumnName, 1)
}

/* StatisticsPartial is used for updating specific columns from STATISTICS. */
type StatisticsPartial struct {
	Statistics
	Touched big.Int
}

func (p *StatisticsPartial) SetTableCatalog(param string) {
	p.TableCatalog = param
	p.Touched.SetBit(&p.Touched, Statistics_TableCatalog, 1)
}

func (p *StatisticsPartial) SetTableSchema(param string) {
	p.TableSchema = param
	p.Touched.SetBit(&p.Touched, Statistics_TableSchema, 1)
}

func (p *StatisticsPartial) SetTableName(param string) {
	p.TableName = param
	p.Touched.SetBit(&p.Touched, Statistics_TableName, 1)
}

func (p *StatisticsPartial) SetNonUnique(param int) {
	p.NonUnique = param
	p.Touched.SetBit(&p.Touched, Statistics_NonUnique, 1)
}

func (p *StatisticsPartial) SetIndexSchema(param string) {
	p.IndexSchema = param
	p.Touched.SetBit(&p.Touched, Statistics_IndexSchema, 1)
}

func (p *StatisticsPartial) SetIndexName(param *string) {
	p.IndexName = param
	p.Touched.SetBit(&p.Touched, Statistics_IndexName, 1)
}

func (p *StatisticsPartial) SetSeqInIndex(param uint) {
	p.SeqInIndex = param
	p.Touched.SetBit(&p.Touched, Statistics_SeqInIndex, 1)
}

func (p *StatisticsPartial) SetColumnName(param *string) {
	p.ColumnName = param
	p.Touched.SetBit(&p.Touched, Statistics_ColumnName, 1)
}

func (p *StatisticsPartial) SetCollation(param *string) {
	p.Collation = param
	p.Touched.SetBit(&p.Touched, Statistics_Collation, 1)
}

func (p *StatisticsPartial) SetCardinality(param *int64) {
	p.Cardinality = param
	p.Touched.SetBit(&p.Touched, Statistics_Cardinality, 1)
}

func (p *StatisticsPartial) SetSubPart(param *int64) {
	p.SubPart = param
	p.Touched.SetBit(&p.Touched, Statistics_SubPart, 1)
}

func (p *StatisticsPartial) SetPacked(param *[]byte) {
	p.Packed = param
	p.Touched.SetBit(&p.Touched, Statistics_Packed, 1)
}

func (p *StatisticsPartial) SetNullable(param string) {
	p.Nullable = param
	p.Touched.SetBit(&p.Touched, Statistics_Nullable, 1)
}

func (p *StatisticsPartial) SetIndexType(param string) {
	p.IndexType = param
	p.Touched.SetBit(&p.Touched, Statistics_IndexType, 1)
}

func (p *StatisticsPartial) SetComment(param string) {
	p.Comment = param
	p.Touched.SetBit(&p.Touched, Statistics_Comment, 1)
}

func (p *StatisticsPartial) SetIndexComment(param string) {
	p.IndexComment = param
	p.Touched.SetBit(&p.Touched, Statistics_IndexComment, 1)
}

func (p *StatisticsPartial) SetIsVisible(param string) {
	p.IsVisible = param
	p.Touched.SetBit(&p.Touched, Statistics_IsVisible, 1)
}

func (p *StatisticsPartial) SetExpression(param *string) {
	p.Expression = param
	p.Touched.SetBit(&p.Touched, Statistics_Expression, 1)
}

/* TablesPartial is used for updating specific columns from TABLES. */
type TablesPartial struct {
	Tables
	Touched big.Int
}

func (p *TablesPartial) SetTableCatalog(param string) {
	p.TableCatalog = param
	p.Touched.SetBit(&p.Touched, Tables_TableCatalog, 1)
}

func (p *TablesPartial) SetTableSchema(param string) {
	p.TableSchema = param
	p.Touched.SetBit(&p.Touched, Tables_TableSchema, 1)
}

func (p *TablesPartial) SetTableName(param string) {
	p.TableName = param
	p.Touched.SetBit(&p.Touched, Tables_TableName, 1)
}

func (p *TablesPartial) SetTableType(param string) {
	p.TableType = param
	p.Touched.SetBit(&p.Touched, Tables_TableType, 1)
}

func (p *TablesPartial) SetEngine(param *string) {
	p.Engine = param
	p.Touched.SetBit(&p.Touched, Tables_Engine, 1)
}

func (p *TablesPartial) SetVersion(param *int) {
	p.Version = param
	p.Touched.SetBit(&p.Touched, Tables_Version, 1)
}

func (p *TablesPartial) SetRowFormat(param *string) {
	p.RowFormat = param
	p.Touched.SetBit(&p.Touched, Tables_RowFormat, 1)
}

func (p *TablesPartial) SetTableRows(param *uint64) {
	p.TableRows = param
	p.Touched.SetBit(&p.Touched, Tables_TableRows, 1)
}

func (p *TablesPartial) SetAvgRowLength(param *uint64) {
	p.AvgRowLength = param
	p.Touched.SetBit(&p.Touched, Tables_AvgRowLength, 1)
}

func (p *TablesPartial) SetDataLength(param *uint64) {
	p.DataLength = param
	p.Touched.SetBit(&p.Touched, Tables_DataLength, 1)
}

func (p *TablesPartial) SetMaxDataLength(param *uint64) {
	p.MaxDataLength = param
	p.Touched.SetBit(&p.Touched, Tables_MaxDataLength, 1)
}

func (p *TablesPartial) SetIndexLength(param *uint64) {
	p.IndexLength = param
	p.Touched.SetBit(&p.Touched, Tables_IndexLength, 1)
}

func (p *TablesPartial) SetDataFree(param *uint64) {
	p.DataFree = param
	p.Touched.SetBit(&p.Touched, Tables_DataFree, 1)
}

func (p *TablesPartial) SetAutoIncrement(param *uint64) {
	p.AutoIncrement = param
	p.Touched.SetBit(&p.Touched, Tables_AutoIncrement, 1)
}

func (p *TablesPartial) SetCreateTime(param time.Time) {
	p.CreateTime = param
	p.Touched.SetBit(&p.Touched, Tables_CreateTime, 1)
}

func (p *TablesPartial) SetUpdateTime(param *time.Time) {
	p.UpdateTime = param
	p.Touched.SetBit(&p.Touched, Tables_UpdateTime, 1)
}

func (p *TablesPartial) SetCheckTime(param *time.Time) {
	p.CheckTime = param
	p.Touched.SetBit(&p.Touched, Tables_CheckTime, 1)
}

func (p *TablesPartial) SetTableCollation(param *string) {
	p.TableCollation = param
	p.Touched.SetBit(&p.Touched, Tables_TableCollation, 1)
}

func (p *TablesPartial) SetChecksum(param *int64) {
	p.Checksum = param
	p.Touched.SetBit(&p.Touched, Tables_Checksum, 1)
}

func (p *TablesPartial) SetCreateOptions(param *string) {
	p.CreateOptions = param
	p.Touched.SetBit(&p.Touched, Tables_CreateOptions, 1)
}

func (p *TablesPartial) SetTableComment(param *string) {
	p.TableComment = param
	p.Touched.SetBit(&p.Touched, Tables_TableComment, 1)
}

/* PersonPartial is used for updating specific columns from person. */
type PersonPartial struct {
	Person
	Touched big.Int
}

func (p *PersonPartial) SetID(param int) {
	p.ID = param
	p.Touched.SetBit(&p.Touched, Person_ID, 1)
}

func (p *PersonPartial) SetName(param string) {
	p.Name = param
	p.Touched.SetBit(&p.Touched, Person_Name, 1)
}

/* TagPartial is used for updating specific columns from tag. */
type TagPartial struct {
	Tag
	Touched big.Int
}

func (p *TagPartial) SetID(param int) {
	p.ID = param
	p.Touched.SetBit(&p.Touched, Tag_ID, 1)
}

func (p *TagPartial) SetName(param string) {
	p.Name = param
	p.Touched.SetBit(&p.Touched, Tag_Name, 1)
}

/* PetPartial is used for updating specific columns from pet. */
type PetPartial struct {
	Pet
	Touched big.Int
}

func (p *PetPartial) SetID(param int) {
	p.ID = param
	p.Touched.SetBit(&p.Touched, Pet_ID, 1)
}

func (p *PetPartial) SetPersonID(param int) {
	p.PersonID = param
	p.Touched.SetBit(&p.Touched, Pet_PersonID, 1)
}

func (p *PetPartial) SetTagID(param int) {
	p.TagID = param
	p.Touched.SetBit(&p.Touched, Pet_TagID, 1)
}

func (p *PetPartial) SetSpecies(param string) {
	p.Species = param
	p.Touched.SetBit(&p.Touched, Pet_Species, 1)
}

/* ExtensivePartial is used for updating specific columns from extensive. */
type ExtensivePartial struct {
	Extensive
	Touched big.Int
}

func (p *ExtensivePartial) SetID(param int) {
	p.ID = param
	p.Touched.SetBit(&p.Touched, Extensive_ID, 1)
}

func (p *ExtensivePartial) SetTinyint(param int) {
	p.Tinyint = param
	p.Touched.SetBit(&p.Touched, Extensive_Tinyint, 1)
}

func (p *ExtensivePartial) SetTinyintUnsigned(param uint) {
	p.TinyintUnsigned = param
	p.Touched.SetBit(&p.Touched, Extensive_TinyintUnsigned, 1)
}

func (p *ExtensivePartial) SetTinyintNull(param *int) {
	p.TinyintNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TinyintNull, 1)
}

func (p *ExtensivePartial) SetSmallint(param int) {
	p.Smallint = param
	p.Touched.SetBit(&p.Touched, Extensive_Smallint, 1)
}

func (p *ExtensivePartial) SetSmallintUnsigned(param uint) {
	p.SmallintUnsigned = param
	p.Touched.SetBit(&p.Touched, Extensive_SmallintUnsigned, 1)
}

func (p *ExtensivePartial) SetSmallintNull(param *int) {
	p.SmallintNull = param
	p.Touched.SetBit(&p.Touched, Extensive_SmallintNull, 1)
}

func (p *ExtensivePartial) SetInt(param int) {
	p.Int = param
	p.Touched.SetBit(&p.Touched, Extensive_Int, 1)
}

func (p *ExtensivePartial) SetIntNull(param *int) {
	p.IntNull = param
	p.Touched.SetBit(&p.Touched, Extensive_IntNull, 1)
}

func (p *ExtensivePartial) SetIntUnsigned(param uint) {
	p.IntUnsigned = param
	p.Touched.SetBit(&p.Touched, Extensive_IntUnsigned, 1)
}

func (p *ExtensivePartial) SetBigint(param int64) {
	p.Bigint = param
	p.Touched.SetBit(&p.Touched, Extensive_Bigint, 1)
}

func (p *ExtensivePartial) SetBigintNull(param *int64) {
	p.BigintNull = param
	p.Touched.SetBit(&p.Touched, Extensive_BigintNull, 1)
}

func (p *ExtensivePartial) SetBigintUnsigned(param uint64) {
	p.BigintUnsigned = param
	p.Touched.SetBit(&p.Touched, Extensive_BigintUnsigned, 1)
}

func (p *ExtensivePartial) SetVarchar(param string) {
	p.Varchar = param
	p.Touched.SetBit(&p.Touched, Extensive_Varchar, 1)
}

func (p *ExtensivePartial) SetVarcharNull(param *string) {
	p.VarcharNull = param
	p.Touched.SetBit(&p.Touched, Extensive_VarcharNull, 1)
}

func (p *ExtensivePartial) SetFloat(param float32) {
	p.Float = param
	p.Touched.SetBit(&p.Touched, Extensive_Float, 1)
}

func (p *ExtensivePartial) SetFloatNull(param *float32) {
	p.FloatNull = param
	p.Touched.SetBit(&p.Touched, Extensive_FloatNull, 1)
}

func (p *ExtensivePartial) SetDouble(param float64) {
	p.Double = param
	p.Touched.SetBit(&p.Touched, Extensive_Double, 1)
}

func (p *ExtensivePartial) SetDoubleNull(param *float64) {
	p.DoubleNull = param
	p.Touched.SetBit(&p.Touched, Extensive_DoubleNull, 1)
}

func (p *ExtensivePartial) SetDecimal(param float64) {
	p.Decimal = param
	p.Touched.SetBit(&p.Touched, Extensive_Decimal, 1)
}

func (p *ExtensivePartial) SetDecimalNull(param *float64) {
	p.DecimalNull = param
	p.Touched.SetBit(&p.Touched, Extensive_DecimalNull, 1)
}

func (p *ExtensivePartial) SetNumeric(param float64) {
	p.Numeric = param
	p.Touched.SetBit(&p.Touched, Extensive_Numeric, 1)
}

func (p *ExtensivePartial) SetNumericNull(param *float64) {
	p.NumericNull = param
	p.Touched.SetBit(&p.Touched, Extensive_NumericNull, 1)
}

func (p *ExtensivePartial) SetCreatedAt(param uint) {
	p.CreatedAt = param
	p.Touched.SetBit(&p.Touched, Extensive_CreatedAt, 1)
}

func (p *ExtensivePartial) SetUpdatedAt(param uint) {
	p.UpdatedAt = param
	p.Touched.SetBit(&p.Touched, Extensive_UpdatedAt, 1)
}

func (p *ExtensivePartial) SetTinyint1(param bool) {
	p.Tinyint1 = param
	p.Touched.SetBit(&p.Touched, Extensive_Tinyint1, 1)
}

func (p *ExtensivePartial) SetTinyint1Null(param *bool) {
	p.Tinyint1Null = param
	p.Touched.SetBit(&p.Touched, Extensive_Tinyint1Null, 1)
}

func (p *ExtensivePartial) SetYear(param int) {
	p.Year = param
	p.Touched.SetBit(&p.Touched, Extensive_Year, 1)
}

func (p *ExtensivePartial) SetYearNull(param *int) {
	p.YearNull = param
	p.Touched.SetBit(&p.Touched, Extensive_YearNull, 1)
}

func (p *ExtensivePartial) SetDate(param time.Time) {
	p.Date = param
	p.Touched.SetBit(&p.Touched, Extensive_Date, 1)
}

func (p *ExtensivePartial) SetDateNull(param *time.Time) {
	p.DateNull = param
	p.Touched.SetBit(&p.Touched, Extensive_DateNull, 1)
}

func (p *ExtensivePartial) SetTime(param time.Time) {
	p.Time = param
	p.Touched.SetBit(&p.Touched, Extensive_Time, 1)
}

func (p *ExtensivePartial) SetTimeNull(param *time.Time) {
	p.TimeNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TimeNull, 1)
}

func (p *ExtensivePartial) SetDatetime(param time.Time) {
	p.Datetime = param
	p.Touched.SetBit(&p.Touched, Extensive_Datetime, 1)
}

func (p *ExtensivePartial) SetDatetimeNull(param *time.Time) {
	p.DatetimeNull = param
	p.Touched.SetBit(&p.Touched, Extensive_DatetimeNull, 1)
}

func (p *ExtensivePartial) SetTimestamp(param time.Time) {
	p.Timestamp = param
	p.Touched.SetBit(&p.Touched, Extensive_Timestamp, 1)
}

func (p *ExtensivePartial) SetTimestampNull(param *time.Time) {
	p.TimestampNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TimestampNull, 1)
}

func (p *ExtensivePartial) SetChar(param string) {
	p.Char = param
	p.Touched.SetBit(&p.Touched, Extensive_Char, 1)
}

func (p *ExtensivePartial) SetCharNull(param *string) {
	p.CharNull = param
	p.Touched.SetBit(&p.Touched, Extensive_CharNull, 1)
}

func (p *ExtensivePartial) SetTinytext(param string) {
	p.Tinytext = param
	p.Touched.SetBit(&p.Touched, Extensive_Tinytext, 1)
}

func (p *ExtensivePartial) SetTinytextNull(param *string) {
	p.TinytextNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TinytextNull, 1)
}

func (p *ExtensivePartial) SetText(param string) {
	p.Text = param
	p.Touched.SetBit(&p.Touched, Extensive_Text, 1)
}

func (p *ExtensivePartial) SetTextNull(param *string) {
	p.TextNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TextNull, 1)
}

func (p *ExtensivePartial) SetMediumtext(param string) {
	p.Mediumtext = param
	p.Touched.SetBit(&p.Touched, Extensive_Mediumtext, 1)
}

func (p *ExtensivePartial) SetMediumtextNull(param *string) {
	p.MediumtextNull = param
	p.Touched.SetBit(&p.Touched, Extensive_MediumtextNull, 1)
}

func (p *ExtensivePartial) SetLongtext(param string) {
	p.Longtext = param
	p.Touched.SetBit(&p.Touched, Extensive_Longtext, 1)
}

func (p *ExtensivePartial) SetLongtextNull(param *string) {
	p.LongtextNull = param
	p.Touched.SetBit(&p.Touched, Extensive_LongtextNull, 1)
}

func (p *ExtensivePartial) SetBinary(param []byte) {
	p.Binary = param
	p.Touched.SetBit(&p.Touched, Extensive_Binary, 1)
}

func (p *ExtensivePartial) SetBinaryNull(param *[]byte) {
	p.BinaryNull = param
	p.Touched.SetBit(&p.Touched, Extensive_BinaryNull, 1)
}

func (p *ExtensivePartial) SetVarbinary(param []byte) {
	p.Varbinary = param
	p.Touched.SetBit(&p.Touched, Extensive_Varbinary, 1)
}

func (p *ExtensivePartial) SetVarbinaryNull(param *[]byte) {
	p.VarbinaryNull = param
	p.Touched.SetBit(&p.Touched, Extensive_VarbinaryNull, 1)
}

func (p *ExtensivePartial) SetTinyblob(param []byte) {
	p.Tinyblob = param
	p.Touched.SetBit(&p.Touched, Extensive_Tinyblob, 1)
}

func (p *ExtensivePartial) SetTinyblobNull(param *[]byte) {
	p.TinyblobNull = param
	p.Touched.SetBit(&p.Touched, Extensive_TinyblobNull, 1)
}

func (p *ExtensivePartial) SetBlob(param []byte) {
	p.Blob = param
	p.Touched.SetBit(&p.Touched, Extensive_Blob, 1)
}

func (p *ExtensivePartial) SetBlobNull(param *[]byte) {
	p.BlobNull = param
	p.Touched.SetBit(&p.Touched, Extensive_BlobNull, 1)
}

func (p *ExtensivePartial) SetMediumblob(param []byte) {
	p.Mediumblob = param
	p.Touched.SetBit(&p.Touched, Extensive_Mediumblob, 1)
}

func (p *ExtensivePartial) SetMediumblobNull(param *[]byte) {
	p.MediumblobNull = param
	p.Touched.SetBit(&p.Touched, Extensive_MediumblobNull, 1)
}

func (p *ExtensivePartial) SetLongblob(param []byte) {
	p.Longblob = param
	p.Touched.SetBit(&p.Touched, Extensive_Longblob, 1)
}

func (p *ExtensivePartial) SetLongblobNull(param *[]byte) {
	p.LongblobNull = param
	p.Touched.SetBit(&p.Touched, Extensive_LongblobNull, 1)
}

// ^^ END OF GENERATED BY CODEGEN. ^^
