package transactionctrl

import (
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/transaction"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/exception"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ControllerHTTP struct {
	svc transaction.Service
}

func New(svc transaction.Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Checkout
// @Description Checkout based on products in cart
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param body body model.CreateTransactionRequest true "Payload Create Transaction Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/transactions/checkout [post]
func (ctrl ControllerHTTP) Checkout(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgutil.HTTPResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "invalid or expired token",
		})
	}

	var req model.CreateTransactionRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	customerIDUUID, err := uuid.Parse(claims.Subject)
	exception.PanicIfNeeded(err)

	req.CustomerID = customerIDUUID

	res, err := ctrl.svc.Checkout(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusCreated,
		Data: res,
	})
}

// @Summary Pay
// @Description Pay transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param body body model.TransactionPayRequest true "Payload Transaction Pay Request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/transactions/pay [post]
func (ctrl ControllerHTTP) Pay(c *fiber.Ctx) error {
	var req model.TransactionPayRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = ctrl.svc.Pay(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
	})
}
