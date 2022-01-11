package main

import (
	mid "auth-service/pkg/middleware"
	"auth-service/platform/database"

	"github.com/gofiber/fiber/v2"

	"os"
)

func main() {
	app := fiber.New()
	mid.FiberMiddleware(app)
	database.Init()
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
