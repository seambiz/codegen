package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PersonRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* PersonRepository interface definition. */
type PersonRepository interface {
	Create(ctx *BaseContext, data *Person) error
	Update(ctx *BaseContext, data *Person) error
	UpdatePartial(ctx *BaseContext, data *PersonPartial) error
	Upsert(ctx *BaseContext, data []*Person) error
	Delete(ctx *BaseContext, data *Person) error
	OneByID(ctx *BaseContext, id int) (*Person, error)

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
}
