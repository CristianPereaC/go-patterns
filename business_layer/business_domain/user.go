package business_domain

/**
Domain no tiene JSON value ya que nunca debería salir de la capa de negocio hacia otras capas
solo mediante dtos

https://medium.com/mastering-software-engineering/data-access-patterns-the-features-of-the-main-data-access-patterns-applied-in-software-industry-6eff86906b4e
**/

type User struct {
	Id string
	FirstName string
	LastName string
	FullName string
	Email string
	CreationDate string
	LastUpdateDate string
	Password string
}