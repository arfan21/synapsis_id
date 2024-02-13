package constant

import "errors"

const (
	ErrSQLUniqueViolation = "23505"
)

var (
	ErrEmailAlreadyRegistered    = errors.New("email already registered")
	ErrEmailOrPasswordInvalid    = errors.New("email or password invalid")
	ErrUnauthorizedAccess        = errors.New("unauthorized access")
	ErrCategoryNotFound          = errors.New("category not found")
	ErrStringNotDecimal          = errors.New("string value is not decimal")
	ErrInvalidUUID               = errors.New("invalid uuid length or format")
	ErrProductNotFound           = errors.New("product not found")
	ErrProductAlreadyAddedToCart = errors.New("product already added into cart")
	ErrCannotAddOwnProductToCart = errors.New("cannot add own product to cart")
	ErrPaymentMethodNotFound     = errors.New("payment method not found")
	ErrNoProductInCart           = errors.New("no product in cart")
	ErrProductStokNotEnough      = errors.New("product stok not enough")
	ErrProductNotFoundOrStok     = errors.New("product not found or stok not enough")
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
