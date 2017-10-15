package main

const VelocidadeMaximaPresa = 2

type Presa struct {
	AgenteImpl
}

func (p *Presa) getCAgente() CAgente {
	return C_Presa
}

func (p *Presa) move(campoVisao CampoVisao) Posicao {
	return p.viver()
}
