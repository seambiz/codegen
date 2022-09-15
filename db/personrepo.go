package db

import (
	"database/sql"

	"bitbucket.org/codegen"
	"bitbucket.org/codegen/db/mysql"
)

// GENERATED BY CODEGEN.

/* PersonRepo implements PersonRepository interface definition. */
type PersonRepo struct {
	conn *sql.DB
}

func NewPersonRepo(conn *sql.DB) codegen.PersonRepository {
	return &PersonRepo{
		conn: conn,
	}
}

func (r PersonRepo) Create(ctx *codegen.Context, data *codegen.Person) error {
	return mysql.NewPersonStore(ctx, r.conn).Insert(data)
}

func (r PersonRepo) Update(ctx *codegen.Context, data *codegen.Person) error {
	_, err := mysql.NewPersonStore(ctx, r.conn).Update(data)
	return err
}

func (r PersonRepo) UpdatePartial(ctx *codegen.Context, data *codegen.PersonPartial) error {
	_, err := mysql.NewPersonStore(ctx, r.conn).SetBits(&data.Touched).Update(&data.Person)
	return err
}

func (r PersonRepo) Delete(ctx *codegen.Context, data *codegen.Person) error {
	_, err := mysql.NewPersonStore(ctx, r.conn).Delete(data)
	return err
}

func (r PersonRepo) OneByID(ctx *codegen.Context, id int) (*codegen.Person, error) {
	return mysql.NewPersonStore(ctx, r.conn).OneByID(id)
}

// ^^ END OF GENERATED BY CODEGEN. ^^
