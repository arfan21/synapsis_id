package server

import "github.com/gofiber/fiber/v2"

func (s *Server) Routes() {

	api := s.app.Group("/api")
	api.Get("/health-check", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	// v1 := api.Group("/v1")

}
