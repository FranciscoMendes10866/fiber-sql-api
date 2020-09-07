package router

import (
	"github.com/fiber-sqlx/handler"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Use(middleware.Logger())
	api.Use(middleware.Compress())
	auth := api.Group("/auth")
	auth.Post("/", handler.CreateUser)
}
