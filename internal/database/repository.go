package database

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Lineup interface {
}

type Repository struct {
	Authorization
	Lineup
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
