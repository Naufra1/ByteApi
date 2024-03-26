package computer

import (
	"net/http"

	"github.com/Naufra1/ByteApi/config"
	"github.com/Naufra1/ByteApi/handler"
	"github.com/Naufra1/ByteApi/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func CreateComputerHandler(ctx *gin.Context) {
	request := CreateComputerRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("erro na validação: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
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
		handler.SendError(ctx, http.StatusInternalServerError, "erro ao criar tabela")
		return
	}
	handler.SendSuccess(ctx, "create-computer", computers)
}
