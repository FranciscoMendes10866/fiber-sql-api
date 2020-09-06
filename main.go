package main

import (
	"github.com/fiber-sqlx/router"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())
	router.SetupRoutes(app)
	app.Listen(8899)
}
