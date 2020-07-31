package business_services

import "github.com/Cristien/go-patterns/business/business_dto"

/*mock*/
type userServiceMock struct {
	HandleCreateUser     func() (id string, err error)
	HandleGetPrivateUser func() (*business_dto.PrivateUserDto)
	HandleGetPublicUser  func() (*business_dto.PublicUserDto)
}

func NewUserServiceMockInstance() userServiceMock {
	return userServiceMock{}
}

func (u userServiceMock) CreateUser(request business_dto.CreateUserRequest) (id string, err error) {
	return u.HandleCreateUser()
}

func (u userServiceMock) GetPrivateUser(id string) (*business_dto.PrivateUserDto) {
	return u.HandleGetPrivateUser()
}

func (u userServiceMock) GetPublicUser(id string) (*business_dto.PublicUserDto) {
	return u.HandleGetPublicUser()
}
