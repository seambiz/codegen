package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TablesRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* TablesRepository interface definition. */
type TablesRepository interface {
	Create(ctx *BaseContext, data *Tables) error
	Update(ctx *BaseContext, data *Tables) error
	UpdatePartial(ctx *BaseContext, data *TablesPartial) error
	Upsert(ctx *BaseContext, data []*Tables) error
	Delete(ctx *BaseContext, data *Tables) error

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
	QueryBySchema(ctx *BaseContext, schema string) ([]*Tables, error)
}
