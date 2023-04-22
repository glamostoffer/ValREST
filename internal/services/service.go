package services

import (
	repo "ValREST/internal/database"
	"ValREST/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Lineup interface {
	Create(userId int, lineup models.Lineup) (int, error)
	GetByAgent(agent string) ([]string, error)
	DeleteLinup(id int) error
}

type Service struct {
	Authorization
	Lineup
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Lineup:        NewLineupService(repos.Lineup),
	}
}
