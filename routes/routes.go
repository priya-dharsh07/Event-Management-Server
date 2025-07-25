package routes

import (
	"github.com/gin-gonic/gin"
	"event-management/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/events", controllers.CreateEvent)
}
