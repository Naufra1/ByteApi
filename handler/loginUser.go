package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}
