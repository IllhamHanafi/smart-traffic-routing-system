package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterPingRoute registers the ping endpoint to the router
func RegisterPingRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
