package business_domain

import (
	"github.com/Cristien/go-patterns/web/web_views"
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/Cristien/go-patterns/business_layer/business_dto"
	"github.com/google/uuid"
	"time"
)

//TODO separar esto implementando el patron builder

func NewPrivateUserDto(user User) web_views.PrivateUserDto {
	return web_views.PrivateUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		CreationDate: user.CreationDate,
	}
}

func NewPublicUserDto(user User) web_views.PublicUserDto {
	return web_views.PublicUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}


func NewUser(request web_request.CreateUserRequest) User{
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