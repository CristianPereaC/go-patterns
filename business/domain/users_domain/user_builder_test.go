package users_domain

import (
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPrivateUserDto(t *testing.T) {
	user := User{
		Id: "123",
		LastName: "Perez",
		FirstName: "Juan",
		Email: "juan@juan.com",
		CreationDate: "12-12-2020",
		LastUpdateDate: "12-12-2020",
		Password: "1234",
	}
	expectedPrivateUser := users_data_transfer.PrivateUserDto{
		Id: "123",
		FirstName: "Juan",
		LastName:  "Perez",
		Email: "juan@juan.com",
		CreationDate: "12-12-2020",
		FullName: "Perez, Juan",
	}

	privateUserDto := NewPrivateUserDto(user)

	assert.Equal(t, expectedPrivateUser, privateUserDto)
}

func TestNewPublicUserDtoUserDto(t *testing.T) {
	user := User{
		Id: "123",
		LastName: "Perez",
		FirstName: "Juan",
		Email: "juan@juan.com",
		CreationDate: "12-12-2020",
		LastUpdateDate: "12-12-2020",
		Password: "1234",
	}
	expectedPublicUser := users_data_transfer.PublicUserDto{
		Id: "123",
		FirstName: "Juan",
		LastName:  "Perez",
	}

	publicUserDto := NewPublicUserDto(user)

	assert.Equal(t, expectedPublicUser, publicUserDto)
}


func TestNewUser(t *testing.T) {

	createuserRequest := users_data_transfer.CreateUserRequest{
		FirstName: "Juan",
		LastName: "Perez",
		Email: "juan@perez@gmail.com",
		Password: "123425645343",
	}

	user := NewUser(createuserRequest)

	assert.Equal(t, "Juan", user.FirstName)
	assert.Equal(t, "Perez", user.LastName)
	assert.Equal(t, "juan@perez@gmail.com", user.Email)
	assert.Equal(t, "123425645343", user.Password)
}


func TestNewUserFromDbDto(t *testing.T) {

	userDb := users_data_transfer.UserDb{
		Id: "123425364",
		FirstName: "Juan",
		LastName: "Perez",
		Email: "juan@perez@gmail.com",
		Password: "123425645343",
		Version: 3,
	}

	user := NewUserFromDbDto(userDb)

	assert.Equal(t, "123425364", user.Id)
	assert.Equal(t, "Juan", user.FirstName)
	assert.Equal(t, "Perez", user.LastName)
	assert.Equal(t, "juan@perez@gmail.com", user.Email)
	assert.Equal(t, "123425645343", user.Password)
}

func TestNewUserDb(t *testing.T) {

	user := User{
		Id: "123425364",
		FirstName: "Juan",
		LastName: "Perez",
		Email: "juan@perez@gmail.com",
		Password: "123425645343",
	}

	userDb := NewUserDb(user)

	assert.Equal(t, "123425364", userDb.Id)
	assert.Equal(t, "Juan", userDb.FirstName)
	assert.Equal(t, "Perez", userDb.LastName)
	assert.Equal(t, "juan@perez@gmail.com", userDb.Email)
	assert.Equal(t, "123425645343", userDb.Password)
	assert.Equal(t, 0, userDb.Version)
}
