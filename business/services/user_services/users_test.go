package user_services

import (
	"errors"
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCreateUserAndDaoFailsThenError(t *testing.T) {
	dao := newUserDaoMockInstance()
	dao.HandleSaveUser = func() (*users_data_transfer.UserDb, error) {
		return nil, errors.New("error creating user")
	}
	service := NewUserServiceInstance(dao)
	request := users_data_transfer.CreateUserRequest{
		FirstName: "Juan",
		LastName: "Perez",
		Email: "pepe@pepe.com",
		Password: "1243254365",
	}
	id, err := service.CreateUser(request)
	expectedErr := errors.New("error creating user")
	assert.Equal(t, "", id)
	assert.Equal(t, expectedErr, err)
}

func TestWhenCreateUserSuccessfulThenUserHasId(t *testing.T) {
	dao := newUserDaoMockInstance()
	dao.HandleSaveUser = func() (*users_data_transfer.UserDb, error) {
		userDb := users_data_transfer.UserDb{
			Id: "1234-1324-1324",
			FirstName: "Juan",
			LastName: "Perez",
			Email: "pepe@pepe.com",
			Password: "1243254365",
			Version: 1,
		}
		return &userDb, nil
	}
	service := NewUserServiceInstance(dao)
	request := users_data_transfer.CreateUserRequest{
		FirstName: "Juan",
		LastName: "Perez",
		Email: "pepe@pepe.com",
		Password: "1243254365",
	}
	id, err := service.CreateUser(request)
	assert.Equal(t, "1234-1324-1324", id)
	assert.Nil(t, err)
}


func TestWhenGetPrivateUserAndUserNotExistThenNil(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*users_data_transfer.UserDb) {
		return nil
	}
	service := NewUserServiceInstance(dao)
	privateUser := service.GetPrivateUser("214")
	assert.Nil(t, privateUser)
}

func TestWhenGetPrivateUserAndUserExistThenUser(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*users_data_transfer.UserDb) {
		userDb := users_data_transfer.UserDb{
			Id: "214",
			FirstName: "Juan",
			LastName: "Perez",
			Version: 1,
			Email: "pepe@pepe.com",
		}
		return &userDb
	}
	service := NewUserServiceInstance(dao)
	privateUser := service.GetPrivateUser("214")
	expectedPrivateUser := users_data_transfer.PrivateUserDto{
		Id: "214",
		FirstName: "Juan",
		LastName: "Perez",
		Email: "pepe@pepe.com",
		FullName: "Perez, Juan",
	}
	assert.Equal(t, expectedPrivateUser, *privateUser)
}

func TestWhenGetPublicUserAndUserNotExistThenNil(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*users_data_transfer.UserDb) {
		return nil
	}
	service := NewUserServiceInstance(dao)
	publicUser := service.GetPublicUser("214")
	assert.Nil(t, publicUser)
}

func TestWhenGetPublicUserAndUserExistThenUser(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*users_data_transfer.UserDb) {
		userDb := users_data_transfer.UserDb{
			Id: "214",
			FirstName: "Juan",
			LastName: "Perez",
			Version: 1,
			Email: "pepe@pepe.com",
		}
		return &userDb
	}
	service := NewUserServiceInstance(dao)
	privateUser := service.GetPublicUser("214")
	expectedPrivateUser := users_data_transfer.PublicUserDto{
		Id: "214",
		FirstName: "Juan",
		LastName: "Perez",
	}
	assert.Equal(t, expectedPrivateUser, *privateUser)
}