package main

import (
	"time"
	"math/rand"
	"sync"
)

const TamanhoMapa = 30

type CAgente int
const (
	C_Predador CAgente = iota + 1
	C_Presa
	C_Vazio
	C_Marca1
	C_Marca2
	C_Marca3
)

type Mapa [TamanhoMapa][TamanhoMapa]CAgente
type AmbienteTela struct {
	Mapa Mapa
	LimiteIteracoes bool
	TamanhoMapa int
	PresasCapturadas bool
}

type Ambiente struct {
	mapa Mapa
	agentes []Agente
	limiteIteracoes bool
	presasCapturadas bool
}

func (a *Ambiente) Init(nPresas, nPredadores int) {
	// inicia todos em branco
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			a.mapa[i][w] = C_Vazio
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// coloca presas (aleatorio)
	for i := 0; i < nPresas; {
		p1, p2 := r.Intn(TamanhoMapa), r.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.mapa[p1][p2] = C_Presa
			presa := &Presa{}
			presa.setPosicaoXY(p1, p2)
			a.agentes = append(a.agentes, presa)
			i++
		}
	}

	// coloca predadores (aleatorio)
	for i := 0; i < nPredadores; {
		p1, p2 := r.Intn(TamanhoMapa), r.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.mapa[p1][p2] = C_Predador
			predador := &Predador{}
			predador.setPosicaoXY(p1, p2)
			a.agentes = append(a.agentes, predador)
			i++
		}
	}
}

func (a *Ambiente) GetAmbienteTela() AmbienteTela {
	return AmbienteTela{
		Mapa: a.mapa,
		LimiteIteracoes: a.limiteIteracoes,
		TamanhoMapa: TamanhoMapa,
		PresasCapturadas: a.presasCapturadas,
	}
}

func (a *Ambiente) Run() {
	for i := 0; i < 5000; i++ {
		a.moveAgentes()
		time.Sleep(500 * time.Millisecond)
	}
	a.limiteIteracoes = true
}

func (a *Ambiente) moveAgentes() {
	qtdeAgentes := len(a.agentes)
	agentes := make(chan bool, qtdeAgentes)
	mutexMapa := &sync.Mutex{}
	for _, ag := range a.agentes {
		go func(agente Agente) {
			posAtual := agente.getPosicao()
			campoVisao := ObterCampoVisao(a.mapa, posAtual)
			posNova := agente.mover(campoVisao)

			morreu := false
			if presa, ehPresa := agente.(*Presa); ehPresa {
				morreu = presa.getMorreu()
			}

			if morreu {
				mutexMapa.Lock()
				a.mapa[posAtual.X][posAtual.Y] = C_Vazio
				mutexMapa.Unlock()
			} else {
				mutexMapa.Lock()
				ok, _ := a.verificaColisao(posNova)
				mutexMapa.Unlock()

				if ok {
					if predador, ehPredador := agente.(*Predador); ehPredador {
						predador.adicionarMarcas(posAtual, posNova)
					}

					mutexMapa.Lock()
					a.mapa[posAtual.X][posAtual.Y] = C_Vazio
					agente.setPosicao(posNova) // move o elemento
					a.mapa[posNova.X][posNova.Y] = agente.getCAgente()
					mutexMapa.Unlock()
				}
			}

			agentes <- true
		}(ag)
	}

	for i := 0; i < qtdeAgentes; i++ {
		<-agentes
	}
}

func (a *Ambiente) verificaColisao(posAgente Posicao) (bool, CAgente) {
	c := a.mapa[posAgente.X][posAgente.Y]
	if c == C_Vazio {
		return true, c
	} else {
		return false, c
	}
}

func VerificaLimites(coordenada int) int {
	if coordenada >= TamanhoMapa {
		coordenada = 0
	} else if coordenada < 0 {
		coordenada = TamanhoMapa - 1
	}
	return coordenada
}
