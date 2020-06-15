package main

import (
	"../service"
)

var userService = service.GetUserServiceInstance()

func main() {

	userService.CreateUser(userService.Options)
}
