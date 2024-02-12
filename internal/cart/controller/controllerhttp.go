package cartctrl

import (
	"github.com/arfan21/synapsis_id/internal/cart"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/exception"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ControllerHTTP struct {
	svc cart.Service
}

func New(svc cart.Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Create Cart
// @Description Create Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param body body model.CreateCartRequest true "Payload Create Cart Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/carts [post]
func (ctrl ControllerHTTP) Create(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgutil.HTTPResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "invalid or expired token",
		})
	}

	var req model.CreateCartRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	customerIDUUID, err := uuid.Parse(claims.Subject)
	exception.PanicIfNeeded(err)

	req.CustomerID = customerIDUUID

	err = ctrl.svc.Create(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusCreated,
	})
}

// @Summary Get Cart By Customer ID
// @Description Get Cart By Customer ID
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} pkgutil.HTTPResponse{data=[]model.GetCartResponse} "Success"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/carts [get]
func (ctrl ControllerHTTP) GetByCustomerID(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgutil.HTTPResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "invalid or expired token",
		})
	}

	results, err := ctrl.svc.GetByCustomerID(c.UserContext(), claims.Subject)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
		Data: results,
	})
}

// @Summary Delete Cart By Product ID
// @Description Delete Cart By Product ID
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param body body model.DeleteCartRequest true "Payload Delete Cart Request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/carts [delete]
func (ctrl ControllerHTTP) Delete(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgutil.HTTPResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "invalid or expired token",
		})
	}

	var req model.DeleteCartRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = ctrl.svc.Delete(c.UserContext(), claims.Subject, req.ProductID.String())
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
	})
}
