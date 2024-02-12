package productctrl

import (
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/product"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/exception"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
)

type ControllerHTTP struct {
	svc product.Service
}

func New(svc product.Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param body body model.ProductCreateRequest true "Payload Create Product Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/products [post]
func (ctrl ControllerHTTP) Create(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkgutil.HTTPResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "invalid or expired token",
		})
	}

	var req model.ProductCreateRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	req.CustomerID = claims.Subject

	err = ctrl.svc.Create(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusCreated,
	})
}

// @Summary Get Product Categories
// @Description Get Product Categories
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} pkgutil.HTTPResponse{data=[]model.GetCategoryResponse}
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/products/categories [get]
func (ctrl ControllerHTTP) GetCategories(c *fiber.Ctx) error {
	res, err := ctrl.svc.GetCategories(c.UserContext())
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
		Data: res,
	})
}
