package routing

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func StartApiRouter(apiHandlers apiHandlers){
	router = gin.Default()
	mapUserResourceHandlers(*apiHandlers.userHandler)
	router.Run()
}


func mapUserResourceHandlers(handler UserHandler){
	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id/private", handler.GetPrivateUser)
	router.GET("/users/:id/public", handler.GetPublicUser)
}