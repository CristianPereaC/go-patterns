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

	if _, err := u.userService.CreateUser(*request); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusCreated)
}
func (u userHandlerImpl) GetPrivateUser(c *gin.Context){
	id := c.Param("id")


}
func (u userHandlerImpl) GetPublicUser(c *gin.Context){

}