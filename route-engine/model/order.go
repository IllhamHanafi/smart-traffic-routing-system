package model

import "github.com/google/uuid"

var (
	SYSTEM_UUID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")
)

type Order struct {
	ID uuid.UUID `json:"id"`
}

type CreateRoutingDecisionRequest struct {
	UserID          uuid.UUID      `json:"user_id"`
	AllocationLogic map[string]int `json:"allocation_logic"`
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
