package main

import (
	"time"
	"math/rand"
)

type Agente struct {
	posicao Posicao
	base Posicao
	temDiamante bool
	caracter Caracter
}

func (a *Agente) Init(base Posicao, caracter Caracter) {
	a.base = base
	a.caracter = caracter
}

func (a *Agente) setPosicao(p Posicao) {
	a.posicao = p
}

func (a *Agente) setPosicaoXY(x, y int) {
	a.posicao = Posicao{x, y}
}

func (a *Agente) getPosicao() Posicao {
	return a.posicao
}

func (a *Agente) setTemDiamante(temDiamante bool) {
	a.temDiamante = temDiamante
}

func (a *Agente) getTemDiamante() bool {
	return a.temDiamante
}

func (a *Agente) getCaracter() Caracter {
	return a.caracter
}

// Nao muda o estado do objeto
func (a *Agente) movePosAleatorio() Posicao {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pos := a.posicao

	if pos.X >= (TamanhoMapa - 1) {
		pos.X--
	} else if pos.X == 0 {
		pos.X++
	} else {
		pos.X += (r.Intn(3) - 1) // random de 0 a 2, se 0 volta uma (-1), 1 fica parado, 2 avanca
	}

	if pos.Y >= (TamanhoMapa - 1) {
		pos.Y--
	} else if pos.Y == 0 {
		pos.Y++
	} else {
		pos.Y += (r.Intn(3) - 1) // random de 0 a 2, se 0 volta uma (-1), 1 fica parado, 2 avanca
	}

	return pos
}

func (a *Agente) voltaBase() Posicao {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pos := a.posicao

	if pos.X >= (TamanhoMapa - 1) {
		pos.X--
	} else if pos.X == 0 || r.Float64() <= 0.2 { // 20% dos casos avanca
		pos.X++
	} else {
		pos.X += (r.Intn(2) - 1) // random de 0 a 1, se 0 volta uma (-1), 1 fica parado
	}

	if pos.Y >= (TamanhoMapa - 1) {
		pos.Y--
	} else if pos.Y == 0 || r.Float64() <= 0.2 { // 20% dos casos avanca
		pos.Y++
	} else {
		pos.Y += (r.Intn(2) - 1) // random de 0 a 1, se 0 volta uma (-1), 1 fica parado
	}

	return pos
}
