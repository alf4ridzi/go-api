package routes

import (
	"api/controllers"
	"api/database"
	"api/repositories"
	"api/services"

	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	db := database.DB

	// index
	indexController := controllers.NewIndexController()

	// user repo
	userRepo := repositories.NewUserRepositories(db)
	// auth service
	authService := services.NewAuthService(userRepo)
	// auth controller
	userController := controllers.NewAuthController(authService)

	api := r.Group("/api")
	{
		// index
		api.GET("/", indexController.Index)
		// login
		api.POST("/login", userController.Login)
		// register
		api.POST("/register", userController.Register)
	}
}
