package router

import (
	"github.com/Naufra1/ByteApi/handler"
	"github.com/Naufra1/ByteApi/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Iniciando o handler
	handler.InitHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/computers", handler.ShowComputersHandler)

		v1.POST("/computer/", handler.CreateComputerHandler)

		v1.GET("/computer/:id", handler.ShowComputerHandler)

		v1.POST("/signup", handler.CreateUserHandler)

		v1.POST("/signin", handler.LoginUserHandler)

		v1.GET("/user/:id", middleware.VerifyToken, handler.ShowUserHandler)

		v1.PATCH("/user/:id", middleware.VerifyToken, handler.ChangeUserPassword)
	}
}
