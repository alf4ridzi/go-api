package controllers

import (
	"api/handlers"
	"api/services"
	"api/utils"
	"context"
	"log"
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
	cookiesManager := utils.Cookies{}
	tokenJwt, err := cookiesManager.GetCookie(ctx, "auth_token")

	if err != nil {
		handlers.ResponseJson(ctx, http.StatusInternalServerError, "error", "unable to read cookies", nil)
		return
	}

	if err := utils.VerifyJwtAuth(tokenJwt); err != nil {
		log.Println(err)
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", "Invalid token", nil)
		return
	}

	// dec, err := utils.DecodeJwtToken(tokenJwt)
	// if err != nil {
	// 	handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
	// 	return
	// }

	username, err := utils.GetUsernameFromJwtAuth(tokenJwt)
	if err != nil {
		log.Println(err)
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", "Error get username", nil)
		return
	}

	userProfile, err := p.service.GetUserProfile(reqCtx, username)
	if err != nil {
		log.Println(err)
		handlers.ResponseJson(ctx, http.StatusUnauthorized, "fail", err.Error(), nil)
		return
	}

	handlers.ResponseJson(ctx, http.StatusOK, "success", "", userProfile)

}
