package Domain

import (
	"context"
	"time"
)


type Task struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string   `json:"status" bson:"status"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
}
type TaskRepository interface {
	Create(c context.Context, task *Task) (*Task, error)
	Update(c context.Context, id string, task *Task) (*Task, error)
	Delete(c context.Context, id string) error
	GetAll(c context.Context) (*[]*Task, error)
	GetById(c context.Context, taskId string) (*Task, error)
}

type TaskUseCase interface {
	Create(c context.Context, payload *Task) (*Task, error)
	Update(c context.Context, taskId string, payload *Task) (*Task, error)
	Delete(c context.Context, taskId string) error
	GetAll(c context.Context) (*[]*Task, error)
	GetById(c context.Context, taskId string) (*Task, error)
}