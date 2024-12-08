package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/controller"
	"github.com/prnndk/final-project-golang-pbkk/middleware"
	"github.com/prnndk/final-project-golang-pbkk/service"
)

func Pembayaran(route *gin.Engine, pembayaranController controller.PembayaranController, jwtService service.JWTService) {
	routes := route.Group("/api/pembayaran")
	{
		routes.POST("/create", middleware.Authenticate(jwtService), pembayaranController.CreatePembayaran)
	}
}
