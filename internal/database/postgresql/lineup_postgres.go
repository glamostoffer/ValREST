package postgrsql

import (
	"ValREST/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type LineupPostgres struct {
	db *sqlx.DB
}

func NewLineupPostgres(db *sqlx.DB) *LineupPostgres {
	return &LineupPostgres{db: db}
}

func (r *LineupPostgres) Create(userId int, lineup models.Lineup) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (agent, mapname, objective, ability, source, author) values ($1, $2, $3, $4, $5, $6) RETURNING id", lineupsTable)
	row := r.db.QueryRow(query, lineup.Agent, lineup.MapName, lineup.Objective, lineup.Ability, lineup.Source, lineup.Author)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("error during insert lineup into db: %s", err.Error())
		return 0, err
	}

	return id, nil

}
