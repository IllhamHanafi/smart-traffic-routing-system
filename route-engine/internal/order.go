package internal

import (
	"net/http"
	"route-engine/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) ProcessOrder(c *gin.Context, order model.Order) {
	res, err := s.repository.GetActiveRoutingLogic(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	allocatedCourier := getCourierFromProbability(res.AllocationLogic)

	c.JSON(http.StatusOK, gin.H{
		"message":          "order created !",
		"order":            order.ID,
		"status":           "success",
		"allocatedCourier": allocatedCourier,
	})
}
