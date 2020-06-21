package http_handlers

import (
	"github.com/Cristien/go-patterns/web/web_request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandlerImpl struct {
	userService UserService
}

func NewuserHandler(userService UserService) userHandlerImpl {
	return userHandlerImpl{
		userService: userService,
	}
}

func (u userHandlerImpl) CreateUser(c *gin.Context){
	request := new(web_request.CreateUserRequest)
	err := c.BindJSON(request)

	if err != nil || !request.IsValid(){
		c.JSON(http.StatusBadRequest, err)
	}

	userId, err := u.userService.CreateUser(*request);

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	id := struct {
		Id string `json:"id"`
	}{
		Id: userId,
	}

	c.JSON(http.StatusCreated, id)
}
func (u userHandlerImpl) GetPrivateUser(c *gin.Context){
	id := c.Param("id")

	if len(id) > 0 {
		c.Status(http.StatusBadRequest)
	}

	if dto := u.userService.GetPrivateUser(id); dto != nil {
		c.JSON(http.StatusOK, dto)
	}

	c.Status(http.StatusNotFound)

}
func (u userHandlerImpl) GetPublicUser(c *gin.Context){
	id := c.Param("id")

	if len(id) > 0 {
		c.Status(http.StatusBadRequest)
	}

	if dto := u.userService.GetPublicUser(id); dto != nil {
		c.JSON(http.StatusOK, dto)
	}

	c.Status(http.StatusNotFound)
}