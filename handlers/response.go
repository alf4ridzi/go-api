package handlers

import "github.com/gin-gonic/gin"

func ResponseJson(ctx *gin.Context, statuscode int, status string, message string, data any) {
	ctx.JSON(statuscode, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
