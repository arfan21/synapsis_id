package customerctrl

import (
	"github.com/arfan21/synapsis_id/internal/customer"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/pkg/exception"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
)

type ControllerHTTP struct {
	svc customer.Service
}

func New(svc customer.Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Register Customer
// @Description Register Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body model.CustomerRegisterRequest true "Payload Customer Register Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/customer/register [post]
func (ctrl ControllerHTTP) Register(c *fiber.Ctx) error {
	var req model.CustomerRegisterRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = ctrl.svc.Register(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusCreated,
	})
}

// @Summary Login Customer
// @Description Login Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body model.CustomerLoginRequest true "Payload Customer Login Request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/customer/login [post]
func (ctrl ControllerHTTP) Login(c *fiber.Ctx) error {
	var req model.CustomerLoginRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	res, err := ctrl.svc.Login(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
		Data: res,
	})
}
