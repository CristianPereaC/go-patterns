package http_handlers

import (
	"github.com/Cristien/go-patterns/web/web_views"
	"github.com/Cristien/go-patterns/web/web_request"
)

/* Service definition */
type UserService interface {
	CreateUser(request web_request.CreateUserRequest) (id string, err error)
	GetPrivateUser(id string) (*web_views.PrivateUserDto)
	GetPublicUser(id string) (*web_views.PublicUserDto)
}

/*mock*/
type userServiceMock struct {
	HandleCreateUser     func() (id string, err error)
	HandleGetPrivateUser func() (*web_views.PrivateUserDto)
	HandleGetPublicUser  func() (*web_views.PublicUserDto)
}

func NewUserServiceMockInstance() userServiceMock {
	return userServiceMock{}
}

func (u userServiceMock) CreateUser(request web_request.CreateUserRequest) (id string, err error) {
	return u.HandleCreateUser()
}

func (u userServiceMock) GetPrivateUser(id string) (*web_views.PrivateUserDto) {
	return u.HandleGetPrivateUser()
}

func (u userServiceMock) GetPublicUser(id string) (*web_views.PublicUserDto) {
	return u.HandleGetPublicUser()
}
