package router

import (
	"github.com/Naufra1/ByteApi/handler"
	"github.com/Naufra1/ByteApi/middleware"
	"github.com/gin-gonic/gin"
	docs "github.com/Naufra1/ByteApi/docs"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	// Iniciando o handler
	handler.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
	{
		v1.GET("/computers", handler.ShowComputersHandler)

		v1.POST("/computer", handler.CreateComputerHandler)

		v1.GET("/computer/:id", handler.ShowComputerHandler)

		v1.POST("/signup", handler.CreateUserHandler)

		v1.POST("/signin", handler.LoginUserHandler)

		v1.GET("/user/:id", middleware.VerifyToken, handler.ShowUserHandler)

		v1.PATCH("/user/:id", middleware.VerifyToken, handler.ChangeUserPassword)

		v1.GET("/cart/:id", middleware.VerifyToken, handler.GetCart)

		v1.POST("/cart/:user_id/:computer_id", middleware.VerifyToken, handler.CreateCartHandler)

		v1.DELETE("/cart/:id", middleware.VerifyToken, handler.DeleteItem)
	}
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
