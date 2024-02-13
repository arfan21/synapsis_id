package server

import (
	cartctrl "github.com/arfan21/synapsis_id/internal/cart/controller"
	cartrepo "github.com/arfan21/synapsis_id/internal/cart/repository"
	cartsvc "github.com/arfan21/synapsis_id/internal/cart/service"
	customerctrl "github.com/arfan21/synapsis_id/internal/customer/controller"
	customerrepo "github.com/arfan21/synapsis_id/internal/customer/repository"
	customersvc "github.com/arfan21/synapsis_id/internal/customer/service"
	"github.com/arfan21/synapsis_id/internal/middleware"
	paymentctrl "github.com/arfan21/synapsis_id/internal/payment/controller"
	paymentrepo "github.com/arfan21/synapsis_id/internal/payment/repository"
	paymentsvc "github.com/arfan21/synapsis_id/internal/payment/service"
	productctrl "github.com/arfan21/synapsis_id/internal/product/controller"
	productrepo "github.com/arfan21/synapsis_id/internal/product/repository"
	productsvc "github.com/arfan21/synapsis_id/internal/product/service"
	transactionctrl "github.com/arfan21/synapsis_id/internal/transaction/controller"
	transactionrepo "github.com/arfan21/synapsis_id/internal/transaction/repository"
	transactionsvc "github.com/arfan21/synapsis_id/internal/transaction/service"
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

	paymentRepo := paymentrepo.New(s.db)
	paymentSvc := paymentsvc.New(paymentRepo)
	paymentCtrl := paymentctrl.New(paymentSvc)

	transactionRepo := transactionrepo.New(s.db)
	transactionSvc := transactionsvc.New(transactionRepo, cartSvc, paymentSvc, productSvc)
	transactionCtrl := transactionctrl.New(transactionSvc)

	s.RoutesCustomer(api, customerCtrl)
	s.RoutesProduct(api, productCtrl)
	s.RoutesCart(api, cartCtrl)
	s.RoutesPayment(api, paymentCtrl)
	s.RoutesTransaction(api, transactionCtrl)
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
	cartsV1.Delete("", middleware.JWTAuth, ctrl.Delete)
}

func (s Server) RoutesPayment(route fiber.Router, ctrl *paymentctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	paymentsV1 := v1.Group("/payments")
	paymentsV1.Get("/methods", ctrl.GetPaymentMethods)
}

func (s Server) RoutesTransaction(route fiber.Router, ctrl *transactionctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	transactionsV1 := v1.Group("/transactions")
	transactionsV1.Post("/checkout", middleware.JWTAuth, ctrl.Checkout)
	transactionsV1.Post("/pay", middleware.JWTAuth, ctrl.Pay)
	transactionsV1.Get("", middleware.JWTAuth, ctrl.GetByCustomerID)
}
