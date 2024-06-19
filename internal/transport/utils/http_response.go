package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteErrorResponse(ctx *gin.Context, err error, code int) {
	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}
