package routing

import "github.com/gin-gonic/gin"

var (
	Router *gin.Engine
)


func init(){
	Router = gin.Default()
}


func MapUserResourceHandlers(handler UserHandler){
	Router.POST("/users", handler.CreateUser)
	Router.GET("/users/:id/private", handler.GetPrivateUser)
	Router.GET("/users/:id/public", handler.GetPublicUser)
}