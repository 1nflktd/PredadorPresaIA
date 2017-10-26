package main

type Direcao int
const (
	P_Norte Direcao = iota
	P_Noroeste
	P_Nordeste
	P_Leste
	P_Oeste
	P_Sudoeste
	P_Sudeste
	P_Sul
	P_Aleatoria Direcao = 99
)

type CampoVisao struct {
	// 0 - y_pls_1 			// norte
	// 1 - y_pls_1_x_min_1  // noroeste
	// 2 - x_pls_1_y_pls_1  // nordeste
	// 3 - x_pls_1 			// leste
	// 4 - x_min_1 			// oeste
	// 5 - y_min_1_x_min_1	// sudoeste
	// 6 - y_min_1_x_pls_1	// sudeste
	// 7 - y_min_1 			// sul

	Posicoes [16]struct {
		Pos Posicao
		Agente CAgente
	}
}

func ObterCampoVisao(mapa Mapa, posAgente Posicao) (CampoVisao) {
	campoVisao := CampoVisao{}

	for i := 1; i < 3; i++ {
		yNorte := posAgente.Y + i
		if yNorte >= TamanhoMapa {
			yNorte = 0
		}

		ySul := posAgente.Y - i
		if ySul < 0 {
			ySul = TamanhoMapa - 1
		}

		xLeste := posAgente.X + i
		if xLeste >= TamanhoMapa {
			xLeste = 0
		}

		xOeste := posAgente.X - i
		if xOeste < 0 {
			xOeste = TamanhoMapa - 1
		}

		posVetor := Direcao((i - 1) * 8)

		// norte
		campoVisao.Posicoes[P_Norte + posVetor].Pos = Posicao{posAgente.X, yNorte}
		campoVisao.Posicoes[P_Norte + posVetor].Agente = mapa[posAgente.X][yNorte]

		// noroeste
		campoVisao.Posicoes[P_Noroeste + posVetor].Pos = Posicao{xOeste, yNorte}
		campoVisao.Posicoes[P_Noroeste + posVetor].Agente = mapa[xOeste][yNorte]

		// nordeste
		campoVisao.Posicoes[P_Nordeste + posVetor].Pos = Posicao{xLeste, yNorte}
		campoVisao.Posicoes[P_Nordeste + posVetor].Agente = mapa[xLeste][yNorte]

		// leste
		campoVisao.Posicoes[P_Leste + posVetor].Pos = Posicao{xLeste, posAgente.Y}
		campoVisao.Posicoes[P_Leste + posVetor].Agente = mapa[xLeste][posAgente.Y]

		// oeste
		campoVisao.Posicoes[P_Oeste + posVetor].Pos = Posicao{xOeste, posAgente.Y}
		campoVisao.Posicoes[P_Oeste + posVetor].Agente = mapa[xOeste][posAgente.Y]

		// sudoeste
		campoVisao.Posicoes[P_Sudoeste + posVetor].Pos = Posicao{xOeste, ySul}
		campoVisao.Posicoes[P_Sudoeste + posVetor].Agente = mapa[xOeste][ySul]

		// sudeste
		campoVisao.Posicoes[P_Sudeste + posVetor].Pos = Posicao{xLeste, ySul}
		campoVisao.Posicoes[P_Sudeste + posVetor].Agente = mapa[xLeste][ySul]

		// sul
		campoVisao.Posicoes[P_Sul + posVetor].Pos = Posicao{posAgente.X, ySul}
		campoVisao.Posicoes[P_Sul + posVetor].Agente = mapa[posAgente.X][ySul]
	}

	return campoVisao
}

func ObterDirecaoOposta(direcao Direcao) Direcao {
	switch (direcao) {
		case P_Norte:
			return P_Sul
		case P_Noroeste:
			return P_Sudeste
		case P_Nordeste:
			return P_Sudoeste
		case P_Leste:
			return P_Oeste
		case P_Oeste:
			return P_Leste
		case P_Sudoeste:
			return P_Nordeste
		case P_Sudeste:
			return P_Noroeste
		case P_Sul:
			return P_Norte
	}

	return direcao
}
