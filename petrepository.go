package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PetRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* PetRepository interface definition. */
type PetRepository interface {
	Create(ctx *BaseContext, data *Pet) error
	Update(ctx *BaseContext, data *Pet) error
	UpdatePartial(ctx *BaseContext, data *PetPartial) error
	Upsert(ctx *BaseContext, data []*Pet) error
	Delete(ctx *BaseContext, data *Pet) error
	OneByID(ctx *BaseContext, id int) (*Pet, error)

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
}
