package repository

import (
	"github.com/google/uuid"
)

type ActiveRoutingLogicResult struct {
	ID              uuid.UUID
	AllocationLogic map[string]int
}

type InsertRoutingDecisionLogParams struct {
	OrderID           uuid.UUID
	CourierID         uuid.UUID
	RoutingDecisionID uuid.UUID
	Status            string
	Reason            string
	CreatedBy         uuid.UUID
}
