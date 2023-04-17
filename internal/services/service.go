package services

import (
	repo "ValREST/internal/database"
	"ValREST/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(name, password string) (string, error)
}

type Lineup interface {
}

type Service struct {
	Authorization
	Lineup
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
