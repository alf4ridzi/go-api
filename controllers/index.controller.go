package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   nil,
	})
}
