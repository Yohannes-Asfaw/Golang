package data

import (
	"context"
	"errors"
	"task6/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(ctx context.Context, user *models.User) error {
	
	users, _ := User_Collection.CountDocuments(context.Background(), bson.M{})
	if users == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.UserID = primitive.NewObjectID().Hex()

	_, err = User_Collection.InsertOne(ctx, user)
	return err
}


func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := User_Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, errors.New("invalid username or password")
	}
	return &user, nil
}

func GetUserByID(ctx context.Context, id string) (*models.User, error) {

    var user models.User
    if err := User_Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
        return nil, errors.New("user not found")
    }

    return &user, nil
}


func PromoteUser(username string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"username": username}
    update := bson.M{"$set": bson.M{"role": "admin"}}

    result, err := User_Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 {
        return errors.New("user not found")
    }

    return nil
}