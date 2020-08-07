package users_domain

import "fmt"

/**
Domain no tiene JSON value ya que nunca debería salir de la capa de negocio hacia otras capas
solo mediante dtos

https://medium.com/mastering-software-engineering/data-access-patterns-the-features-of-the-main-data-access-patterns-applied-in-software-industry-6eff86906b4e
**/

type User struct {
	Id string
	FirstName string
	LastName string
	Email string
	CreationDate string
	LastUpdateDate string
	Password string
}

func (u User) GetFullName() string{
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}