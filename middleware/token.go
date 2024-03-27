package middleware

import (
	"os"
	"time"

	"github.com/Naufra1/ByteApi/config"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET = "JWT_SECRET_TOKEN"
	logger     *config.Logger
)

func CreateToken(email string) (string, error) {
	logger = config.GetLogger("middleware")
	secret := os.Getenv(JWT_SECRET)

	// Criando token jwt
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Errorf("erro ao geral token jwt: %v", err.Error())
		return "", err
	}

	return tokenString, err
}
