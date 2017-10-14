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

	// verifica se move x ou y
	x_y := r.Intn(2)

	// verifica se positivo ou negativo
	pos_neg_para := r.Intn(3) - 1

	if x_y == 0 { // move x
		pos.X += pos_neg_para

		if pos.X >= TamanhoMapa {
			pos.X = 0
		} else if pos.X < 0 {
			pos.X = TamanhoMapa - 1
		}

	} else { // move y
		pos.Y += pos_neg_para

		if pos.Y >= TamanhoMapa {
			pos.Y = 0
		} else if pos.Y < 0 {
			pos.Y = TamanhoMapa - 1
		}
	}

	return pos
}
