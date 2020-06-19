package routing

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetPrivateUser(c *gin.Context)
	GetPublicUser(c *gin.Context)
}
