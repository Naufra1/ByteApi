package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}
