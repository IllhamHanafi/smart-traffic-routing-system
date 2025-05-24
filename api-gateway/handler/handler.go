package handler

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/internal"
)

type handler struct {
	internal internal.InternalInterface
}

func New(i internal.InternalInterface) HandlerInterface {
	return &handler{
		internal: i,
	}
}

type HandlerInterface interface {
}
