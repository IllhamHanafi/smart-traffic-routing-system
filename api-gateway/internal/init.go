package internal

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/jwt"
	"github.com/gin-gonic/gin"
)

type InternalInterface interface {
	ProcessRegisterUser(c *gin.Context, input model.RegisterUserInput)
	ProcessLoginUser(c *gin.Context, input model.LoginUserInput)
	Close()
}

type Service struct {
	config     config.Config
	repository repository.Repository
	jwt        jwt.JWTInterface
}

func New() InternalInterface {
	cfg := config.GetConfig()

	repo := repository.New(cfg.Database)

	jwt, err := jwt.New(cfg.JWT)
	if err != nil {
		panic(err)
	}
	return &Service{
		config:     cfg,
		repository: repo,
		jwt:        jwt,
	}
}

func (s *Service) Close() {
	s.repository.Close()
}
