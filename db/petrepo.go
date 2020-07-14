package db

import (
	"database/sql"

	"bitbucket.org/codegen"
)

// GENERATED BY CODEGEN. DO NOT EDIT.
// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^

/* PetRepo implemente PetRepository interface definition. */
type PetRepo struct {
	conn *sql.DB
}

func NewPetRepo(conn *sql.DB) *PetRepo {
	return &PetRepo{
		conn: conn,
	}
}

func (r PetRepo) Create(data *codegen.Pet) error {
	panic("not implemented")
}

func (r PetRepo) Update(data *codegen.Pet) error {
	panic("not implemented")
}

func (r PetRepo) UpdatePartial(data *codegen.PetPartial) error {
	panic("not implemented")
}

func (r PetRepo) Delete(data *codegen.Pet) error {
	panic("not implemented")
}

func (r PetRepo) Upsert(data []*codegen.Pet) error {
	panic("not implemented")
}

func (r PetRepo) OneByID(id int) (*codegen.Pet, error) {
	panic("not implemented")
}
