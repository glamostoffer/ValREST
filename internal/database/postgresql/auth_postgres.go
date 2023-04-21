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

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password, role) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Password, "default")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(name, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id from %s WHERE name=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, name, password)

	return user, err
}
