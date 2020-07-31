package business_services

import (
	"errors"
	"github.com/Cristien/go-patterns/business/business_dto"
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/Cristien/go-patterns/web/web_views"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCreateUserAndDaoFailsThenError(t *testing.T) {
	dao := newUserDaoMockInstance()
	dao.HandleSaveUser = func() (*business_dto.UserDb, error) {
		return nil, errors.New("error creating user")
	}
	service := NewUserServiceInstance(dao)
	request := web_request.CreateUserRequest{
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
	dao.HandleSaveUser = func() (*business_dto.UserDb, error) {
		userDb := business_dto.UserDb{
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
	request := web_request.CreateUserRequest{
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
	dao.HandleGetUser = func() (*business_dto.UserDb) {
		return nil
	}
	service := NewUserServiceInstance(dao)
	privateUser := service.GetPrivateUser("214")
	assert.Nil(t, privateUser)
}

func TestWhenGetPrivateUserAndUserExistThenUser(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*business_dto.UserDb) {
		userDb := business_dto.UserDb{
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
	expectedPrivateUser := web_views.PrivateUserDto{
		Id: "214",
		FirstName: "Juan",
		LastName: "Perez",
		Email: "pepe@pepe.com",
	}
	assert.Equal(t, expectedPrivateUser, *privateUser)
}

func TestWhenGetPublicUserAndUserNotExistThenNil(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*business_dto.UserDb) {
		return nil
	}
	service := NewUserServiceInstance(dao)
	publicUser := service.GetPublicUser("214")
	assert.Nil(t, publicUser)
}

func TestWhenGetPublicUserAndUserExistThenUser(t *testing.T){
	dao := newUserDaoMockInstance()
	dao.HandleGetUser = func() (*business_dto.UserDb) {
		userDb := business_dto.UserDb{
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
	expectedPrivateUser := web_views.PublicUserDto{
		Id: "214",
		FirstName: "Juan",
		LastName: "Perez",
	}
	assert.Equal(t, expectedPrivateUser, *privateUser)
}