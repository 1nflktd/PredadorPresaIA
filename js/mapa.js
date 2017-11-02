var CAgentes = {
	PREDADOR: 1,
	PRESA: 2,
	VAZIO: 3,
	MARCA1: 4,
	MARCA2: 5,
	MARCA3: 6,
	PRESA_FUGINDO: 7
};

function setLog(log) {
	var logPredadores = "<table class=\"table table-bordered table-responsive table-sm\">";
		logPredadores+= "<caption>Predadores</caption>";
		logPredadores+= "<thead>";
		logPredadores+= 	"<th>Posicao</th>";
		logPredadores+= 	"<th>Cacando</th>";
		logPredadores+= 	"<th>Iteracao cacando</th>";
		logPredadores+= 	"<th>Nro marcas</th>";
		logPredadores+= "</thead>";
		logPredadores+= "<tbody>";

	var logPresas = "<table class=\"table table-bordered table-responsive table-sm\">";
		logPresas+= "<caption>Presas</caption>";
		logPresas+= "<thead>";
		logPresas+= 	"<th>Posicao</th>";
		logPresas+= 	"<th>Fugindo</th>";
		logPresas+= 	"<th>Iteracao fugindo</th>";
		logPresas+= "</thead>";
		logPresas+= "<tbody>";

	var nAgentes = log.Agentes.length;
	for (i = 0; i < nAgentes; i++) {
		var agente = log.Agentes[i];

		var linha = "<tr>";
			linha+= "<td>X:" + agente.Posicao.X + " Y:" + agente.Posicao.Y + "</td>";
		if (agente.CAgente == CAgentes.PREDADOR) {
			logPredadores += linha;
			logPredadores += "<td>" + agente.Cacando + "</td>";
			logPredadores += "<td>" + agente.IteracaoCacando + "</td>";
			logPredadores += "<td>" + agente.NMarcas + "</td>";
			logPredadores += "</tr>";
		} else { // presa
			logPresas += linha;
			logPresas += "<td>" + agente.Fugindo + "</td>";
			logPresas += "<td>" + agente.IteracaoFugindo + "</td>";
			logPresas += "</tr>";
		}
	}

	logPredadores+= "</tbody>";
	logPredadores+= "</table>";

	logPresas+= "</tbody>";
	logPresas+= "</table>";

    document.getElementById("log-predadores").innerHTML = logPredadores;
    document.getElementById("log-presas").innerHTML = logPresas;

    if (log.PresasMortas != null) {
		var logPresasMortas = "<p>Presas mortas</p>";
			logPresasMortas+= "<table class=\"table table-bordered table-responsive table-sm\">";
			logPresasMortas+= "<thead>";
			logPresasMortas+= 	"<th>Posicao</th>";
			logPresasMortas+= 	"<th>Iteracao morreu</th>";
			logPresasMortas+= "</thead>";
			logPresasMortas+= "<tbody>";

		var nPresasMortas = log.PresasMortas.length;
		for (i = 0; i < nPresasMortas; i++) {
			var presaMorta = log.PresasMortas[i];

			logPresasMortas+= "<tr>";
			logPresasMortas+= "<td>X:" + presaMorta.Posicao.X + " Y:" + presaMorta.Posicao.Y + "</td>";
			logPresasMortas+= "<td>" + presaMorta.IteracaoMorreu + "</td>";
			logPresasMortas+= "</tr>";
		}

		logPresasMortas+= "</tbody>";
		logPresasMortas+= "</table>";

		document.getElementById("log-presas-mortas").innerHTML = logPresasMortas;
	}
}

function setAmbiente(ambiente) {
	var ambiente = JSON.parse(ambiente);

	console.log(ambiente);

	var mapa = ambiente.Mapa;
    var tamanhoMapa = ambiente.TamanhoMapa;

	var styleTd = "height: 25px; width: 25px; padding: 0px;";

	var table = "<table class=\"table table-bordered table-responsive\">";
		table+= "<caption>Iteração atual: " + ambiente.IteracaoAtual + "</caption>";
		table+= "<thead>";
		table+= "<tr>";
		table+= "<th style='" + styleTd + "' class='d-inline-block'></th>";
	for (i = 0; i < tamanhoMapa; i++) {
		table+= "<th style='" + styleTd + "' class='d-inline-block'>" + i + "</th>";
	}
		table+= "</tr>";
		table+= "</thead>";
		table+= "<tbody>";
	for (i = 0; i < tamanhoMapa; i++) {
		table += "<tr>";
		table += "<th scope='row' style='" + styleTd + "' class='d-inline-block'>" + i + "</th>";
		for (j = 0; j < tamanhoMapa; j++) {
			var imgName = "";
			var marca = "";
			switch(mapa[i][j]) {
				case CAgentes.PREDADOR:
					imgName = "predator.jpg";
					break;
				case CAgentes.PRESA:
					imgName = "homer.jpg";
					break;
				case CAgentes.PRESA_FUGINDO:
					imgName = "homer_alt.jpg";
					break;
				case CAgentes.MARCA1:
					marca = "#ffff00";
					break;
				case CAgentes.MARCA2:
					marca = "#ff9933";
					break;
				case CAgentes.MARCA3:
					marca = "#ff471a";
					break;
				case CAgentes.VAZIO:
					marca = "white";
					break;
			}

			var style = styleTd;
			var img = imgName != "" ? "<img src=\"../images/" + imgName + "\">" : "";
			style += marca != "" ? "background-color:" + marca : "";
			table += "<td style='" + style + "' class='d-inline-block'>" + img + "</td>";
		}
		table += "</tr>";
	}
	table += "</tbody>";

	document.getElementById("mapa").innerHTML = table;

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
