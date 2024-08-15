package Infrastructure


import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
)


type AuthenticationService struct {
	SecretKey []byte
}

func NewAuthenticationService(secretKey string) *AuthenticationService {
	return &AuthenticationService{
		SecretKey: []byte(secretKey),
	}
}

func (as *AuthenticationService) GenerateToken(userID, username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString(as.SecretKey)
}

func (as *AuthenticationService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return as.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
