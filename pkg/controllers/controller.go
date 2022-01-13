package controllers

import (
	"auth-service/pkg/models"
	signup "auth-service/pkg/services"
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

	createdUser := signup.NewUser(u.Email, hashPassword, u.Name)
	responseUser := signup.NewUser_SignUp(createdUser.Email, createdUser.Name, "")

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

	dbUser := signup.NewUser(u.Email, "", "")

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

	responseUser := signup.NewUser_SignUp(dbUser.Email, dbUser.Name, dbUser.ID.Hex())

	fmt.Println("Successfully logged user", dbUser.Email)

	c.Set("Authorization", "Bearer "+accessToken)
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
