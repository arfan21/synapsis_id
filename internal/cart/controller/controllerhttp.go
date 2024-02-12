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
