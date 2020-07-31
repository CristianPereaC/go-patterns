package business_services

import (
	"github.com/Cristien/go-patterns/business/business_dto"
)

type UserDao interface {
	SaveUser(userDb business_dto.UserDb) (*business_dto.UserDb, error)
	GetUser(id string) (*business_dto.UserDb)
}

/*Mock definition*/
type UserDaoMock struct {
	HandleSaveUser func() (*business_dto.UserDb, error)
	HandleGetUser func() (*business_dto.UserDb)
}

/*constructor*/
func newUserDaoMockInstance() UserDaoMock {
	return UserDaoMock{}
}

func (u UserDaoMock) SaveUser(userDb business_dto.UserDb) (*business_dto.UserDb, error){
	return u.HandleSaveUser()
}
func (u UserDaoMock) GetUser(id string) (*business_dto.UserDb){
	return u.HandleGetUser()
}
