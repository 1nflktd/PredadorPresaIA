package main

type Log struct {
	Agentes []LogAgente
	PresasMortas []LogAgente
}

type LogAgente struct {
	CAgente CAgente
	Posicao Posicao
	 // predador
	Cacando bool
	IteracaoCacando int
	NMarcas int
	 // presa
	Fugindo bool
	IteracaoFugindo int
	IteracaoMorreu int
}

func (l *Log) excluirAgentes() {
	l.Agentes = nil
}

func (l *Log) adicionarAgente(agente Agente) {
	logAgente := LogAgente{}
	logAgente.CAgente = agente.getCAgente()
	logAgente.Posicao = agente.getPosicao()

	if predador, ok := agente.(*Predador); ok {
		logAgente.Cacando = predador.getCacando()
		logAgente.IteracaoCacando = predador.getIteracaoCacando()
		logAgente.NMarcas = len(predador.getMarcas())
	} else if presa, ok := agente.(*Presa); ok {
		logAgente.Fugindo = presa.getFugindo()
		logAgente.IteracaoFugindo = presa.getIteracaoFugindo()
	}

	l.Agentes = append(l.Agentes, logAgente)
}

func (l *Log) adicionarPresaMorta(agente Agente, iteracaoMorreu int) {
	if presaMorta, ok := agente.(*Presa); ok {
		logAgente := LogAgente{}
		logAgente.CAgente = C_Presa
		logAgente.Posicao = presaMorta.getPosicao()
		logAgente.IteracaoMorreu = iteracaoMorreu
		l.PresasMortas = append(l.PresasMortas, logAgente)
	}
}
