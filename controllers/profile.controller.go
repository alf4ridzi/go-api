package controllers

import "github.com/gin-gonic/gin"

type ProfileController struct {
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (p *ProfileController) GetProfiles(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	//tokenJwt := ctx.Request.Header.Get("Authorization")

}
