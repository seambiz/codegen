package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PetRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* PetRepository interface definition. */
type PetRepository interface {
	Create(data *Pet) error
	Update(data *Pet) error
	Delete(data *Pet) error
	OneByID(id int) (*Pet, error)

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
}
