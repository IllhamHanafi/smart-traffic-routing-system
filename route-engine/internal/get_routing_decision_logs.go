package internal

import (
	"net/http"

	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) ProcessGetRoutingDecisionLogs(c *gin.Context, input model.GetRoutingDecisionLogsRequest) {
	repoInput, err := ConstructGetRoutingDecisionLogsRequest(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to construct request",
			"status":  "failed",
			"error":   err.Error(),
		})
		return
	}
	res, err := s.repository.GetRoutingDecisionLogs(c, repoInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get routing decision logs",
			"status":  "failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  "success",
		"data":    res,
	})
}

func ConstructGetRoutingDecisionLogsRequest(r model.GetRoutingDecisionLogsRequest) (repository.GetRoutingDecisionLogsRequest, error) {
	res := repository.GetRoutingDecisionLogsRequest{}
	if r.OrderID != nil {
		orderID, err := uuid.Parse(*r.OrderID)
		if err != nil {
			return res, err
		}
		res.OrderID = &orderID
	}

	if r.CourierID != nil {
		courierID, err := uuid.Parse(*r.CourierID)
		if err != nil {
			return res, err
		}
		res.CourierID = &courierID
	}

	if r.RoutingDecisionID != nil {
		routingDecisionID, err := uuid.Parse(*r.RoutingDecisionID)
		if err != nil {
			return res, err
		}
		res.RoutingDecisionID = &routingDecisionID
	}

	res.Status = r.Status

	if r.Limit == nil || *r.Limit > model.MAX_LIMIT {
		res.Limit = model.DEFAULT_LIMIT
	} else {
		res.Limit = *r.Limit
	}

	if r.Offset == nil {
		res.Offset = model.DEFAULT_OFFSET
	} else {
		res.Offset = *r.Offset
	}

	return res, nil
}
