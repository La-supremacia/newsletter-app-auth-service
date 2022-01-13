package controllers

import (
	"auth-service/pkg/models"
	"auth-service/pkg/services"
	"auth-service/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func PostSignUp(c *fiber.Ctx) error {
	u := new(models.User_SignUp_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	hashPassword := utils.GeneratePassword(u.Password)

	createdUser := services.NewUser(u.Email, hashPassword, u.Name)
	responseUser := services.NewUser_SignUp(createdUser.Email, createdUser.Name)


	err := mgm.Coll(createdUser).Create(createdUser)

	responseUser.Id = createdUser.ID.Hex()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully saved user", responseUser)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func PostSignIn(c *fiber.Ctx) error {
	u := new(models.User_SignUp_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	dbUser := services.NewUser(u.Email, "", "")

	err := mgm.Coll(dbUser).First(bson.M{"email": u.Email}, dbUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if !utils.ComparePasswords(dbUser.PasswordHash, u.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON("Wrong password")
	}

	accessToken, err := utils.GenerateNewAccessToken(dbUser.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	responseUser := services.NewUser_SignUp(dbUser.Email, dbUser.Name)


	fmt.Println("Successfully logged user", dbUser.Email)

	c.Set("Authorization", "Bearer "+accessToken)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func GetRoutes(c *fiber.Ctx) error {
	signInRoute := c.App().GetRoute("Sign In")
	signUpRoute := c.App().GetRoute("Sign Up")

	var ruts = [3]models.Route{services.NewRoute(signInRoute.Path, signInRoute.Method, signInRoute.Name, "email:string;password:string"),
		services.NewRoute(signUpRoute.Path, signUpRoute.Method, signUpRoute.Name, "email:string;password:string;name:string"),
		services.NewRoute(c.Route().Path, c.Route().Method, c.Route().Name, ""),
	}
	return c.Status(fiber.StatusOK).JSON(ruts)
}
