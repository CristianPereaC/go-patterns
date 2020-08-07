package user_services

import "github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"

type UserDao interface {
	SaveUser(userDb users_data_transfer.UserDb) (*users_data_transfer.UserDb, error)
	GetUser(id string) (*users_data_transfer.UserDb)
}

/*Mock definition*/
type UserDaoMock struct {
	HandleSaveUser func() (*users_data_transfer.UserDb, error)
	HandleGetUser func() (*users_data_transfer.UserDb)
}

/*constructor*/
func newUserDaoMockInstance() UserDaoMock {
	return UserDaoMock{}
}

func (u UserDaoMock) SaveUser(userDb users_data_transfer.UserDb) (*users_data_transfer.UserDb, error){
	return u.HandleSaveUser()
}
func (u UserDaoMock) GetUser(id string) (*users_data_transfer.UserDb){
	return u.HandleGetUser()
}
