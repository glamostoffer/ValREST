package services

import (
	repo "ValREST/internal/database"
	"ValREST/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
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
