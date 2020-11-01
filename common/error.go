package common

import "fmt"

// AppError Applcaition code and message
type AppError struct {
	Code    uint32
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%d:%s: AppError", e.Code, e.Message)
}

func AppErrorCode(code uint32) *AppError {
	return &AppError{code, GetMessage(code)}
}
