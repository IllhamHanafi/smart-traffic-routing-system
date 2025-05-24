package internal

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository"
)

type InternalInterface interface {
	Close()
}

type Service struct {
	config     config.Config
	repository repository.Repository
}

func New() InternalInterface {
	cfg := config.GetConfig()
	repo := repository.New(cfg.Database)
	return &Service{
		config:     cfg,
		repository: repo,
	}
}

func (s *Service) Close() {
	s.repository.Close()
}
