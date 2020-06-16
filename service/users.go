package service

import "fmt"
import "../dao"

/* Dependencies */
type ctx struct {
	UsersDao dao.UserDao
}

var UserServiceInstance UserServiceImpl

/* Instance build */
func init() {
	UserServiceInstance = NewUserServiceInstance()
}

func NewUserServiceInstance() UserServiceImpl {
	return UserServiceImpl{
		ctx{
			UsersDao: dao.UserDaoInstance,
		},
	}
}

/* Service definition */
type UserService interface {
	CreateUser()
	UpdateUser()
}

type UserServiceImpl struct {
	Ctx ctx
}

func (u UserServiceImpl) CreateUser() {
	u.Ctx.UsersDao.CreateUserDB()
	fmt.Println("service done")
}

func (u UserServiceImpl) UpdateUser() {
	u.Ctx.UsersDao.UpdateUserDB()
	fmt.Println("service done")
}
