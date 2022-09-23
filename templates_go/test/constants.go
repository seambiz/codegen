package codegen

// GENERATED BY CODEGEN.

// constants for each table's columns
const (

	// Columns for table information_schema.COLUMNS
	Columns_TableCatalog = iota
	Columns_TableSchema
	Columns_TableName
	Columns_ColumnName
	Columns_OrdinalPosition
	Columns_ColumnDefault
	Columns_IsNullable
	Columns_DataType
	Columns_CharacterMaximumLength
	Columns_CharacterOctetLength
	Columns_NumericPrecision
	Columns_NumericScale
	Columns_DatetimePrecision
	Columns_CharacterSetName
	Columns_CollationName
	Columns_ColumnType
	Columns_ColumnKey
	Columns_Extra
	Columns_Privileges
	Columns_ColumnComment
	Columns_GenerationExpression
	Columns_SrsID

	// Columns for table information_schema.KEY_COLUMN_USAGE
	KeyColumnUsage_ConstraintCatalog
	KeyColumnUsage_ConstraintSchema
	KeyColumnUsage_ConstraintName
	KeyColumnUsage_TableCatalog
	KeyColumnUsage_TableSchema
	KeyColumnUsage_TableName
	KeyColumnUsage_ColumnName
	KeyColumnUsage_OrdinalPosition
	KeyColumnUsage_PositionInUniqueConstraint
	KeyColumnUsage_ReferencedTableSchema
	KeyColumnUsage_ReferencedTableName
	KeyColumnUsage_ReferencedColumnName

	// Columns for table information_schema.STATISTICS
	Statistics_TableCatalog
	Statistics_TableSchema
	Statistics_TableName
	Statistics_NonUnique
	Statistics_IndexSchema
	Statistics_IndexName
	Statistics_SeqInIndex
	Statistics_ColumnName
	Statistics_Collation
	Statistics_Cardinality
	Statistics_SubPart
	Statistics_Packed
	Statistics_Nullable
	Statistics_IndexType
	Statistics_Comment
	Statistics_IndexComment
	Statistics_IsVisible
	Statistics_Expression

	// Columns for table information_schema.TABLES
	Tables_TableCatalog
	Tables_TableSchema
	Tables_TableName
	Tables_TableType
	Tables_Engine
	Tables_Version
	Tables_RowFormat
	Tables_TableRows
	Tables_AvgRowLength
	Tables_DataLength
	Tables_MaxDataLength
	Tables_IndexLength
	Tables_DataFree
	Tables_AutoIncrement
	Tables_CreateTime
	Tables_UpdateTime
	Tables_CheckTime
	Tables_TableCollation
	Tables_Checksum
	Tables_CreateOptions
	Tables_TableComment

	// Columns for table fake_benchmark.person
	Person_ID
	Person_Name

	// Columns for table fake_benchmark.tag
	Tag_ID
	Tag_Name

	// Columns for table fake_benchmark.pet
	Pet_ID
	Pet_PersonID
	Pet_TagID
	Pet_Species

	// Columns for table fake_benchmark.extensive
	Extensive_ID
	Extensive_Tinyint
	Extensive_TinyintUnsigned
	Extensive_TinyintNull
	Extensive_Smallint
	Extensive_SmallintUnsigned
	Extensive_SmallintNull
	Extensive_Int
	Extensive_IntNull
	Extensive_IntUnsigned
	Extensive_Bigint
	Extensive_BigintNull
	Extensive_BigintUnsigned
	Extensive_Varchar
	Extensive_VarcharNull
	Extensive_Float
	Extensive_FloatNull
	Extensive_Double
	Extensive_DoubleNull
	Extensive_Decimal
	Extensive_DecimalNull
	Extensive_Numeric
	Extensive_NumericNull
	Extensive_CreatedAt
	Extensive_UpdatedAt
	Extensive_Tinyint1
	Extensive_Tinyint1Null
	Extensive_Year
	Extensive_YearNull
	Extensive_Date
	Extensive_DateNull
	Extensive_Time
	Extensive_TimeNull
	Extensive_Datetime
	Extensive_DatetimeNull
	Extensive_Timestamp
	Extensive_TimestampNull
	Extensive_Char
	Extensive_CharNull
	Extensive_Tinytext
	Extensive_TinytextNull
	Extensive_Text
	Extensive_TextNull
	Extensive_Mediumtext
	Extensive_MediumtextNull
	Extensive_Longtext
	Extensive_LongtextNull
	Extensive_Binary
	Extensive_BinaryNull
	Extensive_Varbinary
	Extensive_VarbinaryNull
	Extensive_Tinyblob
	Extensive_TinyblobNull
	Extensive_Blob
	Extensive_BlobNull
	Extensive_Mediumblob
	Extensive_MediumblobNull
	Extensive_Longblob
	Extensive_LongblobNull
)

// ^^ END OF GENERATED BY CODEGEN. ^^
