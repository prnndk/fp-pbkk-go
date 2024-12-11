package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/controller"
	"github.com/prnndk/final-project-golang-pbkk/middleware"
	"github.com/prnndk/final-project-golang-pbkk/service"
)

func UserTicket(route *gin.Engine, userTicketController controller.UserTicketController, jwtService service.JWTService) {
	routes := route.Group("/api/ticket")
	{
		routes.POST("", middleware.Authenticate(jwtService), userTicketController.BuyTicket)
		routes.GET("/me", middleware.Authenticate(jwtService), userTicketController.GetUserTicket)
		routes.GET("/:id", middleware.Authenticate(jwtService), userTicketController.GetUserTicketById)
	}
}
