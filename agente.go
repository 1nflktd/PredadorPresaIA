package main

import (
	"time"
	"math/rand"
)

type Agente interface {
	setPosicao(p Posicao)
	setPosicaoXY(x, y int)
	getPosicao() Posicao
	viver() Posicao
	getCAgente() CAgente
	mover(CampoVisao) Posicao
	moveAgente(Direcao, int) Posicao
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
func (a *AgenteImpl) viver() Posicao {
	return a.moveAgente(P_Aleatoria, 1)
}

func (a *AgenteImpl) moveAgente(direcao Direcao, velocidade int) Posicao {
	pos := a.posicao

	var x_y, pos_neg_para int

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// verifica se move x ou y
	x_y = r.Intn(2)

	// verifica se positivo ou negativo
	pos_neg_para = r.Intn(3) - 1

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

	return pos
}
