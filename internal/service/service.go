package service

import "github.com/Way-Flare/dagestan-backend/internal/database/psql"

type Authorization interface {
}

type Place interface {
}

type Service struct {
	Authorization
	Place
}

func NewService(repos *psql.Repository) *Service {
	return &Service{}
}
