package paymentctrl

import (
	_ "github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/payment"
	"github.com/arfan21/synapsis_id/pkg/exception"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
)

type ControllerHTTP struct {
	svc payment.Service
}

func New(svc payment.Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Get Payment Methods
// @Description Get Payment Methods
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} pkgutil.HTTPResponse{data=[]model.GetPayemntMethodResponse}
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/payments/methods [get]
func (ctrl ControllerHTTP) GetPaymentMethods(c *fiber.Ctx) error {
	res, err := ctrl.svc.GetPaymentMethods(c.UserContext())
	exception.PanicIfNeeded(err)

	return c.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code: fiber.StatusOK,
		Data: res,
	})
}
