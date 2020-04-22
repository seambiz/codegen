package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PersonRepository

// GENERATED BY CODEGEN. DO NOT EDIT.
/* PersonRepository interface definition. */
type PersonRepository interface {
	Create(data *Person) error
	Update(data *Person) error
	Delete(data *Person) error
	OneByID(id int) (*Person, error)

	// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^
}