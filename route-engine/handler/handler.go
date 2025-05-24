package handler

import (
	"net/http"

	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/internal"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/model"

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
	HandleGetActiveRoutingDecision(c *gin.Context)
	HandleCreateActiveRoutingDecision(c *gin.Context)
	HandleGetRoutingDecisionLogs(c *gin.Context)
}

func (h *handler) HandleCreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.internal.ProcessOrder(c, order)
}

func (h *handler) HandleGetActiveRoutingDecision(c *gin.Context) {
	h.internal.ProcessGetActiveRoutingDecision(c)
}

func (h *handler) HandleCreateActiveRoutingDecision(c *gin.Context) {
	var request model.CreateRoutingDecisionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.internal.ProcessCreateRoutingDecision(c, request)
}

func (h *handler) HandleGetRoutingDecisionLogs(c *gin.Context) {
	var request model.GetRoutingDecisionLogsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.internal.ProcessGetRoutingDecisionLogs(c, request)
}
