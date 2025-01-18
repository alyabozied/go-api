package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) (int, error)
		GetByID(context.Context, int) (User, error)
		GetByEmail(context.Context, string) (User, error)
	}
}

type DBModel struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}
