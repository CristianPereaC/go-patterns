package business_services

import (
	"github.com/Cristien/go-patterns/business/business_domain"
	"github.com/Cristien/go-patterns/business/business_dto"
)

const (
	DEFAULT_USER_ID = ""
)

/* Service definition */
type UserService interface {
	CreateUser(request business_dto.CreateUserRequest) (id string, err error)
	GetPrivateUser(id string) (*business_dto.PrivateUserDto)
	GetPublicUser(id string) (*business_dto.PublicUserDto)
}

type userServiceImpl struct {
	userDao UserDao
}

func NewUserServiceInstance(userDao UserDao) userServiceImpl {
	return userServiceImpl{
		userDao: userDao,
	}
}

func (u userServiceImpl) CreateUser(request business_dto.CreateUserRequest) (id string, err error) {
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

func (u userServiceImpl) GetPrivateUser(id string) (*business_dto.PrivateUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := business_domain.NewUserFromDbDto(*userDb)
	privateUser := business_domain.NewPrivateUserDto(user)
	return &privateUser

}

func (u userServiceImpl) GetPublicUser(id string) (*business_dto.PublicUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := business_domain.NewUserFromDbDto(*userDb)
	publicUser := business_domain.NewPublicUserDto(user)
	return &publicUser
}
