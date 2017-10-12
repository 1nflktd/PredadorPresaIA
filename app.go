package main

type App struct {
	ambiente Ambiente
}

func (a *App) Run(nPresas, nPredadores int) {
	a.ambiente = Ambiente{}
	a.ambiente.Init(nPresas, nPredadores)
	a.ambiente.Run()
}