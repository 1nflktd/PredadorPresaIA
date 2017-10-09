package main

type Predador struct {
	AgenteImpl
}

func (p *Predador) getCAgente() CAgente {
	return C_Predador
}
