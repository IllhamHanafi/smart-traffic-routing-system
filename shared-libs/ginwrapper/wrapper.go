package ginwrapper

import (
	"net/http"

	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/errorwrapper"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, err errorwrapper.ErrorWrapper) {
	c.AbortWithStatusJSON(err.GetHttpStatusCode(), ErrorResponse{
		Message: err.GetMessage(),
		Status:  StatusFailed,
		Errors:  err.GetCode(),
		Data:    nil,
	})
}

func RespondWithSuccess(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Status:  StatusSuccess,
		Data:    data,
	})
}
