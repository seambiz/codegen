package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . KeyColumnUsageRepository

// GENERATED BY CODEGEN.
/* KeyColumnUsageRepository interface definition. */
type KeyColumnUsageRepository interface {
	// ^^ END OF GENERATED BY CODEGEN. ^^
	QueryBySchemaAndRefSchemaAndTable(ctx *Context, schema, refschema, table string) ([]*KeyColumnUsage, error)
}
