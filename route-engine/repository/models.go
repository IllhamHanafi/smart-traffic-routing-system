package repository

import (
	"github.com/google/uuid"
)

type ActiveRoutingLogicResult struct {
	ID              uuid.UUID
	AllocationLogic map[string]int
}
