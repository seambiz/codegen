package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ColumnsRepository

// GENERATED BY CODEGEN.
/* COLUMNSRepository interface definition. */
type COLUMNSRepository interface {
		
		
		
		

	// ^^ END OF GENERATED BY CODEGEN. ^^
	QueryBySchemaAndTable(ctx *Context, schema, table string) ([]*Columns, error)
}
