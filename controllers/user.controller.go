package controllers

import (
	"api/handlers"
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handlers.ResponseJson(ctx, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	if err := c.service.RegisterUser(&user); err != nil {
		handlers.ResponseJson(ctx, http.StatusConflict, "fail", err.Error(), nil)
		return
	}

	handlers.ResponseJson(ctx, http.StatusOK, "success", "", nil)
}

func (c *UserController) Login(ctx *gin.Context) {
	var user models.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handlers.ResponseJson(ctx, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	if err := c.service.VerifyLogin(&user); err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "error", err.Error(), nil)
		return
	}

	handlers.ResponseJson(ctx, http.StatusOK, "success", "", nil)
}
