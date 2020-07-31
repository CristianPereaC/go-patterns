package business_services

import (
	"github.com/Cristien/go-patterns/business/business_domain"
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/Cristien/go-patterns/web/web_views"
)

const (
	DEFAULT_USER_ID = ""
)

type userServiceImpl struct {
	userDao UserDao
}

func NewUserServiceInstance(userDao UserDao) userServiceImpl {
	return userServiceImpl{
		userDao: userDao,
	}
}

func (u userServiceImpl) CreateUser(request web_request.CreateUserRequest) (id string, err error) {
	//Generar objeto de negocio - obtención de id
	user := business_domain.NewUser(request)
	//convertir a dto
	userDbDto := business_domain.NewUserDb(user)
	//persistir
	userDb, err := u.userDao.SaveUser(userDbDto)

	if err != nil {
		return DEFAULT_USER_ID, err
	}
	return userDb.Id, err

}

func (u userServiceImpl) GetPrivateUser(id string) (*web_views.PrivateUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := business_domain.NewUserFromDbDto(*userDb)
	privateUser := business_domain.NewPrivateUserDto(user)
	return &privateUser

}

func (u userServiceImpl) GetPublicUser(id string) (*web_views.PublicUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := business_domain.NewUserFromDbDto(*userDb)
	publicUser := business_domain.NewPublicUserDto(user)
	return &publicUser
}
