package errorsutil

import (
	"fmt"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api"
)

type InternalError struct {
	err  error
	msg  string
	code api.ErrorCode
}

func Wrapf(err error, msg string, code api.ErrorCode) error {
	return &InternalError{
		err:  err,
		msg:  msg,
		code: code,
	}
}

func (e InternalError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.err)
	}

	return e.msg
}

func (e *InternalError) Code() api.ErrorCode {
	return e.code
}
