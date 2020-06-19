package memory_dao

import (
	"errors"
	"github.com/Cristien/go-patterns/business_layer/business_dto"
)

const(
	INVALID_USER_DB_DATA = "The value of the user are not valid"
	FIRST_VERSION = 1
)

/* Dao definition */
type userDaoMemoryImpl struct {
	container map[string]business_dto.UserDb
}

/*constructor*/
func NewUserDaoImplInstance() userDaoMemoryImpl {
	return userDaoMemoryImpl{
		container: map[string]business_dto.UserDb{},
	}
}

func (u userDaoMemoryImpl) SaveUser(userDb business_dto.UserDb) (*business_dto.UserDb, error){
	if !userDb.IsValid() {
		return nil, errors.New(INVALID_USER_DB_DATA)
	}
	if persistedUderDb, ok := u.container[userDb.Id]; ok{
		userDb.Version = persistedUderDb.Version + 1
	}else{
		userDb.Version = FIRST_VERSION
	}
	u.container[userDb.Id] = userDb
	return &userDb, nil
}

func (u userDaoMemoryImpl) GetUser(id string) (*business_dto.UserDb){
	if persistedUderDb, ok := u.container[id]; ok{
		return &persistedUderDb
	}
	return nil

}