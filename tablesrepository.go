package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TablesRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* TablesRepository interface definition. */
type TablesRepository interface {
	Create(data *Tables) error
	Update(data *Tables) error
	UpdatePartial(data *TablesPartial) error
	Upsert(data []*Tables) error
	Delete(data *Tables) error

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
	QueryBySchema(schema string) ([]*Tables, error)
}
