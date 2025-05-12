package internal

import (
	"route-engine/config"
	"route-engine/model"
	"route-engine/repository"

	"github.com/gin-gonic/gin"
)

type InternalInterface interface {
	Close()
	ProcessOrder(c *gin.Context, order model.Order)
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
