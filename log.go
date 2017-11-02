package main

type Log struct {
	Presas []LogPresa
	Predadores []LogPredador
	PresasMortas []LogPresa
}

type LogAgente struct {
	Id int
	CAgente CAgente
	Posicao Posicao
}

type LogPresa struct {
	LogAgente
	Fugindo bool
	IteracaoFugindo int
	IteracaoMorreu int
	QualidadeEmocao int
	IntensidadeEmocao int
}

type LogPredador struct {
	LogAgente
	Cacando bool
	IteracaoCacando int
	NMarcas int
}

func (l *Log) excluirAgentes() {
	l.Presas = nil
	l.Predadores = nil
}

func (l *Log) adicionarAgente(agente Agente) {
	if predador, ok := agente.(*Predador); ok {
		l.adicionarPredador(predador)
	} else if presa, ok := agente.(*Presa); ok {
		l.adicionarPresa(presa)
	}
}

func (l *Log) adicionarPresa(presa *Presa) {
	logPresa := LogPresa{}
	logPresa.Id = presa.getId()
	logPresa.CAgente = presa.getCAgente()
	logPresa.Posicao = presa.getPosicao()
	logPresa.Fugindo = presa.getFugindo()
	logPresa.IteracaoFugindo = presa.getIteracaoFugindo()
	logPresa.QualidadeEmocao = presa.getQualidadeEmocao()
	logPresa.IntensidadeEmocao = presa.getIntensidadeEmocao()

	l.Presas = append(l.Presas, logPresa)
}

func (l *Log) adicionarPredador(predador *Predador) {
	logPredador := LogPredador{}
	logPredador.Id = predador.getId()
	logPredador.CAgente = predador.getCAgente()
	logPredador.Posicao = predador.getPosicao()
	logPredador.Cacando = predador.getCacando()
	logPredador.IteracaoCacando = predador.getIteracaoCacando()
	logPredador.NMarcas = len(predador.getMarcas())

	l.Predadores = append(l.Predadores, logPredador)
}

func (l *Log) adicionarPresaMorta(agente Agente, iteracaoMorreu int) {
	if presaMorta, ok := agente.(*Presa); ok {
		logPresa := LogPresa{}
		logPresa.Id = presaMorta.getId()
		logPresa.CAgente = C_Presa
		logPresa.Posicao = presaMorta.getPosicao()
		logPresa.IteracaoMorreu = iteracaoMorreu
		l.PresasMortas = append(l.PresasMortas, logPresa)
	}
}
