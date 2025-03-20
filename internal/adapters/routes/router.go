package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/somphonee/go-fiber-hex/internal/adapters/handlers"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// ตั้งค่า Middleware
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User routes
	users := v1.Group("/users")
	users.Post("/", userHandler.Create)
	users.Get("/:id", userHandler.GetByID)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}