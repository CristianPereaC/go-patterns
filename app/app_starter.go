package app

import (
	"github.com/Cristien/go-patterns/app/app_context"
	"github.com/Cristien/go-patterns/routing"
)

func StartApplication(){
	context := app_context.BuildAppContext()

	apiHandlers := routing.NewApiHandlers()
	apiHandlers.SetUserHandler(&context.HandlerContext.UserHandler)

	routing.StartApiRouter(apiHandlers)
}
