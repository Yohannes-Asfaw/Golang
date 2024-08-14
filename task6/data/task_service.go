package data

import (
	"context"
	"errors"
	"fmt"
	"task6/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks() ([]models.Task, error) {
	cursor, err := Task_Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {

	var task models.Task
	err := Task_Collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, fmt.Errorf("error fetching task: %v", err)
	}

	return task, nil
}

func CreateTask(task models.Task) error {
	task.DueDate = time.Now().AddDate(0, 0, 7)
	_, err := Task_Collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	update := bson.M{}
	setFields := bson.M{}

	if updatedTask.Title != "" {
		setFields["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		setFields["description"] = updatedTask.Description
	}
	if !updatedTask.DueDate.IsZero() {
		setFields["due_date"] = updatedTask.DueDate
	}
	if updatedTask.Status != "" {
		setFields["status"] = updatedTask.Status
	}

	if len(setFields) > 0 {
		update["$set"] = setFields
	}

	_, err := Task_Collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = Task_Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTask(id string) error {
	_, err := Task_Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
