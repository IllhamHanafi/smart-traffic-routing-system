package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) ProcessGetActiveRoutingDecision(c *gin.Context) {
	// to do: cache for faster process
	res, err := s.repository.GetActiveRoutingLogic(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "success",
		"status":           "success",
		"allocation_logic": res.AllocationLogic,
	})
}
