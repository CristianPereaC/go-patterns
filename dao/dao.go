package dao

import "fmt"

/* Dependencies */
type ctx struct{}

var UserDao UserDaoInterface

/* Instance build */
func init() {
	UserDao = NewUserDaoInstance()
}

func NewUserDaoInstance() UserDaoInstance {
	return UserDaoInstance{
		ctx{},
	}
}

/* Dao definition */
type UserDaoInterface interface {
	CreateUserDB()
	UpdateUserDB()
}

type UserDaoInstance struct {
	Ctx ctx
}

func (u UserDaoInstance) CreateUserDB() {
	fmt.Println("DB USER CREATED")
}
func (u UserDaoInstance) UpdateUserDB() {
	fmt.Println("DB USER UPDATED")
}
