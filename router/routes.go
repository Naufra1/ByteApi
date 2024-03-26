package router

import (
	"github.com/Naufra1/ByteApi/handler"
	"github.com/Naufra1/ByteApi/handler/computer"
	"github.com/Naufra1/ByteApi/handler/user"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Iniciando o handler
	handler.InitHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/computers", computer.ShowComputersHandler)

		v1.POST("/computer", computer.CreateComputerHandler)

		v1.GET("/computer", computer.ShowComputersHandler)

		v1.POST("/signup", user.CreateUserHandler)

		v1.GET("/signin", user.LoginUserHandler)

		v1.GET("/user", user.ShowUserHandler)
	}
}
