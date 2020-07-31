package business_dto

import "strings"

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (request CreateUserRequest) IsValid() bool{
	validFirstName := len(request.FirstName) > 0
	validLastName := len(request.LastName) > 0
	validPassword := len(request.Password) > 8
	atIndex := strings.LastIndex(request.Email, "@")
	validEmail := len(request.Email) > 10 &&
		len(request.Email) <= 100 &&
		strings.Contains(request.Email, "@") &&
		atIndex > 0 && atIndex < len(request.Email)

	return validFirstName && validLastName && validPassword && validEmail
}