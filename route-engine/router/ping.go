package router

import (
	"net/http"
	"route-engine/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoute registers the ping endpoint to the router
func RegisterRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/order", handler.HandleCreateOrder)
}
