package router

import (
	"github.com/Ashu23042000/docker-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitImageRoutes(group *gin.RouterGroup) {
	router := group.Group("/image")
	{
		router.GET("/list", controllers.GetImages)
	}
}
