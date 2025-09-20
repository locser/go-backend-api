package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Status:  statusCode,
		Message: Msg[statusCode],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(http.StatusOK, BaseResponse{
		Status:  statusCode,
		Message: Msg[statusCode],
		Data:    nil,
	})
}
