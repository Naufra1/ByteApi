package handler

import (
	"fmt"
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operação do handler %s feita com sucesso", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateComputerResponse struct {
	Message string `json:"message"`
	Data schemas.ComputerResponse `json:"data"`
}

type ListComputerResponse struct {
	Message string `json:"message"`
	Data schemas.ComputerResponse `json:"data"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Data schemas.UserResponse `json:"data"`
}

type CreateLoginResponse struct {
	Message string `json:"message"`
	Data schemas.LoginResponse `json:"data"`
}

type ListUserResponse struct {
	Message string `json:"message"`
	Data schemas.UserResponse `json:"data"`
}

type CreatePasswordChangeResponse struct {
	Message string `json:"message"`
	Data schemas.ChangePasswordResponse `json:"data"`
}