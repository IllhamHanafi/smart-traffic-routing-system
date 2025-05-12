package model

import "github.com/google/uuid"

type Order struct {
	ID uuid.UUID `json:"id"`
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
