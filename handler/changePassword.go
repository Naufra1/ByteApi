package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// @Summary Change password
// @Description Change the user password
// @Tags User
// @Accept json
// @Produce json 
// @Param id path int true "User ID"
// @Param request body ChangePasswordRequest true "Request body"
// @Success 200 {object} CreatePasswordChangeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /user/{id} [patch]
func ChangeUserPassword(ctx *gin.Context) {
	id := ctx.Param("id")
	request := ChangePasswordRequest{}

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

	if uintID != users.Model.ID {
		SendError(ctx, http.StatusBadRequest, "usuário invalido")
		return
	}
	
	if err := db.Model(&users).Where("id = ?", id).Update("cript_password", users.CriptPassword).Error; err != nil {
		logger.Errorf("erro ao acessar informações do usuário: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao acessar informações do usuário")
		return
	}

	SendSuccess(ctx, "patch-password", "null")
}
