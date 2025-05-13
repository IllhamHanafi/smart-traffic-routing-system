package internal

import (
	"route-engine/config"
	"route-engine/model"
	"route-engine/repository"

	"github.com/gin-gonic/gin"
)

type InternalInterface interface {
	ProcessGetActiveRoutingDecision(c *gin.Context)
	ProcessOrder(c *gin.Context, order model.Order)
	ProcessCreateRoutingDecision(c *gin.Context, order model.CreateRoutingDecisionRequest)
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
