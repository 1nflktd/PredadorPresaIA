package main

import (
	"time"
	"math/rand"
)

type Agente interface {
	setPosicao(p Posicao)
	setPosicaoXY(x, y int)
	getPosicao() Posicao
	movePosAleatorio() Posicao
	getCAgente() CAgente
}

type AgenteImpl struct {
	posicao Posicao
	cAgente CAgente
}

func (a *AgenteImpl) Init() {}

func (a *AgenteImpl) setPosicao(p Posicao) {
	a.posicao = p
}

func (a *AgenteImpl) setPosicaoXY(x, y int) {
	a.posicao = Posicao{x, y}
}

func (a *AgenteImpl) getPosicao() Posicao {
	return a.posicao
}

// Nao muda o estado do objeto
func (a *AgenteImpl) movePosAleatorio() Posicao {
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
