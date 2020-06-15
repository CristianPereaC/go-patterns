package dao

import "fmt"

func GetUserDaoInstance() UserDaoImpl {
	return UserDaoImpl{
		CreateUserDB,
		UpdateUserDB,
	}
}

type UserDaoImpl struct {
	CreateUserDB func()
	UpdateUserDB func()
}

func CreateUserDB() {
	fmt.Println("DB USER CREATED")
}
func UpdateUserDB() {
	fmt.Println("DB USER UPDATED")
}
