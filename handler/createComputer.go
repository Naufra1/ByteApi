package handler

import (
	"net/http"

	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
)

func CreateComputerHandler(ctx *gin.Context) {
	request := CreateComputerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	computers := schemas.Computer{
		Name:        request.Name,
		Image:       request.Image,
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
		sendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	sendSuccess(ctx, "create-computer", computers)
}
