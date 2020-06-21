package http_handlers

import (
	"encoding/json"
	"github.com/Cristien/go-patterns/web/http_utils"
	"github.com/Cristien/go-patterns/web/web_views"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/*
Unit test
 */
func TestUserHandlerImpl_GetPublicUser(t *testing.T) {

	userServiceMock := NewUserServiceMockInstance()
	userServiceMock.HandleGetPublicUser = func() *web_views.PublicUserDto {
		dto := web_views.PublicUserDto{
			Id: "1234",
			FirstName: "Juan",
			LastName: "Perez",
		}
		return &dto
	}
	handler := NewuserHandler(userServiceMock)

	context, response := http_utils.BuildGetHttpRequestContext("/users/1234", nil)

	handler.GetPublicUser(context)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResponse := map[string]string{
		"id": "1234",
		"first_name": "Juan",
		"last_name": "Perez",
	}

	apiResponse := map[string]string{}

	json.Unmarshal(response.Body.Bytes(), &apiResponse)

	assert.Equal(t, expectedResponse, apiResponse)
}