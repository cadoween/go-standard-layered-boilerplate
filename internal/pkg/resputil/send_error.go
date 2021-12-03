package resputil

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/pkg/errorsutil"
)

type errorResponse struct {
	Code    api.ErrorCode `json:"code"`
	Message string        `json:"message"`
	Errors  interface{}   `json:"errors"`
}

func SendError(c *gin.Context, err error) {
	var httpStatus int
	var resp errorResponse
	var internalErr *errorsutil.InternalError

	if errors.As(err, &internalErr) {
		switch internalErr.Code() {
		case api.ErrCodeNotFound:
			httpStatus = http.StatusNotFound
		case api.ErrCodeBadRequest:
			httpStatus = http.StatusBadRequest
		case api.ErrCodeInternalUnknown:
			fallthrough
		default:
			httpStatus = http.StatusInternalServerError
		}

		resp.Code = internalErr.Code()
		resp.Message = internalErr.Error()

	} else {
		httpStatus = http.StatusInternalServerError

		resp.Code = api.ErrCodeInternalUnknown
		resp.Message = "Internal server error"
	}

	c.JSON(httpStatus, &resp)
}
