function setMapa(ambiente) {
	var ambiente = JSON.parse(ambiente);

	var CAgentes = {
		PREDADOR: 1,
		PRESA: 2,
		VAZIO: 3,
		MARCA1: 4,
		MARCA2: 5,
		MARCA3: 6,
		PRESA_FUGINDO: 7
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
				case CAgentes.PRESA_FUGINDO:
					imgName = "homer_alt.jpg";
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
		setMapa(e.data);
	};

	source.onerror = function(e){
		$(".msgs").hide();
		$("#servidor-morto").show();
	};
}

