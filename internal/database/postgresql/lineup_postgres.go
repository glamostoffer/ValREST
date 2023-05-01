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
	lineup.Author = userId
	query := fmt.Sprintf("INSERT INTO %s (agent, mapname, objective, ability, source, author) values ($1, $2, $3, $4, $5, $6) RETURNING id", lineupsTable)
	row := r.db.QueryRow(query, lineup.Agent, lineup.MapName, lineup.Objective, lineup.Ability, lineup.Source, lineup.Author)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("error during insert lineup into db: %s", err.Error())
		return 0, err
	}

	return id, nil

}

func (r *LineupPostgres) GetByAgent(agent string) ([]string, error) {
	var sources []string
	query := fmt.Sprintf("SELECT source from %s where agent=$1", lineupsTable)
	err := r.db.Select(&sources, query, agent)

	return sources, err
}

func (r *LineupPostgres) DeleteLinup(id int) error {
	query := fmt.Sprintf("DELETE from %s WHERE id=$1", lineupsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *LineupPostgres) GetByAgentAndMap(agent, mapName string) ([]string, error) {
	var sources []string
	query := fmt.Sprintf("SELECT source from %s where agent=$1 and mapname=$2", lineupsTable)
	err := r.db.Select(&sources, query, agent, mapName)

	return sources, err
}

func (r *LineupPostgres) GetByAgentMapObjective(agent, mapName, objective string) ([]string, error) {
	var sources []string
	query := fmt.Sprintf("SELECT source from %s where agent=$1 and mapname=$2 and objective=$3", lineupsTable)
	err := r.db.Select(&sources, query, agent, mapName, objective)

	return sources, err
}

func (r *LineupPostgres) GetByAgentMapObjectiveAbility(agent, mapName, objective, ability string) ([]string, error) {
	var sources []string
	query := fmt.Sprintf("SELECT source from %s where agent=$1 and mapname=$2 and objective=$3 and ability=$4", lineupsTable)
	err := r.db.Select(&sources, query, agent, mapName, objective, ability)

	return sources, err
}
