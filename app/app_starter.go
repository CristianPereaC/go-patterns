package app

var Context appContext

func StartApplication(){
	BuildAppContext()
	startAppRouter()
}

func BuildAppContext(){
	Context = newAppContext()
}

func startAppRouter(){
	Context.AppRouter.RunAppRouter()
}



