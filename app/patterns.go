package main

import (
	"../service"
)

var userService service.UserServiceImpl

func init() {
	userService = service.NewUserServiceInstance()
}

func main() {

	userService.CreateUser()
}
