package routing


type apiHandlers struct {
	userHandler *UserHandler
}

func NewApiHandlers() apiHandlers{
	return apiHandlers{}
}

func (apiHandlers *apiHandlers) SetUserHandler(userHandler *UserHandler){
	apiHandlers.userHandler = userHandler
}