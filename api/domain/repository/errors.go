package repository

import "fmt"

type ErrorCode int

const (
	_ ErrorCode = iota
	ErrUnknown
	ErrNotFound
)

// OperationError datastore operation error
type OperationError struct {
	cause error
	Code  ErrorCode
}

func (e OperationError) Error() string {
	return fmt.Sprintf("operation error. code:[%v] %+v", e.Code, e.cause)
}

func (e OperationError) UnWrap() error {
	return e.cause
}

func NewOperationError(code ErrorCode, err error) error {
	return &OperationError{Code: code, cause: err}
}
