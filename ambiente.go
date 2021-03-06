package main

import (
	"time"
	"math/rand"
	"sync"

//	"log"
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
	C_Presa_Fugindo
)

type Mapa [TamanhoMapa][TamanhoMapa]CAgente

type AmbienteTela struct {
	Mapa Mapa
	Log Log
	LimiteIteracoes bool
	TamanhoMapa int
	PresasTotais int
	IteracaoAtual int
}

type Ambiente struct {
	mapa Mapa
	log Log
	agentes []Agente
	limiteIteracoes bool
	presasTotais int
	mutexMapa *sync.Mutex
	mutexLog *sync.Mutex
	mutexIteracaoAtual *sync.Mutex
	iteracaoAtual int
}

func (a *Ambiente) Init(nPresas, nPredadores int) {
	a.mutexMapa = &sync.Mutex{}
	a.mutexLog = &sync.Mutex{}
	a.mutexIteracaoAtual = &sync.Mutex{}

	a.log = Log{}

	// inicia todos em branco
	for i := 0; i < TamanhoMapa; i++ {
		for w := 0; w < TamanhoMapa; w++ {
			a.mapa[i][w] = C_Vazio
		}
	}

	randPos := rand.New(rand.NewSource(time.Now().UnixNano()))
	mutexRandPos := &sync.Mutex{}

	// coloca presas (aleatorio)
	a.presasTotais = nPresas
	for i := 0; i < nPresas; {
		p1, p2 := randPos.Intn(TamanhoMapa), randPos.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.mapa[p1][p2] = C_Presa
			presa := &Presa{}
			presa.Init(i)
			presa.setRandPos(randPos, mutexRandPos)
			presa.setPosicaoXY(p1, p2)
			a.agentes = append(a.agentes, presa)
			i++
		}
	}

	// coloca predadores (aleatorio)
	for i := 0; i < nPredadores; {
		p1, p2 := randPos.Intn(TamanhoMapa), randPos.Intn(TamanhoMapa)
		if a.mapa[p1][p2] == C_Vazio {
			a.mapa[p1][p2] = C_Predador
			predador := &Predador{}
			predador.Init(i)
			predador.setRandPos(randPos, mutexRandPos)
			predador.setPosicaoXY(p1, p2)
			a.agentes = append(a.agentes, predador)
			i++
		}
	}
}

func (a *Ambiente) GetAmbienteTela() AmbienteTela {
	a.mutexMapa.Lock()
	mapa := a.mapa
	limiteIteracoes := a.limiteIteracoes
	presasTotais := a.presasTotais
	a.mutexMapa.Unlock()

	a.mutexIteracaoAtual.Lock()
	iteracaoAtual := a.iteracaoAtual
	a.mutexIteracaoAtual.Unlock()

	a.mutexLog.Lock()
	log := a.log
	a.mutexLog.Unlock()

	ambienteTela := AmbienteTela{
		Mapa: mapa,
		LimiteIteracoes: limiteIteracoes,
		TamanhoMapa: TamanhoMapa,
		PresasTotais: presasTotais,
		IteracaoAtual: iteracaoAtual,
		Log: log,
	}

	return ambienteTela
}

func (a *Ambiente) Run() {
	for i := 0; i < 5000; i++ {
		a.mutexIteracaoAtual.Lock()
		a.iteracaoAtual = i
		a.mutexIteracaoAtual.Unlock()
		a.moveAgentes()
		time.Sleep(500 * time.Millisecond)
	}
	a.limiteIteracoes = true
}

func (a *Ambiente) moveAgentes() {
	a.mutexLog.Lock()
	a.log.excluirAgentes()
	a.mutexLog.Unlock()

	qtdeAgentes := len(a.agentes)
	agentes := make(chan bool, qtdeAgentes)

	presasFaltantes := make(chan int, a.presasTotais)
	for iAg, ag := range a.agentes {
		go func(agente Agente, iAg int) {
			posAtual := agente.getPosicao()
			a.mutexMapa.Lock()
			var campoVisao CampoVisao
			if _, ehPresa := agente.(*Presa); ehPresa {
				campoVisao = ObterCampoVisaoPresa(a.mapa, posAtual)
			} else {
				campoVisao = ObterCampoVisaoPredador(a.mapa, posAtual)
			}
			a.mutexMapa.Unlock()
			posNova, posMovimento := agente.mover(campoVisao)

			morreu := false
			if presa, ehPresa := agente.(*Presa); ehPresa {
				morreu = presa.getMorreu()
			}

			if morreu {
				a.mutexMapa.Lock()
				a.mapa[posAtual.X][posAtual.Y] = C_Vazio
				a.presasTotais--
				a.mutexMapa.Unlock()
				presasFaltantes <- iAg
			} else {
				a.mutexMapa.Lock()
				ok, _ := a.verificaColisao(posNova)

				if ok {
					if predador, ehPredador := agente.(*Predador); ehPredador {
						a.eliminarMarcasMapa(predador.getMarcas())
						predador.adicionarMarcas(posAtual, posNova, posMovimento)
						a.adicionarMarcasMapa(predador.getMarcas())
					}
					a.mapa[posAtual.X][posAtual.Y] = C_Vazio
					agente.setPosicao(posNova) // move o elemento
					a.mapa[posNova.X][posNova.Y] = agente.getCAgente()
				}
				a.mutexMapa.Unlock()

				a.mutexLog.Lock()
				a.log.adicionarAgente(agente)
				a.mutexLog.Unlock()
			}

			agentes <- true
		}(ag, iAg)
	}

	for i := 0; i < qtdeAgentes; i++ {
		<-agentes
	}

	for i := 0; i < len(presasFaltantes); i++ {
		select {
			case idAg := <-presasFaltantes:
				a.mutexIteracaoAtual.Lock()
				iteracaoAtual := a.iteracaoAtual
				a.mutexIteracaoAtual.Unlock()

				a.mutexLog.Lock()
				a.log.adicionarPresaMorta(a.agentes[idAg], iteracaoAtual)
				a.mutexLog.Unlock()

				// remove
				a.agentes = append(a.agentes[:idAg], a.agentes[idAg+1:]...)
		}
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
		if intensidade > (2 * IntensidadeMarcaMul) {
			return C_Marca3
		} else if intensidade > IntensidadeMarcaMul {
			return C_Marca2
		} else if intensidade > 0 {
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
