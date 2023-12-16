package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suhail34/sorting_server/internal/routes"
)

func NewServer() *fiber.App {
  app := fiber.New()
  routes.SetupRoutes(app)
  return app
}
