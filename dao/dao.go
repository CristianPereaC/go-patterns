package dao

import "fmt"

/* Dependencies */
type ctx struct{}

var UserDaoInstance UserDaoImpl

/* Instance build */
func init() {
	UserDaoInstance = NewUserDaoInstance()
}

func NewUserDaoInstance() UserDaoImpl {
	return UserDaoImpl{
		ctx{},
	}
}

/* Dao definition */
type UserDao interface {
	CreateUserDB()
	UpdateUserDB()
}

type UserDaoImpl struct {
	Ctx ctx
}

func (u UserDaoImpl) CreateUserDB() {
	fmt.Println("DB USER CREATED")
}
func (u UserDaoImpl) UpdateUserDB() {
	fmt.Println("DB USER UPDATED")
}
