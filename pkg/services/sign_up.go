package services

import "auth-service/pkg/models"

func NewUser_SignUp(email string, name string) *models.User_SignUp_Response {
	return &models.User_SignUp_Response{
		Email: email,
		Name:  name,
	}
}

func NewUser(email string, password string, name string) *models.User {
	return &models.User{
		Email:        email,
		Name:         name,
		PasswordHash: password,
	}
}
