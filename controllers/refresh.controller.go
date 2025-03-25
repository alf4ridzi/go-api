package controllers

import "api/services"

type RefreshController struct {
	service *services.RefreshService
}

func NewRefreshController(service *services.RefreshService) *RefreshController {
	return &RefreshController{service: service}
}

func (r *RefreshController) RefreshToken() {

}
