package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prnndk/final-project-golang-pbkk/controller"
	"github.com/prnndk/final-project-golang-pbkk/service"
)

func Storage(route *gin.Engine, storageController controller.StorageController, jwtService service.JWTService) {
	routes := route.Group("/api/storage")
	{
		routes.POST("/upload", storageController.UploadFile)
	}
}
