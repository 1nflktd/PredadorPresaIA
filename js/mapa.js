function setMapa(ambiente) {
	var ambiente = JSON.parse(ambiente);

	if (ambiente.LimiteIteracoes) {
		document.getElementById("mapa").innerHTML = "<p>Limite de 5000 iterações atingido.</p>";
	} else if (ambiente.PresasTotais <= 0) {
		document.getElementById("mapa").innerHTML = "<p>Todas as presas foram capturadas.</p>";
	} else {
		var CAgentes = {
			PREDADOR: 1,
			PRESA: 2,
			VAZIO: 3,
			MARCA1: 4,
			MARCA2: 5,
			MARCA3: 6
		};

		var mapa = ambiente.Mapa;

	    var tamanhoMapa = ambiente.TamanhoMapa;
		var table = "<table class=\"table table-bordered table-responsive\">";
		table += "<tbody>";
		for (i = 0; i < tamanhoMapa; i++) {
			table += "<tr>";
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
					case CAgentes.MARCA1:
						marca = "blue"
						break;
					case CAgentes.MARCA2:
						marca = "yellow"
						break;
					case CAgentes.MARCA3:
						marca = "red"
						break;
				}

				var img = imgName != "" ? "<img src=\"../images/" + imgName + "\">" : "";
				var style = marca != "" ? "style='background-color:" + marca + "'" : "";
				table += "<td " + style + ">" + img + "</td>";
			}
			table += "</tr>";
		}
		table += "</tbody>";

		document.getElementById("mapa").innerHTML = table;
	}
}

function iniciarEventSource() {
	var source = new EventSource('/events/');

	source.onmessage = function(e) {
		setMapa(e.data);
	};

	source.onerror = function(e){
		document.getElementById("mapa").innerHTML = "<p>Servidor morreu!</p>";
	};
}

