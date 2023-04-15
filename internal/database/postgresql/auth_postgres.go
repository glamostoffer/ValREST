package postgrsql

import (
	"ValREST/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password, role) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := s.db.QueryRow(query, user.Name, user.Password, "default")
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
