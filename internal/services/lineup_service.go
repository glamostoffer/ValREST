package services

import (
	repository "ValREST/internal/database"
	"ValREST/internal/models"
)

type LineupService struct {
	repo repository.Lineup
}

func NewLineupService(repo repository.Lineup) *LineupService {
	return &LineupService{repo: repo}
}

func (s *LineupService) Create(userId int, lineup models.Lineup) (int, error) {
	return s.repo.Create(userId, lineup)
}

func (s *LineupService) GetByAgent(agent string) ([]string, error) {
	return s.repo.GetByAgent(agent)
}

func (s *LineupService) DeleteLinup(id int) error {
	return s.repo.DeleteLinup(id)
}

func (s *LineupService) GetByAgentAndMap(agent, mapName string) ([]string, error) {
	return s.repo.GetByAgentAndMap(agent, mapName)
}

func (s *LineupService) GetByAgentMapObjective(agent, mapName, objective string) ([]string, error) {
	return s.repo.GetByAgentMapObjective(agent, mapName, objective)
}

func (s *LineupService) GetByAgentMapObjectiveAbility(agent, mapName, objective, ability string) ([]string, error) {
	return s.repo.GetByAgentMapObjectiveAbility(agent, mapName, objective, ability)
}
