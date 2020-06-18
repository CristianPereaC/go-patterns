package http_handlers


type userHandler struct {
	userService UserService
}

func NewuserHandler(userService UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}