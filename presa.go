package main

const VelocidadeMaximaPresa = 2

type Presa struct {
	AgenteImpl
	morreu bool
}

func (p *Presa) getCAgente() CAgente {
	return C_Presa
}

func (p *Presa) mover(campoVisao CampoVisao) (Posicao, PosMovimento) {
	// verifica se tem predador
	// verifica se tem presa que mudou de cor (???)

	direcao := Direcao(-1)
	qtdePredadores := 0
	for i, campo := range campoVisao.Posicoes {
		if campo.Agente == C_Predador {
			direcao = Direcao(i)
			qtdePredadores++
		}
	}

	if qtdePredadores >= 3 {
		return p.morrer()
	} else if qtdePredadores > 0 {
		return p.fugir(direcao)
	} else {
		return p.viver()
	}
}

func (p *Presa) fugir(direcao Direcao) (Posicao, PosMovimento) {
	// vai na direcao oposta
	return p.moveAgente(ObterDirecaoOposta(direcao), 1)
}

func (p *Presa) morrer() (Posicao, PosMovimento) {
	p.morreu = true
	return Posicao{}, PosMovimento(-1)
}

func (p *Presa) getMorreu() bool {
	return p.morreu
}
