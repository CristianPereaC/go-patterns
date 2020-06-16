package service

import "fmt"
import "../dao"

/* Dependencies */
type ctx struct {
	UsersDao dao.UserDaoInterface
}

var UserService UserServiceInstance

/* Instance build */
func init() {
	UserService = NewUserServiceInstance()
}

func NewUserServiceInstance() UserServiceInstance {
	return UserServiceInstance{
		ctx{
			UsersDao: dao.UserDao,
		},
	}
}

/* Service definition */
type UserServiceInterface interface {
	CreateUser()
	UpdateUser()
}

type UserServiceInstance struct {
	Ctx ctx
}

func (u UserServiceInstance) CreateUser() {
	u.Ctx.UsersDao.CreateUserDB()
	fmt.Println("service done")
}

func (u UserServiceInstance) UpdateUser() {
	u.Ctx.UsersDao.UpdateUserDB()
	fmt.Println("service done")
}
