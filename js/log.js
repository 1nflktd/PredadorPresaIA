function getTableLogPredadores(logPredadores) {
	if (logPredadores == null) {
		return "";
	}

	var table = "<table class=\"table table-bordered table-responsive table-sm\">" +
				"<caption>Predadores</caption>" +
				"<thead>" +
					"<th>Id</th>" +
					"<th>Posição</th>" +
					"<th>Caçando</th>" +
					"<th>Iter. caçando</th>" +
					"<th>Marcas</th>" +
				"</thead>" +
				"<tbody>";

	var nPredadores = logPredadores.length;
	for (i = 0; i < nPredadores; i++) {
		for (j = 0; j < nPredadores; j++) {
			var predador = logPredadores[j];

			if (predador.Id != i) {
				continue;
			}

			table+= "<tr>" +
						"<td>" + predador.Id + "</td>" +
						"<td>X:" + predador.Posicao.X + " Y:" + predador.Posicao.Y + "</td>" +
						"<td>" + predador.Cacando + "</td>" +
						"<td>" + predador.IteracaoCacando + "</td>" +
						"<td>" + predador.NMarcas + "</td>" +
					"</tr>";

			break;
		}
	}

	table+= "</tbody>" +
			"</table>";

	return table;
}

function getTableLogPresas(logPresas, qtdePresas) {
	if (logPresas == null) {
		return "";
	}

	var table = "<table class=\"table table-bordered table-responsive table-sm\">" +
				 "<caption>Presas</caption>" +
				 "<thead>" +
				 	"<th>Id</th>" +
				 	"<th>Posição</th>" +
				 	"<th>Fugindo</th>" +
				 	"<th>Iter. fugindo</th>" +
				 	"<th>QE</th>" +
				 	"<th>IE</th>" +
				 "</thead>" +
				 "<tbody>";

	var nPresas = qtdePresas;
	for (i = 0; i < nPresas; i++) {
		for (j = 0; j < nPresas; j++) {
			if (logPresas[j] == null) {
				continue;
			}

			var presa = logPresas[j];

			if (presa.Id != i) {
				continue;
			}

			table+= "<tr>" +
						"<td>" + presa.Id + "</td>" +
						"<td>X:" + presa.Posicao.X + " Y:" + presa.Posicao.Y + "</td>" +
						"<td>" + presa.Fugindo + "</td>" +
						"<td>" + presa.IteracaoFugindo + "</td>" +
						"<td>" + presa.QualidadeEmocao + "</td>" +
						"<td>" + presa.IntensidadeEmocao + "</td>" +
					"</tr>";

			break;
		}
	}

	table+= "</tbody>" +
			"</table>";

	return table;
}

function getTableLogPresasMortas(logPresasMortas) {
    if (logPresasMortas == null) {
    	return "";
    }

	var table = "<table class=\"table table-bordered table-responsive table-sm\">" +
				"<caption>Presas mortas</caption>" +
				"<thead>" +
					"<th>Id</th>" +
					"<th>Posição</th>" +
					"<th>Iter. morreu</th>" +
				"</thead>" +
				"<tbody>";

	var nPresasMortas = logPresasMortas.length;
	for (i = 0; i < nPresasMortas; i++) {
		var presaMorta = logPresasMortas[i];

		table+= "<tr>" +
					"<td>" + presaMorta.Id + "</td>" +
					"<td>X:" + presaMorta.Posicao.X + " Y:" + presaMorta.Posicao.Y + "</td>" +
					"<td>" + presaMorta.IteracaoMorreu + "</td>" +
				"</tr>";
	}

	table+= "</tbody>" +
			"</table>";

	return table;
}

function setLog(log) {
	var qtdePresas = 0;
	qtdePresas += log.Presas != null ? log.Presas.length : 0;
	qtdePresas += log.PresasMortas != null ? log.PresasMortas.length : 0;

	console.log("qtdePresas " + qtdePresas);

    document.getElementById("log-predadores").innerHTML = getTableLogPredadores(log.Predadores);
    document.getElementById("log-presas").innerHTML = getTableLogPresas(log.Presas, qtdePresas);
	document.getElementById("log-presas-mortas").innerHTML = getTableLogPresasMortas(log.PresasMortas);
}
