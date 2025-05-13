package model

import (
	"time"

	"github.com/google/uuid"
)

var (
	SYSTEM_UUID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")

	DEFAULT_LIMIT  = int32(10)
	MAX_LIMIT      = int32(100)
	DEFAULT_OFFSET = int32(0)
)

const (
	ROUTING_DECISION_STATUS_ACTIVE   = "active"
	ROUTING_DECISION_STATUS_INACTIVE = "inactive"
)

type Order struct {
	ID uuid.UUID `json:"id"`
}

type CreateRoutingDecisionRequest struct {
	UserID          uuid.UUID      `json:"user_id"`
	AllocationLogic map[string]int `json:"allocation_logic"`
}

type GetRoutingDecisionLogsRequest struct {
	OrderID           *string `form:"order_id" binding:"omitempty,uuid"`
	CourierID         *string `form:"courier_id" binding:"omitempty,uuid"`
	RoutingDecisionID *string `form:"routing_decision_id" binding:"omitempty,uuid"`
	Status            *string `form:"status"`
	Limit             *int32  `form:"limit" binding:"omitempty,min=0"`
	Offset            *int32  `form:"offset" binding:"omitempty,min=0"`
}

type JSONResponse struct {
	RequestID uuid.UUID `json:"request_id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
}

type SuccessCreateOrder struct {
	JSONResponse
	OrderID uuid.UUID `json:"order_id"`
}

type Courier struct {
	ID   uuid.UUID
	Name string
	Code string
}

type RoutingDecisionLog struct {
	ID                uuid.UUID `json:"id"`
	OrderID           uuid.UUID `json:"order_id"`
	CourierID         uuid.UUID `json:"courier_id"`
	RoutingDecisionID uuid.UUID `json:"routing_decision_id"`
	Status            string    `json:"status"`
	Reason            string    `json:"reason"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         uuid.UUID `json:"created_by"`
}
