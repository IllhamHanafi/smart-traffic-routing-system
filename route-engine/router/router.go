package router

import (
	"route-engine/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoute registers the ping endpoint to the router
func RegisterRoute(r *gin.Engine) {
	r.POST("/route/order", handler.HandleCreateOrder)
}
