package controllers

import (
	"api/database"
	"api/handlers"
	"api/initializers"
	"api/models"
	"api/repositories"
	"api/services"
	"api/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	cookiesManager := utils.Cookies{}
	authToken, err := cookiesManager.GetCookie(ctx, "auth_token")
	if err == nil && utils.VerifyJwtAuth(authToken) == nil {
		handlers.ResponseJson(ctx, http.StatusAccepted, "success", "Already log in", nil)
		return
	}

	authToken, refreshToken, err := c.service.VerifyLogin(reqCtx, &user)
	if err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "error", err.Error(), nil)
		return
	}

	cookiesManager.SetCookie(ctx, "auth_token", authToken, 15*time.Minute)
	cookiesManager.SetCookie(ctx, "refresh_token", refreshToken, 24*time.Hour)

	handlers.ResponseJson(ctx, http.StatusOK, "success", "Success Login", nil)
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	cookiesManager := utils.Cookies{}
	refresh_token, err := cookiesManager.GetCookie(ctx, "refresh_token")
	if err != nil {
		handlers.ResponseJson(ctx, 200, "fail", err.Error(), nil)
		return
	}

	if refresh_token == "" {
		handlers.ResponseJson(ctx, 200, "fail", "Refresh token is not found", nil)
		return
	}

	refresh_token_secret := utils.Env("REFRESH_SECRET")

	if err := utils.VerifyJwtRefresh(refresh_token); err != nil {
		handlers.ResponseJson(ctx, 200, "fail", err.Error(), nil)
		return
	}

	username, err := utils.GetValueJwt(initializers.GetRefreshSecret(), refresh_token, "username")
	if err != nil {
		handlers.ResponseJson(ctx, 500, "error", err.Error(), nil)
		return
	}

	userRepositories := repositories.NewUserRepositories(database.DB)

	Reqctx, cancel := context.WithTimeout(ctx.Request.Context(), 10)
	defer cancel()

	user, err := userRepositories.GetUserByUsername(Reqctx, username)
	if err != nil {
		handlers.ResponseJson(ctx, 500, "error", err.Error(), nil)
		return
	}

	if user == nil {
		handlers.ResponseJson(ctx, 200, "fail", "user is not found", nil)
		return
	}

	claims := jwt.MapClaims{
		"username": username,
		"role":     user.Role,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}

	token, err := utils.CreateJwtToken(refresh_token_secret, claims)
	if err != nil {
		handlers.ResponseJson(ctx, 200, "fail", "error create new token", nil)
		return
	}

	cookiesManager.SetCookie(ctx, "auth_token", token, 15*time.Minute)

	data := map[string]string{
		"token": token,
	}

	handlers.ResponseJson(ctx, 200, "success", "success create new token", data)
}
