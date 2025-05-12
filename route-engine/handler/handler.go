package handler

import (
	"net/http"
	"route-engine/internal"
	"route-engine/model"

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
	HandleCreateOrder(c *gin.Context)
}

func (h *handler) HandleCreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.internal.ProcessOrder(c, order)
}
