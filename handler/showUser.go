package handler

import (
	"net/http"
	"strconv"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

type selectUserInfo struct {
	Id            uint
	Name          string
	Email         string
	Address       string
	Birthday      string
	Number        int64
}

// @BasePath /api/v1

// @Summary List user
// @Description List a specific user
// @Tags User
// @Accept json
// @Produce json 
// @Param id path int true "User ID"
// @Success 200 {object} ListUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /user/{id} [get]
func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	users := schemas.User{}
	userInfo := selectUserInfo{
		Id:       users.ID,
		Name:     users.Name,
		Email:    users.Email,
		Address:  users.Address,
		Birthday: users.Birthday,
		Number:   users.Number,
	}

	uintID := ConvertStr(id)
	if uintID == 0 {
		return
	}

	if err := db.Model(&users).Where("id= ?", uintID).Find(&userInfo).Error; err != nil {
		logger.Errorf("erro ao listar informações do usuário: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao listar usuário")
		return
	}

	if uintID != userInfo.Id {
		SendError(ctx, http.StatusBadRequest, "usuário invalido")
		return
	}

	SendSuccess(ctx, "get-user", userInfo)
}

func ConvertStr(s string) uint {
	string := s
	u64, err := strconv.ParseUint(string, 10, 32)
	if err != nil {
		logger.Errorf("erro ao converter string para uint: %v", err)
		return 0
	}
	wd := uint(u64)
	return wd
}
