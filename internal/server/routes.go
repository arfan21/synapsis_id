package server

import (
	cartctrl "github.com/arfan21/synapsis_id/internal/cart/controller"
	cartrepo "github.com/arfan21/synapsis_id/internal/cart/repository"
	cartsvc "github.com/arfan21/synapsis_id/internal/cart/service"
	customerctrl "github.com/arfan21/synapsis_id/internal/customer/controller"
	customerrepo "github.com/arfan21/synapsis_id/internal/customer/repository"
	customersvc "github.com/arfan21/synapsis_id/internal/customer/service"
	"github.com/arfan21/synapsis_id/internal/middleware"
	productctrl "github.com/arfan21/synapsis_id/internal/product/controller"
	productrepo "github.com/arfan21/synapsis_id/internal/product/repository"
	productsvc "github.com/arfan21/synapsis_id/internal/product/service"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Routes() {

	api := s.app.Group("/api")
	api.Get("/health-check", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	customerRepo := customerrepo.New(s.db)
	customerRepoRedis := customerrepo.NewRedis(s.redisClient)
	customerSvc := customersvc.New(customerRepo, customerRepoRedis)
	customerCtrl := customerctrl.New(customerSvc)

	productRepo := productrepo.New(s.db)
	productSvc := productsvc.New(productRepo)
	productCtrl := productctrl.New(productSvc)

	cartRepo := cartrepo.New(s.db)
	cartSvc := cartsvc.New(cartRepo, productSvc)
	cartCtrl := cartctrl.New(cartSvc)

	s.RoutesCustomer(api, customerCtrl)
	s.RoutesProduct(api, productCtrl)
	s.RoutesCart(api, cartCtrl)
}

func (s Server) RoutesCustomer(route fiber.Router, ctrl *customerctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	customersV1 := v1.Group("/customers")
	customersV1.Post("/register", ctrl.Register)
	customersV1.Post("/login", ctrl.Login)
	customersV1.Post("/refresh-token", ctrl.RefreshToken)
}

func (s Server) RoutesProduct(route fiber.Router, ctrl *productctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	productsV1 := v1.Group("/products")
	productsV1.Post("", middleware.JWTAuth, ctrl.Create)
	productsV1.Get("", ctrl.GetProducts)
	productsV1.Get("/categories", ctrl.GetCategories)
}

func (s Server) RoutesCart(route fiber.Router, ctrl *cartctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	cartsV1 := v1.Group("/carts")
	cartsV1.Post("", middleware.JWTAuth, ctrl.Create)
	cartsV1.Get("", middleware.JWTAuth, ctrl.GetByCustomerID)
}
