package service

import (
	"fmt"
	"testing"
)

func TestUserServiceImpl_CreateUser(t *testing.T) {

	var UserServiceInstance = GetUserServiceInstance()

	UserServiceInstance.Options.UsersDao.CreateUserDB = func() {
		fmt.Println("Mocked response")
		return
	}

	UserServiceInstance.CreateUser(UserServiceInstance.Options)
}
