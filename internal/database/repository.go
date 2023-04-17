package database

import (
	auth "ValREST/internal/database/postgresql"
	"ValREST/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name, pssword string) (models.User, error)
}

type Lineup interface {
}

type Repository struct {
	Authorization
	Lineup
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPostgres(db),
	}
}

func NewLineup(db *sqlx.DB) *Lineup {
	return nil
}
