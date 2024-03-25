package router

import (
	"github.com/Naufra1/ByteApi/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Iniciando o handler
	handler.InitHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/computers", handler.ShowComputersHandler)

		v1.POST("/computer", handler.CreateComputerHandler)

		v1.GET("/computer", handler.ShowComputersHandler)

		v1.POST("/signup", handler.CreateUserHandler)

		v1.GET("/signin", handler.LoginUserHandler)

		v1.GET("/user", handler.ShowUserHandler)
	}
}
