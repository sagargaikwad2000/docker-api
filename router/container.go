package router

import (
	"github.com/Ashu23042000/docker-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitContainerRoutes(group *gin.RouterGroup) {

	router := group.Group("/container")
	{
		router.GET("/list", controllers.GetContainers)

		router.GET("/list/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "Hello0"})
		})

		router.POST("/create", func(c *gin.Context) {})
		router.DELETE("/delete", func(c *gin.Context) {})
		router.DELETE("/delete/:id", func(c *gin.Context) {})
	}
}
