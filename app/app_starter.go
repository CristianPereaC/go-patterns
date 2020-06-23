package app

var Context appContext

func StartApplication(){
	BuildAppContext()
	Context.AppRouter.RunAppRouter()
}

func BuildAppContext(){
	Context = newAppContext()
}



