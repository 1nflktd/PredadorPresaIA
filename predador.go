package main

const VelocidadeMaximaPredador = 4

type Predador struct {
	AgenteImpl
	cacando bool
	iteracaoCacando int
}

func (p *Predador) getCAgente() CAgente {
	return C_Predador
}

func (p *Predador) mover(campoVisao CampoVisao) Posicao {
	// verifica se tem presa no campo de visao
	// se tem comeca a caca
	p.cacando = false
	direcao := Direcao(-1)
	for i, campo := range campoVisao.Posicoes {
		if campo.Agente == C_Presa {
			p.cacando = true

			if p.iteracaoCacando == 4 {
				p.iteracaoCacando = 0 // ultima iteracao velocidade maxima
			} else {
				p.iteracaoCacando++
			}

			direcao = Direcao(i)
			break
		}
	}

	if p.cacando {
		return p.cacar(direcao)
	} else {
		return p.viver()
	}
}

func (p *Predador) cacar(direcao Direcao) Posicao {
	velocidade := 1
	if (p.iteracaoCacando > 0) {
		velocidade = VelocidadeMaximaPredador
	}

	return p.moveAgente(direcao, velocidade)
}