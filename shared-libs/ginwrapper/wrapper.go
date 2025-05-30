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
		Detail:  err.GetDetail(),
	})
}

func RespondWithSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: MessageSuccess,
		Status:  StatusSuccess,
		Data:    data,
	})
}
