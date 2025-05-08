package handler

import (
	"net/http"
	"route-engine/internal"
	"route-engine/model"

	"github.com/gin-gonic/gin"
)

func HandleCreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	internal.ProcessOrder(c, order)
}
