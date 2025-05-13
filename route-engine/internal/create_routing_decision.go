package internal

import (
	"net/http"
	"route-engine/model"
	"route-engine/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) ProcessCreateRoutingDecision(c *gin.Context, input model.CreateRoutingDecisionRequest) {
	isRequestValid, message := s.IsCreateRoutingDecisionRequestValid(c, input)
	if !isRequestValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"status":  "error",
		})
		return
	}

	err := s.repository.CreateActiveRoutingDecision(c.Request.Context(), repository.InsertNewRoutingDecisionParams{
		UserID:          input.UserID,
		AllocationLogic: input.AllocationLogic,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
	}

	// To Do: delete cache on current active routing decision

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  "success",
	})
}

func (s *Service) IsCreateRoutingDecisionRequestValid(c *gin.Context, input model.CreateRoutingDecisionRequest) (bool, string) {
	probabilitySum := 0
	for courierCode, prob := range input.AllocationLogic {
		courier, err := s.repository.GetCourierByCode(c, courierCode)
		if err != nil {
			return false, "courier not found"
		}
		if courier.ID == uuid.Nil {
			return false, "courier not found"
		}
		probabilitySum += prob
	}

	if probabilitySum != 100 {
		return false, "probability sum is not 100"
	}
	return true, ""
}
