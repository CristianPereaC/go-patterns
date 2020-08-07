package memory_dao

import (
	"errors"
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenUserDbHasInvalidDataWhenSaveUserThenInvalidDataError(t *testing.T) {
	userDbToPersist := users_data_transfer.UserDb{Id: "",}
	dao := NewUserDaoImplInstance()
	userDb, err := dao.SaveUser(userDbToPersist)
	assert.Nil(t, userDb)
	assert.Equal(t, errors.New("The value of the user are not valid"), err)
}

func TestWhenSaveUserForFirstTimeThenVersionIs1(t *testing.T) {
	userDbToPersist := users_data_transfer.UserDb{
		Id: "12345",
		FirstName: "Juan",
		LastName: "Perez",
	}
	dao := NewUserDaoImplInstance()
	userDb, err := dao.SaveUser(userDbToPersist)
	assert.Equal(t, 1, userDb.Version)
	assert.Nil(t, err)
}

func TestGivenUserIsPersistedWhenSaveUserThenCreationConflict(t *testing.T) {
	userDbToPersist := users_data_transfer.UserDb{
		Id: "12345",
		FirstName: "Juan",
		LastName: "Perez",
	}
	dao := NewUserDaoImplInstance()
	userDb, err := dao.SaveUser(userDbToPersist)
	userDb, err = dao.SaveUser(*userDb)
	assert.Nil(t, userDb)
	assert.Equal(t, errors.New("User is already persisted"), err)
}


func TestGivenUserWithId56IsNotPersistedWhenGetUserByIdThenNil(t *testing.T) {
	dao := NewUserDaoImplInstance()
	userDb := dao.GetUser("56")
	assert.Nil(t , userDb)
}

func TestGivenUserWithId56IsPersistedWhenGetUserByIdThenUserDbFetched(t *testing.T) {
	userId := "56"
	dao := NewUserDaoImplInstance()
	userDbToPersist := users_data_transfer.UserDb{
		Id: userId,
	}
	dao.SaveUser(userDbToPersist)
	userDb := dao.GetUser(userId)
	assert.Equal(t , userId, userDb.Id)
}