package internal

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/repository"

	"github.com/gin-gonic/gin"
)

type InternalInterface interface {
	ProcessGetActiveRoutingDecision(c *gin.Context)
	ProcessOrder(c *gin.Context, order model.Order)
	ProcessCreateRoutingDecision(c *gin.Context, order model.CreateRoutingDecisionRequest)
	ProcessGetRoutingDecisionLogs(c *gin.Context, input model.GetRoutingDecisionLogsRequest)
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
