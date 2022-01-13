package routes

import (
	"auth-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/sign-up", controllers.PostSignUp).Name("Sign Up")
	route.Post("/sign-in", controllers.PostSignIn).Name("Sign In")
	route.Get("/", controllers.GetRoutes).Name("Root")
}
