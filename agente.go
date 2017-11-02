package main

import (
	// "time"
	"math/rand"
	"sync"
)

type Agente interface {
	Init(int)
	setRandPos(*rand.Rand, *sync.Mutex)
	setPosicao(Posicao)
	setPosicaoXY(int, int)
	getPosicao() Posicao
	viver() (Posicao, PosMovimento)
	getCAgente() CAgente
	getId() int
	mover(CampoVisao) (Posicao, PosMovimento)
	moveAgente(Direcao, int) (Posicao, PosMovimento)
}

type AgenteImpl struct {
	id int
	posicao Posicao
	cAgente CAgente
	randPos *rand.Rand
	mutexRandPos *sync.Mutex
}

type PosMovimento int
const (
	AUM_X PosMovimento = iota
	DIM_X
	AUM_Y
	DIM_Y
)

func (a *AgenteImpl) Init(id int) {}

func (a *AgenteImpl) setRandPos(randPos *rand.Rand, mutexRandPos *sync.Mutex) {
	a.mutexRandPos = mutexRandPos
	a.randPos = randPos
}

func (a *AgenteImpl) setPosicao(p Posicao) {
	a.posicao = p
}

func (a *AgenteImpl) setPosicaoXY(x, y int) {
	a.posicao = Posicao{x, y}
}

func (a *AgenteImpl) getPosicao() Posicao {
	return a.posicao
}

func (a *AgenteImpl) getId() int {
	return a.id
}

// Nao muda o estado do objeto
func (a *AgenteImpl) viver() (Posicao, PosMovimento) {
	return a.moveAgente(P_Aleatoria, 1)
}

func (a *AgenteImpl) moveAgente(direcao Direcao, velocidade int) (Posicao, PosMovimento) {
	pos := a.posicao

	var x_y, pos_neg_para int

	a.mutexRandPos.Lock()

	// verifica se move x ou y
	x_y = a.randPos.Intn(1000) % 2

	// verifica se positivo ou negativo
	pos_neg_para = (a.randPos.Intn(1002) % 3) - 1

	a.mutexRandPos.Unlock()

	if direcao != P_Aleatoria {
		switch (direcao) {
		case P_Norte:
			if x_y == 0 {
				pos_neg_para *= velocidade
			} else {
				pos_neg_para = velocidade
			}
		case P_Noroeste:
			if x_y == 0 {
				pos_neg_para = -velocidade
			} else {
				pos_neg_para = velocidade
			}
		case P_Nordeste:
			pos_neg_para = velocidade
		case P_Leste:
			if x_y == 0 {
				pos_neg_para = velocidade
			} else {
				pos_neg_para *= velocidade
			}
		case P_Oeste:
			if x_y == 0 {
				pos_neg_para = -velocidade
			} else {
				pos_neg_para *= velocidade
			}
		case P_Sudoeste:
			pos_neg_para = -velocidade
		case P_Sul:
			if x_y == 0 {
				pos_neg_para *= velocidade
			} else {
				pos_neg_para = -velocidade
			}
		}
	}

	if x_y == 0 { // move x
		pos.X += pos_neg_para
		pos.X = VerificaLimites(pos.X)
	} else { // move y
		pos.Y += pos_neg_para
		pos.Y = VerificaLimites(pos.Y)
	}

	return pos, VerificaPosMovimento(x_y, pos_neg_para)
}

func VerificaPosMovimento(x_y, pos_neg int) PosMovimento {
	if x_y == 0 {
		if pos_neg > 0 {
			return AUM_X
		} else {
			return DIM_X
		}
	}

	if pos_neg > 0 {
		return AUM_Y
	}

	return DIM_Y
}
