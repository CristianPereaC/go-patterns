package user_services

import (
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
	"github.com/Cristien/go-patterns/business/domain/users_domain"
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

func (u userServiceImpl) CreateUser(request users_data_transfer.CreateUserRequest) (id string, err error) {
	//Generar objeto de negocio - obtención de id
	user := users_domain.NewUser(request)
	//convertir a db dto
	userDbDto := users_domain.NewUserDb(user)
	//persistir
	userDb, err := u.userDao.SaveUser(userDbDto)

	if err != nil {
		return DEFAULT_USER_ID, err
	}
	return userDb.Id, err
}

func (u userServiceImpl) GetPrivateUser(id string) (*users_data_transfer.PrivateUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := users_domain.NewUserFromDbDto(*userDb)
	privateUser := users_domain.NewPrivateUserDto(user)
	return &privateUser

}

func (u userServiceImpl) GetPublicUser(id string) (*users_data_transfer.PublicUserDto) {
	userDb := u.userDao.GetUser(id)
	if userDb == nil {
		return nil
	}
	user := users_domain.NewUserFromDbDto(*userDb)
	publicUser := users_domain.NewPublicUserDto(user)
	return &publicUser
}
