package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCartHandler(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	computerID := ctx.Param("computer_id")

	uintUserID := ConvertStr(userID)
	uintComputerID := ConvertStr(computerID)
	if uintUserID == 0 || uintComputerID == 0 {
		return
	}

	cart := schemas.Cart{
		UserId:     uintUserID,
		ComputerId: uintComputerID,
	}

	if err := db.Create(&cart).Error; err != nil {
		logger.Errorf("erro adicionando usuários: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	SendSuccess(ctx, "create-cart-item", cart)
}
