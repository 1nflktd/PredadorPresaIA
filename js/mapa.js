function setMapa(mapa, tamanhoMapa, iteracaoAtual) {
	var styleTd = "height: 25px; width: 25px; padding: 0px;";

	var table = "<table class=\"table table-bordered table-responsive\">";
		table+= "<caption>Iteração atual: " + iteracaoAtual + "</caption>";
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
}
