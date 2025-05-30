package ginwrapper

type SuccessResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
	Detail  any    `json:"detail"`
}

const (
	StatusSuccess = "success"
	StatusFailed  = "failed"

	MessageSuccess = "success"
)
