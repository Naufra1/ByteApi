package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func ShowComputersHandler(ctx *gin.Context) {
	computers := []schemas.Computer{}

	if err := db.Find(&computers).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao listar computadores")
		return
	}
	SendSuccess(ctx, "list-computers", computers)
}
