package testing

import (
	"fmt"
	"github.com/Cristien/go-patterns/app"
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserFunctional(t *testing.T){

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

	req, response := buildGetRequest(fmt.Sprintf("/users/%s/private", id), nil)

	app.Context.AppRouter.ServeHTTP(req, response)

	assert.Equal(t, http.StatusOK, response.Code)

	fmt.Println(response.Body.String())

}


func buildGetRequest(url string, headers map[string]string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest("GET", url, nil)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response := httptest.NewRecorder()

	return request, response
}
