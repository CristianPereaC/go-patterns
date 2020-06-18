package business_services

import (
	"fmt"
	"testing"
)

type UserDaoMock struct {
	CreateUserDBMock func()
	UpdateUserDBMock func()
}

func (u UserDaoMock) CreateUserDB() {
	u.CreateUserDBMock()
}
func (u UserDaoMock) UpdateUserDB() {
	u.UpdateUserDBMock()
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	var UserServiceInstance = NewUserServiceInstance()
	var userDaoMock UserDaoMock

	userDaoMock.CreateUserDBMock = func() {
		fmt.Println("Mocked response")
		return
	}

	UserServiceInstance.Ctx.UsersDao = userDaoMock

	UserServiceInstance.CreateUser()
}
