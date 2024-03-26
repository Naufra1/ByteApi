package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowComputerHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Success",
	})
}
