package handler

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/internal"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/errorwrapper"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/ginwrapper"
	"github.com/gin-gonic/gin"
)

type handler struct {
	internal internal.InternalInterface
}

func New(i internal.InternalInterface) HandlerInterface {
	return &handler{
		internal: i,
	}
}

type HandlerInterface interface {
	HandleRegisterUser(c *gin.Context)
	HandleLoginUser(c *gin.Context)
}

func (h *handler) HandleRegisterUser(c *gin.Context) {
	var input model.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		ginwrapper.RespondWithError(c, errorwrapper.ErrValidationNotPassed.WithDetail(map[string]any{"error": err.Error()}))
		return
	}

	h.internal.ProcessRegisterUser(c, input)
}

func (h *handler) HandleLoginUser(c *gin.Context) {
	var input model.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		ginwrapper.RespondWithError(c, errorwrapper.ErrValidationNotPassed.WithDetail(map[string]any{"error": err.Error()}))
		return
	}

	h.internal.ProcessLoginUser(c, input)
}
