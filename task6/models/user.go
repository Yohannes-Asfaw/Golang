package models

import (
	"github.com/go-playground/validator/v10"
)

var validateuser *validator.Validate

func init() {
	validateuser = validator.New()
}

type User struct {
	UserID   string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required,min=4"`
	Role     string `json:"role" bson:"role"`
}

func ValidateUser(user User) map[string]string {
	errors := make(map[string]string)

	err := validateuser.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			if tag == "required" {
				errors[field] = field + " is required"
			} else if tag == "min" {
				errors[field] = field + " must be at least " + err.Param() + " characters long"
			}
		}
	}

	return errors
}
