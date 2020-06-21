package main

import (
	"github.com/Cristien/go-patterns/app_context"
	"github.com/Cristien/go-patterns/routing"
)

func main() {

	context := app_context.BuildAppContext()

	apiHandlers := routing.NewApiHandlers()
	apiHandlers.SetUserHandler(&context.HandlerContext.UserHandler)

	routing.StartApiRouter(apiHandlers)
}
