package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show cart
// @Description Show a specific user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} ListCartResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /cart/{id} [get]
func GetCart(ctx *gin.Context) {
	id := ctx.Param("id")
	cart := []schemas.Cart{}

	if err := db.Where("user_id=?", id).Find(&cart).Error; err != nil {
		logger.Errorf("erro ao listar informações do carrinho: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao listar itens do carrinho")
		return
	}

	if len(cart) == 0 {
		SendError(ctx, http.StatusNotFound, "nenhum item encontrado para o usuário especificado")
		return
	}

	SendSuccess(ctx, "list-cart-items", cart)
}
