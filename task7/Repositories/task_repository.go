package Repositories

import (
	"context"
	"errors"
	"fmt"
	"task7/Domain"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	DB *mongo.Collection
}

func NewTaskRepository(db *mongo.Collection) Domain.TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *Domain.Task) (*Domain.Task, error) {
	task.DueDate = time.Now().AddDate(0, 0, 7)
	task.Id = primitive.NewObjectID().Hex()
	_, err := r.DB.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) Update(ctx context.Context, id string, updatedTask *Domain.Task) (*Domain.Task, error) {
	update := bson.M{}
	setFields := bson.M{}

	if updatedTask.Title != "" {
		setFields["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		setFields["description"] = updatedTask.Description
	}
	if !updatedTask.DueDate.IsZero() {
		setFields["dueDate"] = updatedTask.DueDate
	}
	if updatedTask.Status != "" {
		setFields["status"] = updatedTask.Status
	}

	if len(setFields) > 0 {
		update["$set"] = setFields
	}

	_, err := r.DB.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	var task Domain.Task
	err = r.DB.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.DB.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *TaskRepository) GetAll(ctx context.Context) (*[]*Domain.Task, error) {
	cursor, err := r.DB.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var tasks []*Domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *TaskRepository) GetById(ctx context.Context, taskId string) (*Domain.Task, error) {
	var task Domain.Task
	err := r.DB.FindOne(ctx, bson.D{{Key: "_id", Value: taskId}}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, fmt.Errorf("error fetching task: %v", err)
	}
	return &task, nil
}
