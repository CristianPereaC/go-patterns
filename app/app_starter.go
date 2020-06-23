package app

var Context appContext

func StartApplication(){
	BuildAppContext()
	startAppRouter()
}

func BuildAppContext(){
	Context = NewAppContext()
}

func startAppRouter(){
	Context.AppRouter.RunAppRouter()
}



