package repository

import "fmt"

// ValidationError input param validation error
type ValidationError struct {
	cause error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed. %+v", e.cause)
}

func (e *ValidationError) UnWrap() error {
	return e.cause
}

// InfrastructureError infrastructure error
type InfrastructureError struct {
	cause error
}

func (e *InfrastructureError) Error() string {
	return fmt.Sprintf("infrastructure error. %+v", e.cause)
}

func (e *InfrastructureError) UnWrap() error {
	return e.cause
}

type StoreOperationCode int

const (
	_ StoreOperationCode = iota
	StoreOperationCodeUnkownError
	StoreOperationCodeNotFound
)

// StoreOperationError datastore operation error
type StoreOperationError struct {
	cause error
	Code  StoreOperationCode
}

func (e *StoreOperationError) Error() string {
	return fmt.Sprintf("store operation error. code:[%v] %+v", e.Code, e.cause)
}

func (e *StoreOperationError) UnWrap() error {
	return e.cause
}

func NewStoreOperationError(code StoreOperationCode, err error) error {
	return &StoreOperationError{Code: code, cause: err}
}
