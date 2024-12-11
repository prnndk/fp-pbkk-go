package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/controller"
	"github.com/prnndk/final-project-golang-pbkk/middleware"
	"github.com/prnndk/final-project-golang-pbkk/service"
)

func Event(route *gin.Engine, eventController controller.EventController, jwtService service.JWTService) {
	routes := route.Group("/api/event")
	{
		routes.GET("", middleware.Authenticate(jwtService), eventController.GetAllEvent)
		routes.GET("/:id", middleware.Authenticate(jwtService), eventController.GetSingleEvent)
	}
}

func Quota(route *gin.Engine, eventController controller.EventController, jwtService service.JWTService) {
	routes := route.Group("/api/event")
	{
		routes.PUT("/:id", middleware.Authenticate(jwtService), eventController.UpdateQuotaEvent)
	}
}
