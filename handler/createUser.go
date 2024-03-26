package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), cost)
	if err != nil {
		logger.Errorf("erro ao encriptar a senha: %v", err)
	}

	users := schemas.User{
		Name:          request.Name,
		Email:         request.Email,
		CriptPassword: string(hash),
		Birthday:      request.Birthday,
		Number:        request.Number,
		Address:       request.Address,
	}

	if err := db.Create(&users).Error; err != nil {
		logger.Errorf("erro adicionando usuários: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	SendSuccess(ctx, "create-user", users)
}
