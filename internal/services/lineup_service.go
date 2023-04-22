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
