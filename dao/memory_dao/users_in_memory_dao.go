package memory_dao

import (
	"errors"
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
)

const(
	CREATION_CONFLICT = "User is already persisted"
	INVALID_USER_DB_DATA = "The value of the user are not valid"
	FIRST_VERSION = 1
)

/* Dao definition */
type userDaoMemoryImpl struct {
	container map[string]users_data_transfer.UserDb
}

/*constructor*/
func NewUserDaoImplInstance() userDaoMemoryImpl {
	return userDaoMemoryImpl{
		container: map[string]users_data_transfer.UserDb{},
	}
}

func (u userDaoMemoryImpl) SaveUser(userDb users_data_transfer.UserDb) (*users_data_transfer.UserDb, error){
	if !userDb.IsValid() {
		return nil, errors.New(INVALID_USER_DB_DATA)
	}
	if _, ok := u.container[userDb.Id]; ok{
		return nil, errors.New(CREATION_CONFLICT)
	}
	userDb.Version = FIRST_VERSION
	u.container[userDb.Id] = userDb
	return &userDb, nil
}

func (u userDaoMemoryImpl) GetUser(id string) (*users_data_transfer.UserDb){
	if persistedUderDb, ok := u.container[id]; ok{
		return &persistedUderDb
	}
	return nil
}