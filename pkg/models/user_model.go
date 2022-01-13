package models

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash     string `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	Name             string `db:"name" json:"name" validate:"lte=255"`
}

type User_SignUp_Request struct {
	Email    string `db:"email" json:"email" validate:"required,email,lte=255"`
	Password string `db:"password" json:"password,omitempty" validate:"required,lte=255"`
	Name     string `db:"name" json:"name,omitempty" validate:"lte=255"`
}

type User_SignUp_Response struct {
	Email string `db:"email" json:"email" validate:"required,email,lte=255"`
	Name  string `db:"name" json:"name" validate:"lte=255"`
	Id    string `json:"_id,omitempty" validate:"lte=255"`
}
