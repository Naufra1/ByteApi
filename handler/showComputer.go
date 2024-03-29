package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func ShowComputerHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	computers := schemas.Computer{}
	if err := db.Where("id = ?", id).Find(&computers).Error; err != nil {
		logger.Errorf("erro ao listar informações do computador: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao listar informaões do computador")
		return
	}

	uintID := ConvertStr(id)
	if uintID == 0 {
		return
	}
	if uintID != computers.Model.ID {
		SendError(ctx, http.StatusBadRequest, "computador invalido")
		return
	}

	SendSuccess(ctx, "get-computer", computers)
}
