package main

type App struct {
	ambiente Ambiente
}

func (a *App) Init(nPresas, nPredadores int) {
	a.ambiente = Ambiente{}
	a.ambiente.Init(nPresas, nPredadores)
}

func (a *App) Run() {
	a.ambiente.Run()
}