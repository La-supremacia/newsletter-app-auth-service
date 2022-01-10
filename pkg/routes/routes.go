package routes

import (
	"auth-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/sign-up", controllers.PostSignUp)
	//route.Get("/", controllers.GetRoutes) //Vamos a usar esta ruta en raiz para que devuelva todas las rutas del microservicio, su metodo y su body
	//route.Post("/sign-in", controllers.PostSignIn)
}
