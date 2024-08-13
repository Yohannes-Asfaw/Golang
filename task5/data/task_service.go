package data

import (
	"context"
	"errors"
	"fmt"
	"task5/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks() ([]models.Task, error) {
	cursor, err := collection.Find(context.Background(), bson.D{})
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
	err := collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&task)
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
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(id string, updatedTask models.Task) error {
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: updatedTask.Title},
				{Key: "description", Value: updatedTask.Description},
				{Key: "due_date", Value: updatedTask.DueDate},
				{Key: "status", Value: updatedTask.Status},
			}},
		},
	)
	return err
}

func DeleteTask(id string) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
