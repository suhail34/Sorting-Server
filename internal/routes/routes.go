package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suhail34/sorting_server/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
  app.Post("/process-single", handlers.ProcessSingle)
  app.Post("/process-concurrent", handlers.ProcessConcurrent)
}
