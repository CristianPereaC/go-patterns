package app

import (
	"github.com/Cristien/go-patterns/business/business_services"
	"github.com/Cristien/go-patterns/dao/memory_dao"
	"github.com/Cristien/go-patterns/routing"
	"github.com/Cristien/go-patterns/web/http_handlers"
)

type appContext struct {
	DaoContext     *daoContext
	ServiceContext *serviceContext
	HandlerContext *handlerContext
	AppRouter      *routing.AppRouter
}

//todo stub para ejecutar local la api y hacer pruebas
func NewAppContext() appContext{
	appContext := new(appContext)
	appContext.buildDaos()
	appContext.buildServices()
	appContext.buildHandlers()
	appContext.buildAppRouter()
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

func (c *appContext) buildAppRouter(){
	c.AppRouter = routing.NewAppRouter()
	c.AppRouter.MapUserResourceHandlers(c.HandlerContext.UserHandler)
}




