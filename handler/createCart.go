package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create cart item
// @Description Create a new cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param computer_id path int true "Computer ID"
// @Success 200 {object} CreateCartItemResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /cart/{user_id}/{computer_id} [post]
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

	if err := db.Preload("Computer").Create(&cart).Error; err != nil {
		logger.Errorf("erro adicionando usuários: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	SendSuccess(ctx, "create-cart-item", cart)
}
