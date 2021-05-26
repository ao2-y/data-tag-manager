package usecase

import "fmt"

type ValidationType int

const (
	// ValidationTypeLength 文字列長/数値長チェック
	ValidationTypeLength ValidationType = iota
	// ValidationTypeType 型チェック
	ValidationTypeType
	// ValidationTypeDuplicated 重複チェック
	ValidationTypeDuplicated
	// ValidationTypeExist 存在チェック
	ValidationTypeExist
	// ValidationTypeUsed 他のリソースに依存されている
	ValidationTypeUsed
)

type ValidationError struct {
	Field          string
	ValidationType ValidationType
	Value          interface{}
	Cause          error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed, type=%v ,field=%s ,value=%v ,cause=%v", e.ValidationType, e.Field, e.Value, e.Cause)
}

func (e ValidationError) UnWrap() error {
	return e.Cause
}

func NewValidationError(vType ValidationType, field string, value interface{}, cause error) error {
	return &ValidationError{
		Field:          field,
		ValidationType: vType,
		Value:          value,
		Cause:          cause,
	}
}

type InternalServerError struct {
	Message string
	Cause   error
}

func (e InternalServerError) Error() string {
	return fmt.Sprintf("%s:%v", e.Message, e.Cause)
}

func (e InternalServerError) UnWrap() error {
	return e.Cause
}

func NewInternalServerError(msg string, cause error) error {
	return &InternalServerError{
		Message: msg,
		Cause:   cause,
	}
}

type ResourceNotFoundError struct {
	ResourceType string
	Cause        error
}

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("resource not found.[%s]", e.ResourceType)
}

func (e ResourceNotFoundError) UnWrap() error {
	return e.Cause
}

func NewResourceNorFoundError(resource string) error {
	return &ResourceNotFoundError{
		ResourceType: resource,
	}
}
