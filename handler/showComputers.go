package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List computers
// @Description List all computers
// @Tags Computer
// @Accept json
// @Produce json 
// @Success 200 {object} ListComputerResponse
// @Failure 500 {object} ErrorResponse
// @Router /computers [get]
func ShowComputersHandler(ctx *gin.Context) {
	computers := []schemas.Computer{}

	if err := db.Find(&computers).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "erro ao listar computadores")
		return
	}
	SendSuccess(ctx, "list-computers", computers)
}
