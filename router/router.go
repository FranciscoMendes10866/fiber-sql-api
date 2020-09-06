package router

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Use(middleware.Logger())
	api.Use(middleware.Compress())
	api.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
}
