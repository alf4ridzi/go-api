package main

import (
	"api/database"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	database.InitDB()

	r := gin.Default()
	routes.MapRoutes(r)
	r.Run(":8080")
}
