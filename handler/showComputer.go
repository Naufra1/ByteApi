package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List computer
// @Description List a specific computer
// @Tags Computer
// @Accept json
// @Produce json 
// @Param id path int true "Computer ID"
// @Success 200 {object} ListComputerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /computer/{id} [get]
func ShowComputerHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	computers := schemas.Computer{}

	uintID := ConvertStr(id)
	if uintID == 0 {
		return
	}

	if err := db.Where("id = ?", id).Find(&computers).Error; err != nil {
		logger.Errorf("erro ao listar informações do computador: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao listar informaões do computador")
		return
	}

	if uintID != computers.Model.ID {
		SendError(ctx, http.StatusBadRequest, "computador invalido")
		return
	}
	
	SendSuccess(ctx, "get-computer", computers)
}
