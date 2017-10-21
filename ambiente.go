package main

import (
	"time"
	"math/rand"
	"sync"

//	"log"
)

const TamanhoMapa = 5

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
	PresasTotais int
}

type Ambiente struct {
	mapa Mapa
	agentes []Agente
	limiteIteracoes bool
	presasTotais int
	mutexMapa *sync.Mutex
}

func (a *Ambiente) Init(nPresas, nPredadores int) {
	a.mutexMapa = &sync.Mutex{}
	// inicia todos em branco
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			a.mapa[i][w] = C_Vazio
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// coloca presas (aleatorio)
	a.presasTotais = nPresas
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
	a.mutexMapa.Lock()
	ambienteTela := AmbienteTela{
		Mapa: a.mapa,
		LimiteIteracoes: a.limiteIteracoes,
		TamanhoMapa: TamanhoMapa,
		PresasTotais: a.presasTotais,
	}
	a.mutexMapa.Unlock()

	return ambienteTela
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
	for _, ag := range a.agentes {
		go func(agente Agente) {
			a.mutexMapa.Lock()
			posAtual := agente.getPosicao()
			campoVisao := ObterCampoVisao(a.mapa, posAtual)
			posNova := agente.mover(campoVisao)

			morreu := false
			if presa, ehPresa := agente.(*Presa); ehPresa {
				morreu = presa.getMorreu()
			}

			if morreu {
				a.mapa[posAtual.X][posAtual.Y] = C_Vazio
				a.presasTotais--
			} else {
				ok, _ := a.verificaColisao(posNova)

				if ok {
					if predador, ehPredador := agente.(*Predador); ehPredador {
						a.eliminarMarcasMapa(predador.getMarcas())
						predador.adicionarMarcas(posAtual, posNova)
						a.adicionarMarcasMapa(predador.getMarcas())
					}
					a.mapa[posAtual.X][posAtual.Y] = C_Vazio
					agente.setPosicao(posNova) // move o elemento
					a.mapa[posNova.X][posNova.Y] = agente.getCAgente()
				}
			}
			a.mutexMapa.Unlock()

			agentes <- true
		}(ag)
	}

	for i := 0; i < qtdeAgentes; i++ {
		<-agentes
	}
}

func (a *Ambiente) verificaColisao(posAgente Posicao) (bool, CAgente) {
	c := a.mapa[posAgente.X][posAgente.Y]
	if c == C_Vazio || c == C_Marca3 || c == C_Marca2 || c == C_Marca1 {
		return true, c
	} else {
		return false, c
	}
}

func (a *Ambiente) eliminarMarcasMapa(marcas []Marca) {
	for _, marca := range marcas {
		if ok, _ := a.verificaColisao(marca.Pos); ok {
			a.mapa[marca.Pos.X][marca.Pos.Y] = C_Vazio
		}
	}
}

func (a *Ambiente) adicionarMarcasMapa(marcas []Marca) {
	fnCAgenteMarca := func(intensidade int) CAgente {
		switch(intensidade) {
			case 3:
				return C_Marca3
			case 2:
				return C_Marca2
			case 1:
				return C_Marca1
		}
		return C_Vazio
	}

	for _, marca := range marcas {
		if ok, _ := a.verificaColisao(marca.Pos); ok {
			a.mapa[marca.Pos.X][marca.Pos.Y] = fnCAgenteMarca(marca.Intensidade)
		}
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

func VerificaSeEhMarca(cAgente CAgente) bool {
	return cAgente == C_Marca1 || cAgente == C_Marca2 || cAgente == C_Marca3
}
