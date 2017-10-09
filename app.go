package main

import (
	"net/http"
)

type App struct {
	ambiente Ambiente
}

func (a *App) Run(w http.ResponseWriter, nPresas, nPredadores int) {
	a.ambiente = Ambiente{}
	a.ambiente.Init(w, nPresas, nPredadores)
	a.ambiente.Run()
}