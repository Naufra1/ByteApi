package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete cart item
// @Description Delete a item from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param item_id path int true "Item ID"
// @Success 200 {object} CreateCartItemResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /cart/{item_id} [delete]
func DeleteItem(ctx *gin.Context) {
	id := ctx.Param("item_id")
	cart := schemas.Cart{}

	if err := db.Preload("Computer").Delete(&cart, id).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao deletar carrinho")
		return
	}
	SendSuccess(ctx, "delete-cart", nil)
}
