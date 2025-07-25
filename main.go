package main

import (
	"event-management/config"
	"event-management/controllers"
	"event-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	controllers.InitAuthController()
	controllers.InitEventController()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
