package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/middleware"
	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary Login user
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json 
// @Param request body CreateLoginRequest true "Request body"
// @Success 200 {object} CreateLoginResponse
// @Header  200 {string}  token     "token"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /signin [post]
func LoginUserHandler(ctx *gin.Context) {
	request := CreateLoginRequest{}

	ctx.BindJSON(&request)

	if err := request.ValidateUser(); err != nil {
		logger.Errorf("erro na validação de usuário: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	users := schemas.User{
		Email: request.Email,
	}
	if err := db.Where("email= ?", request.Email).Find(&users).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao listar usuário")
		return
	}

	// Comparando senhas
	if err := bcrypt.CompareHashAndPassword([]byte(users.CriptPassword), []byte(request.Password)); err != nil {
		SendError(ctx, http.StatusBadRequest, "Senha inválida, digite a senha correta")
		logger.Errorf("senha inválida, digite a senha correta")
		return
	}

	// Criando Token JWT
	token, err := middleware.CreateToken(users.Model.ID, users.Name, users.Email, users.Birthday, users.Address, users.Number)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao gerar o token")
		logger.Errorf("erro ao gerar o token")
		return
	}

	ctx.Header("Authorization", token)
	SendSuccess(ctx, "login-user", request)
}
