package users_domain

import (
	"github.com/Cristien/go-patterns/business/data_transfer/users_data_transfer"
	"github.com/google/uuid"
	"time"
)

func NewPrivateUserDto(user User) users_data_transfer.PrivateUserDto {
	return users_data_transfer.PrivateUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		FullName: user.GetFullName(),
		Email: user.Email,
		CreationDate: user.CreationDate,
	}
}

func NewPublicUserDto(user User) users_data_transfer.PublicUserDto {
	return users_data_transfer.PublicUserDto{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}


func NewUser(request users_data_transfer.CreateUserRequest) User{
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


func NewUserFromDbDto(userDbDto users_data_transfer.UserDb) User{
	return User{
		Id: userDbDto.Id,
		FirstName: userDbDto.FirstName,
		LastName: userDbDto.LastName,
		Email: userDbDto.Email,
		Password: userDbDto.Password,
		CreationDate: userDbDto.CreationDate,
		LastUpdateDate:userDbDto.LastUpdateDate,
	}
}

func NewUserDb(user User) users_data_transfer.UserDb{
	return users_data_transfer.UserDb{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
		CreationDate: user.CreationDate,
		LastUpdateDate: user.LastUpdateDate,
	}
}