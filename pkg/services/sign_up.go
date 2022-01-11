package signup

import "auth-service/pkg/models"

func NewUser_SignUp(email string, organization_id int) *models.User_SignUp_Response {
	return &models.User_SignUp_Response{
		Email:          email,
		OrganizationId: organization_id,
	}
}

func NewUser(email string, password string, name string) *models.User {
	//GO TO ORGANIZATION MICROSERVICE AND GENERATE A NEW ORG
	orgId := 1
	return &models.User{
		Email:          email,
		Name:           name,
		PasswordHash:   password,
		OrganizationId: orgId,
	}
}
