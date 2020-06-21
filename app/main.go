package main

import (
	"github.com/Cristien/go-patterns/app_context"
	"github.com/Cristien/go-patterns/routing"
)

func main() {
	context := app_context.BuildAppContext()
	routing.MapUserResourceHandlers(context.HandlerContext.UserHandler)
	routing.Router.Run()
}
