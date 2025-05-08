package internal

import (
	"net/http"
	"route-engine/model"

	"github.com/gin-gonic/gin"
)

func ProcessOrder(c *gin.Context, order model.Order) {
	c.JSON(http.StatusOK, gin.H{
		"message": "order created !",
		"order":   order.ID,
		"status":  "success",
	})
}
