package routing

import "github.com/gin-gonic/gin"

var (
	Router *gin.Engine
)


func init(){
	Router = gin.Default()
}


func MapUserResourceHandlers(handler UserHandler){
	Router.POST("/user", handler.CreateUser)
	Router.GET("/user/:id/private", handler.GetPrivateUser)
	Router.GET("/user/:id/public", handler.GetPublicUser)
}