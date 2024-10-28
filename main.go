package main

import (
	"github.com/Naufra1/ByteApi/config"
	"github.com/Naufra1/ByteApi/router"
)

var (
	logger *config.Logger
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logger = config.GetLogger("main")

	// Inicializando a Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Erro na inicialização da configuração: %v", err)
		return
	}

	// Incializando o Router
	router.Initialize()
}
