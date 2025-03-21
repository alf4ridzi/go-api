package controllers

import (
	"api/handlers"
	"api/models"
	"api/services"
	"api/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var reg models.Register

	// create context
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&reg); err != nil {
		handlers.ResponseJson(ctx, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	user := models.User{
		Username: reg.Username,
		Name:     reg.Name,
		Email:    &reg.Email,
		Password: reg.Password,
	}

	if err := c.service.RegisterUser(reqCtx, &user); err != nil {
		handlers.ResponseJson(ctx, http.StatusConflict, "fail", err.Error(), nil)
		return
	}

	handlers.ResponseJson(ctx, http.StatusOK, "success", "Success create account!", nil)
}

func (c *AuthController) Login(ctx *gin.Context) {
	// create context
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	var user models.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handlers.ResponseJson(ctx, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	authToken, refreshToken, err := c.service.VerifyLogin(reqCtx, &user)
	if err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "error", err.Error(), nil)
		return
	}

	cookiesManager := utils.Cookies{}
	cookiesManager.SetCookie(ctx, "auth_token", authToken, 15*time.Minute)
	cookiesManager.SetCookie(ctx, "refresh_token", refreshToken, 24*time.Hour)

	handlers.ResponseJson(ctx, http.StatusOK, "success", "Success Login", nil)
}
