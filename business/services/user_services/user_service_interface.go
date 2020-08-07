package user_services

import "github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"

/* Service definition */
type UserService interface {
	CreateUser(request users_data_transfer.CreateUserRequest) (id string, err error)
	GetPrivateUser(id string) (*users_data_transfer.PrivateUserDto)
	GetPublicUser(id string) (*users_data_transfer.PublicUserDto)
}

/*mock*/
type userServiceMock struct {
	HandleCreateUser     func() (id string, err error)
	HandleGetPrivateUser func() (*users_data_transfer.PrivateUserDto)
	HandleGetPublicUser  func() (*users_data_transfer.PublicUserDto)
}

func NewUserServiceMockInstance() userServiceMock {
	return userServiceMock{}
}

func (u userServiceMock) CreateUser(request users_data_transfer.CreateUserRequest) (id string, err error) {
	return u.HandleCreateUser()
}

func (u userServiceMock) GetPrivateUser(id string) (*users_data_transfer.PrivateUserDto) {
	return u.HandleGetPrivateUser()
}

func (u userServiceMock) GetPublicUser(id string) (*users_data_transfer.PublicUserDto) {
	return u.HandleGetPublicUser()
}
