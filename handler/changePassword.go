package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ChangeUserPassword(ctx *gin.Context) {
	id := ctx.Param("id")
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.ValidatePassword(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	users := schemas.User{}

	if err := db.Where("id = ?", id).Find(&users).Error; err != nil {
		logger.Errorf("erro ao listar informações do usuário: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao listar senha")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users.CriptPassword), []byte(request.Password)); err == nil {
		SendError(ctx, http.StatusBadRequest, "Senha inválida, as senhas não podem ser iguais")
		logger.Errorf("senha inválida, as senhas não podem ser iguais: %v", err)
		return
	}

	cost := bcrypt.DefaultCost
	newHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), cost)
	if err != nil {
		logger.Errorf("erro ao encriptar a senha: %v", err)
		return
	}

	users = schemas.User{
		Model: gorm.Model{ID: users.ID},
		CriptPassword: string(newHash),
	}

	uintID := ConvertStr(id)
	if uintID == 0 {
		return
	}

	SendSuccess(ctx, "", users)

	if uintID != users.Model.ID {
		SendError(ctx, http.StatusBadRequest, "usuário invalido")
		return
	}

	if err := db.Model(&users).Where("id = ?", id).Update("cript_password", users.CriptPassword).Error; err != nil {
		logger.Errorf("erro ao acessar informações do usuário: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao acessar informações do usuário")
		return
	}

	SendSuccess(ctx, "patch-password", users.CriptPassword)
}
