package app_context

import (
	"github.com/Cristien/go-patterns/business_layer/business_services"
	"github.com/Cristien/go-patterns/dao/memory_dao"
	"github.com/Cristien/go-patterns/routing"
	"github.com/Cristien/go-patterns/web/http_handlers"
)

type appContext struct{
	DaoContext     *daoContext
	ServiceContext *serviceContext
	HandlerContext *handlerContext
}

func BuildAppContext() appContext{
	appContext := new(appContext)
	appContext.buildDaos()
	appContext.buildServices()
	appContext.buildHandlers()
	return *appContext
}

type daoContext struct {
	UsersDao business_services.UserDao
}

type serviceContext struct {
	UserService http_handlers.UserService
}

type handlerContext struct {
	UserHandler routing.UserHandler
}

func (c *appContext) buildDaos(){
	c.DaoContext = new(daoContext)
	c.DaoContext.UsersDao = memory_dao.NewUserDaoImplInstance()
}

func (c *appContext) buildServices(){
	c.ServiceContext = new(serviceContext)
	c.ServiceContext.UserService = business_services.NewUserServiceInstance(c.DaoContext.UsersDao)
}

func (c *appContext) buildHandlers(){
	c.HandlerContext = new(handlerContext)
	c.HandlerContext.UserHandler = http_handlers.NewuserHandler(c.ServiceContext.UserService)

}




