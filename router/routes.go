package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.GET("/computers", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Success",
			})
		})

		v1.GET("/computer", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Success",
			})
		})

		v1.GET("/signup", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Success",
			})
		})

		v1.GET("/signin", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Success",
			})
		})

		v1.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Success",
			})
		})
	}
}	