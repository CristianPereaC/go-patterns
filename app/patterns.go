package main

import (
	"../service"
)

var userService service.UserServiceInstance

func init() {
	userService = service.NewUserServiceInstance()
}

func main() {

	userService.CreateUser()
}
