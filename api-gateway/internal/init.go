package internal

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository"
	"github.com/gin-gonic/gin"
)

type InternalInterface interface {
	ProcessRegisterUser(c *gin.Context, input model.RegisterUserInput)
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
