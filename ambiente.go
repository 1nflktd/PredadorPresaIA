package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"net/http"
)

const TamanhoMapa = 15

type Caracter string
var C_Diamante Caracter = "*"
var C_Pedra Caracter = "#"
var C_Base Caracter = "B"
var C_Vazio Caracter = " "

type AmbienteTela struct {
	Mapa [TamanhoMapa][TamanhoMapa]Caracter
}

type Ambiente struct {
	mapa [TamanhoMapa][TamanhoMapa]Caracter
	diamantes int
	agentes []*Agente
	base Posicao
	w http.ResponseWriter
}

func (a *Ambiente) Init(w http.ResponseWriter, nDiamantes, nPedras, nAgentes int) {
	a.w = w
	// inicia todos em branco
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			a.mapa[i][w] = C_Vazio
		}
	}

	// coloca a base na posicao certa
	a.base = Posicao{0, 0}
	a.mapa[0][0] = C_Base

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// coloca diamantes (aleatorio)
	for i := 0; i < nDiamantes; {
		p1, p2 := r.Intn(TamanhoMapa), r.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.diamantes++
			a.mapa[p1][p2] = C_Diamante
			i++
		}
	}

	// coloca pedras (aleatorio)
	for i := 0; i < nPedras; {
		p1, p2 := r.Intn(TamanhoMapa), r.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.mapa[p1][p2] = C_Pedra
			i++
		}
	}

	// coloca agente (aleatorio)
	for i := 0; i < nAgentes; {
		p1, p2 := r.Intn(TamanhoMapa), r.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			agente1 := &Agente{}
			agente1.Init(a.base, Caracter(rune(i) + 'a'))
			a.mapa[p1][p2] = agente1.getCaracter()
			agente1.setPosicaoXY(p1, p2)
			a.agentes = append(a.agentes, agente1)
			i++
		}
	}

}

func (a *Ambiente) PrintMapa() {
	executeTemplate(a.w, AmbienteTela{Mapa: a.mapa})
}

func (a *Ambiente) PrintInfo() {
	for _, ag := range a.agentes {
		fmt.Printf("Agente %s:...", ag.getCaracter())
		if ag.getTemDiamante() {
			fmt.Printf("Tem diamante. Voltando para base\n")
		} else {
			fmt.Printf("Procurando diamantes\n")
		}
	}
}

func (a *Ambiente) Run() {
	// laco ate encontrar todos os diamantes
	for {
		//limpaTela()
		a.PrintMapa()
		//a.PrintInfo()
		a.moveAgentes()
		if a.diamantes == 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Todos diamantes foram encontrados e levados para a base!\n")
}

func (a *Ambiente) moveAgentes() {
	qtdeAgentes := len(a.agentes)
	agentes := make(chan bool, qtdeAgentes)
	mutexMapa := &sync.Mutex{}
	for _, ag := range a.agentes {
		go func(agente *Agente) {
			posAtual := agente.getPosicao()
			var p_ag Posicao
			if agente.getTemDiamante() {
				p_ag = agente.voltaBase()
			} else {
				p_ag = agente.movePosAleatorio()
			}

			mutexMapa.Lock()
			ok, caracter := a.verificaColisao(p_ag)
			mutexMapa.Unlock()

			if ok {
				atualizarPos := true
				if caracter == C_Diamante {
					if agente.getTemDiamante() { // ja tem diamante, pula
						atualizarPos = false
					} else {
						agente.setTemDiamante(true)
					}
				}

				if atualizarPos {
					mutexMapa.Lock()
					a.mapa[posAtual.X][posAtual.Y] = C_Vazio
					agente.setPosicao(p_ag) // move o elemento
					a.mapa[p_ag.X][p_ag.Y] = agente.getCaracter()
					mutexMapa.Unlock()
				}
			} else if caracter == C_Base {
				if agente.getTemDiamante() {
					agente.setTemDiamante(false)
					a.diamantes--
				}
			}
			agentes <- true
		}(ag)
	}

	for i := 0; i < qtdeAgentes; i++ {
		<-agentes
	}
}

func (a *Ambiente) verificaColisao(posAgente Posicao) (bool, Caracter) {
	c := a.mapa[posAgente.X][posAgente.Y]
	if c == C_Diamante || c == C_Vazio {
		return true, c
	} else {
		return false, c
	}
}
