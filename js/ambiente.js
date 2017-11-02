var CAgentes = {
	PREDADOR: 1,
	PRESA: 2,
	VAZIO: 3,
	MARCA1: 4,
	MARCA2: 5,
	MARCA3: 6,
	PRESA_FUGINDO: 7
};

function setAmbiente(ambiente) {
	var ambiente = JSON.parse(ambiente);

	setMapa(ambiente.Mapa, ambiente.TamanhoMapa, ambiente.IteracaoAtual);
    setLog(ambiente.Log);

	if (ambiente.LimiteIteracoes) {
		$(".msgs").hide();
		$("#limite-iteracoes").show();
	} else if (ambiente.PresasTotais <= 0) {
		$(".msgs").hide();
		$("#presas-capturadas").show();
	}
}

function iniciarEventSource() {
	var source = new EventSource('/events/');

	source.onmessage = function(e) {
		setAmbiente(e.data);
	};

	source.onerror = function(e){
		$(".msgs").hide();
		$("#servidor-morto").show();
	};
}
