package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TagRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* TagRepository interface definition. */
type TagRepository interface {
	Create(data *Tag) error
	Update(data *Tag) error
	Delete(data *Tag) error
	OneByID(id int) (*Tag, error)

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
}