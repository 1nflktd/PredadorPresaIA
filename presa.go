package main

type Presa struct {
	AgenteImpl
}

func (p *Presa) getCAgente() CAgente {
	return C_Presa
}
