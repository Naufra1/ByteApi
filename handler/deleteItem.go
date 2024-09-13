package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")
	cart := schemas.Cart{}

	if err := db.Delete(&cart, id).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao deletar carrinho")
		return
	}
	SendSuccess(ctx, "delete-cart", nil)
}
