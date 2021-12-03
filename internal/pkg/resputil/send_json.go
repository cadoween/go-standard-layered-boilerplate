package resputil

import "github.com/gin-gonic/gin"

type baseResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendJSON(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, &baseResponse{
		Code:    0, // Zero represent success code.
		Message: msg,
		Data:    data,
	})
}
