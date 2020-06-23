package testing

import (
	"github.com/Cristien/go-patterns/app"
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
cada test podría levantar su contexto
cada test puede reemplazar una parte del contexto para mockear cierta parte del contexto
y hacer test de integración entre ciertas capas
 */

func TestCreateUserUserService(t *testing.T){

	service := app.Context.ServiceContext.UserService

	request := web_request.CreateUserRequest{
		FirstName:"Juan",
		LastName: "Perez",
		Email: "pepe@pepe.com",
		Password: "123456545234",
	}

	id, error := service.CreateUser(request)

	assert.Nil(t, error)

	assert.True(t, len(id)>0)

}