package main

import (
	"net/http"
)

type App struct {
	ambiente Ambiente
}

func (a *App) Run(w http.ResponseWriter, nDiamantes, nPedras, nAgentes int) {
	a.ambiente = Ambiente{}
	a.ambiente.Init(w, nDiamantes, nPedras, nAgentes)
	a.ambiente.Run()
}