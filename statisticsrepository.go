package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StatisticsRepository

// GENERATED BY CODEGEN.
/* STATISTICSRepository interface definition. */
type STATISTICSRepository interface {
		
		
		
		

	// ^^ END OF GENERATED BY CODEGEN. ^^
	IndexNameBySchemaAndTable(ctx *Context, schema, table string) ([]*Statistics, error)
	QueryBySchemaAndTableAndIndex(ctx *Context, schema, table, index string) ([]*Statistics, error)
}
