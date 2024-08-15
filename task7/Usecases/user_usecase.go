package Usecases

import (
	"context"
	"errors"
	"task7/Domain"
	"task7/Infrastructure"
)
type UserUseCase struct {
	UserRepo        Domain.UserRepository
	JWTService      Infrastructure.JWTService
	PasswordService Infrastructure.PasswordService
}

func NewUserUseCase(userRepo Domain.UserRepository, jwtService Infrastructure.JWTService, passwordService Infrastructure.PasswordService) *UserUseCase {
	return &UserUseCase{
		UserRepo:        userRepo,
		JWTService:      jwtService,
		PasswordService: passwordService,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, user *Domain.User) error {
	hashedPassword, err := uc.PasswordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return uc.UserRepo.Create(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := uc.UserRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Use the PasswordService to compare the hashed password
	if err := uc.PasswordService.CompareHashAndPassword(user.Password, password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Use the JWTService to generate the token
	token, err := uc.JWTService.GenerateToken(user.UserID, user.Username, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (uc *UserUseCase) PromoteUser(ctx context.Context, username string) error {
	return uc.UserRepo.PromoteUser(ctx, username)
}
