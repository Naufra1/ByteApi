package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create computer
// @Description Create a new computer
// @Tags Computer
// @Accept json
// @Produce json 
// @Param request body CreateComputerRequest true "Request body"
// @Success 200 {object} CreateComputerResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /computer [post]
func CreateComputerHandler(ctx *gin.Context) {
	request := CreateComputerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	computers := schemas.Computer{
		Name:        request.Name,
		Price:       request.Price,
		Case:        request.Case,
		Cpu:         request.Cpu,
		Cooler:      request.Cooler,
		Graphics:    request.Graphics,
		Motherboard: request.Motherboard,
		Ram:         request.Ram,
		Storage:     request.Storage,
	}

	if err := db.Create(&computers).Error; err != nil {
		logger.Errorf("erro adicionando computador: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	SendSuccess(ctx, "create-computer", computers)
}
