package service

import (
	"github.com/Egor-Tihonov/SandMailLogic/internal/repository"
)

type Service struct {
	repo *repository.PostgresDB
}

func New(rps *repository.PostgresDB) *Service {
	return &Service{repo: rps}
}
