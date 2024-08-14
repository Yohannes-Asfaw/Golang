package data

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "log"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Task_Collection *mongo.Collection
var User_Collection *mongo.Collection


func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal("Could not connect to MongoDB:", err)
    }

    log.Println("Connected to MongoDB!")
    Task_Collection=client.Database("task_manager").Collection("tasks")
    User_Collection=client.Database("task_manager").Collection("users")

    
}