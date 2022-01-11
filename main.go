package main

import (
	mid "auth-service/pkg/middleware"
	"auth-service/pkg/routes"
	"auth-service/platform/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	mid.FiberMiddleware(app)
	routes.PublicRoutes(app)
	database.Init()
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
