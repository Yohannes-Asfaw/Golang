package Infrastructure

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type AuthorizationService struct{}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

func (as *AuthorizationService) Authorize(token *jwt.Token, requiredRole string) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("unauthorized")
	}

	role := claims["role"].(string)
	if role != requiredRole {
		return errors.New("forbidden: insufficient permissions")
	}

	return nil
}
