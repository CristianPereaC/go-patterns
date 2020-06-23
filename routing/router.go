package routing

import (
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	router *gin.Engine
}

func NewAppRouter() *AppRouter{
	return &AppRouter{
		router: gin.Default(),
	}
}

func (appRouter *AppRouter) MapUserResourceHandlers(handler UserHandler){
	appRouter.router.POST("/users", handler.CreateUser)
	appRouter.router.GET("/users/:id/private", handler.GetPrivateUser)
	appRouter.router.GET("/users/:id/public", handler.GetPublicUser)
}

func (appRouter *AppRouter) RunAppRouter() {
	appRouter.router.Run()
}