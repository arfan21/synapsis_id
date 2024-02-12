package constant

import "errors"

const (
	ErrSQLUniqueViolation = "23505"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrEmailOrPasswordInvalid = errors.New("email or password invalid")
	ErrUnauthorizedAccess     = errors.New("unauthorized access")
	ErrCategoryNotFound       = errors.New("category not found")
	ErrStringNotDecimal       = errors.New("string value is not decimal")
	ErrInvalidUUID            = errors.New("invalid uuid length or format")
)

type ErrNotFound struct {
	Message string
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

type ErrUnauthorized struct {
	Message string
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}

type ErrValidation struct {
	Message string
}

func (e *ErrValidation) Error() string {
	return e.Message
}
