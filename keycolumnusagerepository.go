package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . KeyColumnUsageRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* KeyColumnUsageRepository interface definition. */
type KeyColumnUsageRepository interface {
	Create(data *KeyColumnUsage) error
	Update(data *KeyColumnUsage) error
	UpdatePartial(data *KeyColumnUsagePartial) error
	Upsert(data []*KeyColumnUsage) error
	Delete(data *KeyColumnUsage) error

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
	QueryBySchemaAndRefSchemaAndTable(schema, refschema, table string) ([]*KeyColumnUsage, error)
}
