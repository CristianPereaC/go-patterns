package main

import (
	"github.com/Cristien/go-patterns/business_layer/business_services"
	"github.com/Cristien/go-patterns/dao/memory_dao"
	"github.com/Cristien/go-patterns/web/http_handlers"
)

var(
	usersDao = memory_dao.NewUserDaoImplInstance()
	usersService = business_services.NewUserServiceInstance(usersDao)
	userHandler = http_handlers.NewuserHandler(usersService)
)

func init() {

}

func main() {

}
