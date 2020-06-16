package service

import (
	"fmt"
	"testing"
)

type UserDaoImplMock struct {
	MockCreateUserDB func()
	MockUpdateUserDB func()
}

func (u UserDaoImplMock) CreateUserDB() {
	u.MockCreateUserDB()
}
func (u UserDaoImplMock) UpdateUserDB() {
	u.MockUpdateUserDB()
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	var UserServiceInstance = NewUserServiceInstance()
	var userDaoImplMock UserDaoImplMock

	userDaoImplMock.MockCreateUserDB = func() {
		fmt.Println("Mocked response")
		return
	}

	UserServiceInstance.Ctx.UsersDao = userDaoImplMock

	UserServiceInstance.CreateUser()
}
