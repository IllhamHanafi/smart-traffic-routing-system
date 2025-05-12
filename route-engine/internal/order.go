package internal

import (
	"net/http"
	"route-engine/model"
	"route-engine/repository"

	"github.com/gin-gonic/gin"
)

func (s *Service) ProcessOrder(c *gin.Context, order model.Order) {
	// to do: cache for faster process
	res, err := s.repository.GetActiveRoutingLogic(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	allocatedCourier := getCourierFromProbability(res.AllocationLogic)
	// to do: cache
	cour, err := s.repository.GetCourierByCode(c.Request.Context(), allocatedCourier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	// to do: make proses async using message queue
	err = s.repository.InsertRoutingDecisionLog(c.Request.Context(), repository.InsertRoutingDecisionLogParams{
		OrderID:           order.ID,
		CourierID:         cour.ID,
		RoutingDecisionID: res.ID,
		Status:            "success",
		Reason:            "default allocation",
		CreatedBy:         model.SYSTEM_UUID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "order created !",
		"order":            order.ID,
		"status":           "success",
		"allocatedCourier": allocatedCourier,
	})
}
