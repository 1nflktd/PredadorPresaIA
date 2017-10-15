package main

const VelocidadeMaximaPredador = 4

type Predador struct {
	AgenteImpl
	cacando bool
	iteracaoCacando int
	marcas []Marca
}

type Marca struct {
	Pos Posicao
	Intensidade int
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

	// diminuir intensidade das marcas atuais
	/*
	marcas := []Marca{}
	for _, marca := range p.marcas {
		marca.Intensidade--
		if marca.Intensidade > 0 {
			// manter
			marcas = append(marcas, marca)
		}
	}

	p.marcas = marcas
	*/

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

func (p *Predador) adicionarMarcas(posAtual, posNova Posicao) {
	if p.cacando {
		/*
		fValAltera := func(pAtual, pNovo int) int {
			var valAltera int
			if pAtual < pNovo {
				valAltera = -1
			} else {
				valAltera = 1
			}
			return valAltera
		}

		if posAtual.X != posNova.X {
			// mudou x
			valAltera := fValAltera(posAtual.X, posNova.X)
			for x := posAtual.X; x <= posNova.X; x += valAltera {
				x = VerificaLimites(x)
				p.marcas = append(p.marcas, Marca{Pos: Posicao{X: x, Y: posNova.Y}, Intensidade: 3})
			}
		} else {
			// mudou y
			valAltera := fValAltera(posAtual.Y, posNova.Y)
			for y := posAtual.Y; y <= posNova.Y; y += valAltera {
				y = VerificaLimites(y)
				p.marcas = append(p.marcas, Marca{Pos: Posicao{X: posNova.X, Y: y}, Intensidade: 3})
			}
		}
		*/
	}
}
