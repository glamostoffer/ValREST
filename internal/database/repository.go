package database

import (
	postgres "ValREST/internal/database/postgresql"
	"ValREST/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name, pssword string) (models.User, error)
}

type Lineup interface {
	Create(userId int, lineup models.Lineup) (int, error)
}

type Repository struct {
	Authorization
	Lineup
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Lineup:        postgres.NewLineupPostgres(db),
	}
}

func NewLineup(db *sqlx.DB) *Lineup {
	return nil
}
