package Repositories

import (
	"context"
	"errors"
	"task7/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) Domain.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *Domain.User) error {
	users, _ := ur.DB.CountDocuments(context.Background(), bson.M{})
	if users == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	user.UserID = primitive.NewObjectID().Hex()

	_, err := ur.DB.InsertOne(ctx, user)
	return err
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (*Domain.User, error) {
	var user Domain.User
	if err := ur.DB.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, errors.New("invalid username or password")
	}
	return &user, nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id string) (*Domain.User, error) {
	var user Domain.User
	if err := ur.DB.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (ur *UserRepository) PromoteUser(ctx context.Context, username string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := ur.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
