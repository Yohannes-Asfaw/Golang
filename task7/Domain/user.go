package Domain

import (
	"context"
)

type User struct {
	UserID   string `json:"userId" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	PromoteUser(ctx context.Context, username string) error
}

type UserUseCase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, username, password string) (string, error)
	PromoteUser(ctx context.Context, username string) error
}
