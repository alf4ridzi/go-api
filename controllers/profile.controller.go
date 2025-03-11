package controllers

import (
	"api/handlers"
	"api/services"
	"api/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	service *services.ProfileService
}

func NewProfileController(service *services.ProfileService) *ProfileController {
	return &ProfileController{service: service}
}

func (p *ProfileController) GetProfiles(ctx *gin.Context) {
	// create context
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	ctx.Header("Content-Type", "application/json")
	tokenJwt := ctx.Request.Header.Get("Authorization")

	if tokenJwt == "" {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", "missing token header", nil)
		return
	}

	if err := utils.VerifyTokenJwt(tokenJwt); err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
		return
	}

	// dec, err := utils.DecodeJwtToken(tokenJwt)
	// if err != nil {
	// 	handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
	// 	return
	// }

	username, err := utils.GetUsernameFromJwt(tokenJwt)
	if err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
		return
	}

	userProfile, err := p.service.GetUserProfile(reqCtx, username)
	if err != nil {
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
		return
	}

	handlers.ResponseJson(ctx, http.StatusOK, "success", "", userProfile)

}
