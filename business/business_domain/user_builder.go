package business_domain

import (
	"github.com/Cristien/go-patterns/business/business_dto"
	"github.com/google/uuid"
	"time"
)

//TODO separar esto implementando el patron builder

func NewPrivateUserDto(user User) business_dto.PrivateUserDto {
	return business_dto.PrivateUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		CreationDate: user.CreationDate,
	}
}

func NewPublicUserDto(user User) business_dto.PublicUserDto {
	return business_dto.PublicUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}


func NewUser(request business_dto.CreateUserRequest) User{
	currentTime := time.Now().String()
	return User{
		Id: uuid.New().String(),
		FirstName: request.FirstName,
		LastName: request.LastName,
		Email: request.Email,
		Password: request.Password,
		CreationDate: currentTime,
		LastUpdateDate:currentTime,
	}
}


func NewUserFromDbDto(userDbDto business_dto.UserDb) User{
	return User{
		Id: userDbDto.Id,
		FirstName: userDbDto.FirstName,
		LastName: userDbDto.LastName,
		Email: userDbDto.Email,
		Password: userDbDto.Password,
		CreationDate: userDbDto.CreationDate,
		LastUpdateDate:userDbDto.LastUpdateDate,
		FullName: userDbDto.FirstName + userDbDto.LastName,
	}
}

func NewUserDb(user User) business_dto.UserDb{
	return business_dto.UserDb{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
		CreationDate: user.CreationDate,
		LastUpdateDate: user.LastUpdateDate,
	}
}