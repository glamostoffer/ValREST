package services

import (
	repo "ValREST/internal/database/postgrsql"
)

type Authorization interface {
}

type Lineup interface {
}

type Service struct {
	Authorization
	Lineup
}

func NewService(repos *repo.Repository) *Service {
	return &Service{}
}
