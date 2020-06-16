package main

import (
	"github.com/mercadolibre/go-patterns/service"
)

var userService = service.UserService

func init() {

}

func main() {

	userService.CreateUser()
}
