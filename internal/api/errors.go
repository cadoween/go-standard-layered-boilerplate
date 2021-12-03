package api

type ErrorCode uint

const (
	ErrCodeInternalUnknown ErrorCode = iota
	ErrCodeInternalDatabase
	ErrCodeNotFound
	ErrCodeBadRequest
	ErrCodeFailedValidation
)
