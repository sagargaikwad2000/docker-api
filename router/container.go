package router

import (
	"github.com/Ashu23042000/docker-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitContainerRoutes(group *gin.RouterGroup) {

	router := group.Group("/container")
	{
		router.GET("/list", controllers.GetContainers)
		router.POST("/create", controllers.CreateContainer)
		router.GET("/inspect/:id", controllers.GetContainers)

	}
}
