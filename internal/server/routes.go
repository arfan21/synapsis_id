package server

import (
	customerctrl "github.com/arfan21/synapsis_id/internal/customer/controller"
	customerrepo "github.com/arfan21/synapsis_id/internal/customer/repository"
	customersvc "github.com/arfan21/synapsis_id/internal/customer/service"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Routes() {

	api := s.app.Group("/api")
	api.Get("/health-check", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	customerRepo := customerrepo.New(s.db)
	customerSvc := customersvc.New(customerRepo)
	customerCtrl := customerctrl.New(customerSvc)

	s.RoutesCustomer(api, customerCtrl)

}

func (s Server) RoutesCustomer(route fiber.Router, ctrl *customerctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	v1.Post("/customer/register", ctrl.Register)
}
