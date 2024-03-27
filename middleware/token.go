package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Naufra1/ByteApi/config"
	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var (
	JWT_SECRET = "JWT_SECRET_TOKEN"
	logger     *config.Logger
)

func CreateToken(id uint, name, email, birthday, address string, number int64) (string, error) {
	logger = config.GetLogger("middleware")
	secret := os.Getenv(JWT_SECRET)

	// Criando token jwt
	claims := jwt.MapClaims{
		"ID":       id,
		"name":     name,
		"email":    email,
		"birthday": birthday,
		"address":  address,
		"number":   number,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Errorf("erro ao geral token jwt: %v", err.Error())
		return "", err
	}

	return tokenString, err
}

func VerifyToken(ctx *gin.Context) {
	secret := os.Getenv(JWT_SECRET)
	tokenValue := RemoveBearer(ctx.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("token invalido")
	})
	if err != nil {
		logger.Errorf("erro ao verificar token: %v", err)
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Errorf("token invalido: %v", err)
		ctx.Abort()
		return
	}

	user := &schemas.User{
		Model: gorm.Model{
			ID: uint(claims["ID"].(float64)),
		},
		Name:     claims["name"].(string),
		Email:    claims["email"].(string),
		Birthday: claims["birthday"].(string),
		Address:  claims["address"].(string),
		Number:   int64(claims["number"].(float64)),
	}

	logger.Infof("usuário autorizado: %v", user)
}

func RemoveBearer(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}
	return token
}
