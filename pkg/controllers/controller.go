package controllers

import (
	"auth-service/pkg/models"
	signup "auth-service/pkg/services"
	"auth-service/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func PostSignUp(c *fiber.Ctx) error {
	u := new(models.User_SignUp_Request)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	hashPassword := utils.GeneratePassword(u.Password)

	createdUser := signup.NewUser_SignUp(u.Email, hashPassword, u.Name)
	responseUser := signup.NewUser(u.Email, createdUser.OrganizationId)

	err := mgm.Coll(createdUser).Create(createdUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully saved user", responseUser)

	return c.Status(fiber.StatusOK).JSON(responseUser)
}

/*

func GetObject(c *fiber.Ctx) error {
	creds := &Credentials{}
	coll := mgm.Coll(creds)
	fmt.Println(coll)
	id, _ := primitive.ObjectIDFromHex("61dc4140b13df2d9f41e950a")
	fmt.Println(id)
	_ = coll.FindByID(id, creds)

	//byId := coll.First(bson.M{}, creds)
	//found := coll.SimpleFind(creds, bson.M{})
	//fmt.Println(found)

	if creds == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(creds)
	}
	return c.Status(fiber.StatusOK).JSON(creds)
}
func GetObjects(c *fiber.Ctx) error {
	result := []Credentials{}

	err := mgm.Coll(&Credentials{}).SimpleFind(&result, bson.M{"email": bson.M{operator.Eq: "cesaasaf"}})
	fmt.Println(result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
*/

/* func GetTest(c *fiber.Ctx) error {
	ninja := models.User{
		Email:          "cesaraugp01@gmail.com",
		PasswordHash:   "fdsaf34240124k",
		OrganizationId: 3,
		Name:           "Cesar Perez",
	}

	return c.Status(fiber.StatusOK).JSON(ninja)
} */
