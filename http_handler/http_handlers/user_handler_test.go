package http_handlers

import (
	"encoding/json"
	"github.com/Cristien/go-patterns/business/business_dto"
	"github.com/Cristien/go-patterns/business/business_services"
	"github.com/Cristien/go-patterns/http_handler/http_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/*
Unit test
 */
func TestUserHandlerImpl_GetPublicUser(t *testing.T) {

	userServiceMock := business_services.NewUserServiceMockInstance()
	userServiceMock.HandleGetPublicUser = func() *business_dto.PublicUserDto {
		dto := business_dto.PublicUserDto{
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