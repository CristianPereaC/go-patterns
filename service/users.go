package service

import "fmt"
import "../dao"

type options struct {
	UsersDao dao.UserDaoImpl
}

func GetUserServiceInstance() UserServiceImpl {

	var dependencies = options{
		UsersDao: dao.GetUserDaoInstance(),
	}

	return UserServiceImpl{
		CreateUser,
		UpdateUser,
		dependencies,
	}
}

type UserServiceImpl struct {
	CreateUser func(options)
	UpdateUser func(options)
	Options    options
}

func CreateUser(ops options) {
	ops.UsersDao.CreateUserDB()
	fmt.Println("service USER CREATED")
}

func UpdateUser(ops options) {
	ops.UsersDao.UpdateUserDB()
	fmt.Println("service USER UPDATED")
}
