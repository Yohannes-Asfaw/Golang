package models

import (
	"time"
	"github.com/go-playground/validator/v10"
)

var validatetask *validator.Validate

func init() {
	validatetask = validator.New()
}

type Task struct {
	TaskID      string    `json:"task_id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title" validate:"required,min=5"`
	Description string    `json:"description" bson:"description" validate:"required,min=5"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status" validate:"required"`
}

func ValidateTask(task Task) map[string]string {
	errors := make(map[string]string)

	err := validatetask.Struct(task)
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
