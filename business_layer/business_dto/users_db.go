package business_dto

import "github.com/Cristien/go-patterns/business_layer/business_domain"

type UserDb struct {
	Id string
	FirstName string
	LastName string
	Email string
	CreationDate string
	LastUpdateDate string
	Password string
	Version int
}

func (u UserDb) IsValid() bool{
	return len(u.Id) > 0;
}


func NewUserDb(user business_domain.User) UserDb{
	return UserDb{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
		CreationDate: user.CreationDate,
		LastUpdateDate: user.LastUpdateDate,
	}
}