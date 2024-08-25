package serviceerr

import (
	"errors"
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ErrorServiceType типы ошибок сервисов
type ErrorServiceType int

const (
	InvalidInputDataErrorType ErrorServiceType = 0
	RuntimeErrorType          ErrorServiceType = 1
	NotFoundErrorType         ErrorServiceType = 2
	ConflictErrorType         ErrorServiceType = 3
	FailPreconditionErrorType ErrorServiceType = 4
	PermissionDeniedType      ErrorServiceType = 5
)

type FieldViolation struct {
	Field    string
	ErrorMsg string
}

type ErrorInfo struct {
	Description     string
	FieldViolations []FieldViolation
}

// ErrorService ошибка сервиса
type ErrorService struct {
	Type    ErrorServiceType
	Err     error
	ErrInfo ErrorInfo
}

func (r *ErrorService) Error() string {
	if r == nil || r.Err == nil {
		return ""
	}
	return r.Err.Error()
}

func (r *ErrorService) Unwrap() error {
	return r.Err
}

// NotFoundf создание ошибки - не найдено
func NotFoundf(description string, a ...any) error {
	return &ErrorService{
		Type: NotFoundErrorType,
		Err:  fmt.Errorf(description, a...),
		ErrInfo: ErrorInfo{
			Description: fmt.Sprintf(description, a...),
		},
	}
}

// Conflictf создание ошибки конфликта
func Conflictf(description string, a ...any) error {
	return &ErrorService{
		Type: ConflictErrorType,
		Err:  fmt.Errorf(description, a...),
		ErrInfo: ErrorInfo{
			Description: fmt.Sprintf(description, a...),
		},
	}
}

func PermissionDeniedf(description string, a ...any) error {
	return &ErrorService{
		Type: PermissionDeniedType,
		Err:  fmt.Errorf(description, a...),
		ErrInfo: ErrorInfo{
			Description: fmt.Sprintf(description, a...),
		},
	}
}

// FailPreconditionf создание ошибки fail prediction
func FailPreconditionf(description string, a ...any) error {
	return &ErrorService{
		Type: FailPreconditionErrorType,
		Err:  fmt.Errorf(description, a...),
		ErrInfo: ErrorInfo{
			Description: fmt.Sprintf(description, a...),
		},
	}
}

func InvalidInputf(description string, a ...any) error {
	return &ErrorService{
		Type: InvalidInputDataErrorType,
		Err:  fmt.Errorf(description, a...),
		ErrInfo: ErrorInfo{
			Description: fmt.Sprintf(description, a...),
		},
	}
}

// InvalidInput создание ошибки входных данных
func InvalidInput(trans ut.Translator, err error, formName string) error {
	fields := make([]FieldViolation, 0)
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, validationErr := range validationErrors {
			fields = append(fields, FieldViolation{
				Field:    validationErr.Field(),
				ErrorMsg: validationErr.Translate(trans),
			})
		}
	}
	return &ErrorService{
		Type: InvalidInputDataErrorType,
		Err:  fmt.Errorf(formName+": %w", err),
		ErrInfo: ErrorInfo{
			Description:     fmt.Sprintf("Form data validation error %s", formName),
			FieldViolations: fields,
		},
	}
}

// MakeErr создание
func MakeErr(err error, method string) error {
	return &ErrorService{
		Type: RuntimeErrorType,
		Err:  fmt.Errorf(method+": %w", err),
		ErrInfo: ErrorInfo{
			Description: "Runtime error",
		},
	}
}
