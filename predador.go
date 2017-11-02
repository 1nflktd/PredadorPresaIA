package main

import (
//	"log"
)

const VelocidadeMaximaPredador = 4

type Predador struct {
	AgenteImpl
	cacando bool
	iteracaoCacando int
	marcas []Marca
}

const IntensidadeMarcaMul = 5
type Marca struct {
	Pos Posicao
	Intensidade int
}

func (p *Predador) Init() {
	p.cacando = false
	p.iteracaoCacando = 0
}

func (p *Predador) getCAgente() CAgente {
	return C_Predador
}

func (p *Predador) mover(campoVisao CampoVisao) (Posicao, PosMovimento) {
	// verifica se tem presa no campo de visao
	// se tem comeca a caca
	p.cacando = false
	direcao := Direcao(-1)
	for i, campo := range campoVisao.Posicoes {
		if campo.Agente == C_Presa || campo.Agente == C_Presa_Fugindo {
			p.cacando = true

			if p.iteracaoCacando == 4 {
				p.iteracaoCacando = 0 // ultima iteracao velocidade maxima
			} else {
				p.iteracaoCacando++
			}

			direcao = Direcao(i % 8)
			break
		} else if VerificaSeEhMarca(campo.Agente) {
			direcao = Direcao(i % 8)
		}
	}

	// diminuir intensidade das marcas atuais
	marcas := []Marca{}
	for _, marca := range p.marcas {
		marca.Intensidade--

		if marca.Intensidade > -5 {
			marcas = append(marcas, marca)
		}
	}
	p.marcas = marcas

	if p.cacando {
		return p.cacar(direcao)
	} else if !p.cacando && direcao != Direcao(-1) {
		return p.seguirMarca(direcao)
	} else {
		return p.viver()
	}
}

func (p *Predador) seguirMarca(direcao Direcao) (Posicao, PosMovimento) {
	return p.moveAgente(direcao, 1)
}

func (p *Predador) cacar(direcao Direcao) (Posicao, PosMovimento) {
	velocidade := 1
	if (p.iteracaoCacando > 0) {
		velocidade = VelocidadeMaximaPredador
	}

	return p.moveAgente(direcao, velocidade)
}

func (p *Predador) adicionarMarcas(posAtual, posNova Posicao, posMovimento PosMovimento) {
	if p.cacando {
		fValAltera := func(posM PosMovimento) int {
			if posM == AUM_X || posM == AUM_Y {
				return 1
			}
			return -1
		}

		fMaiorMenor := func(pAtual, pNovo, valAltera int) bool {
			if valAltera < 0 {
				return pAtual > pNovo
			}
			return pNovo < pAtual
		}

		if posAtual.X != posNova.X {
			// mudou x
			valAltera := fValAltera(posMovimento)
			for x := posAtual.X; fMaiorMenor(VerificaLimites(x), posNova.X, valAltera); x += valAltera {
				x = VerificaLimites(x)
				p.marcas = append(p.marcas, Marca{Pos: Posicao{X: x, Y: posNova.Y}, Intensidade: 3 * IntensidadeMarcaMul})
			}
		} else {
			// mudou y
			valAltera := fValAltera(posMovimento)
			for y := posAtual.Y; fMaiorMenor(VerificaLimites(y), posNova.Y, valAltera); y += valAltera {
				y = VerificaLimites(y)
				p.marcas = append(p.marcas, Marca{Pos: Posicao{X: posNova.X, Y: y}, Intensidade: 3 * IntensidadeMarcaMul})
			}
		}
	}
}

func (p *Predador) getMarcas() []Marca {
	return p.marcas
}

func (p *Predador) getCacando() bool {
	return p.cacando
}

func (p *Predador) getIteracaoCacando() int {
	return p.iteracaoCacando
}
