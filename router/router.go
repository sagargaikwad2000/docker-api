package router

import (
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()

	router := r.Group("/api")

	InitContainerRoutes(router)
	InitImageRoutes(router)

	r.Run(":8080")
}
