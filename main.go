package main

import (
	mid "auth-service/pkg/middleware"

	"auth-service/pkg/routes"
	"auth-service/platform/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	mid.FiberMiddleware(app)
	database.Init()
	routes.PublicRoutes(app)
	app.Listen(":3000")
}
