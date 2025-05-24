package router

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoute registers the ping endpoint to the router
func RegisterRoute(r *gin.Engine, h handler.HandlerInterface) {
	r.POST("/route/order", h.HandleCreateOrder)
	r.GET("/routing-decision/active", h.HandleGetActiveRoutingDecision)
	r.POST("/routing-decision", h.HandleCreateActiveRoutingDecision)
	r.GET("/routing-decision/logs", h.HandleGetRoutingDecisionLogs)
}
