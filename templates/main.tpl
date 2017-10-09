<!DOCTYPE html>
<html>
<head>
	<title>Página de teste</title>
	<script type="text/javascript">
		function setMapa(mapa) {
			var table = "<table style=\"border: solid\">";
			for (i = 0; i < 30; i++) {
				table += "<tr>";
				for (j = 0; j < 30; j++) {
					var imgName = "";
					switch(mapa[i][j]) {
						case 1:
							imgName = "predator.jpg";
							break;
						case 2:
							imgName = "homer.jpg";
							break;
						case 3:
							imgName = "fundo.jpg";
							break;
					}

					var img =  "<img src=\"../images/" + imgName + "\">";
					table += "<td style=\"border: solid\">" + img + "</td>";
				}
				table += "</tr>";
			}

			document.getElementById("mapa").innerHTML = table;
		}

	</script>
	{{ define "script" }}
	<script type="text/javascript">
		setMapa({{ .Mapa }});
	</script>
	{{ end }}
</head>
<body>
	<div id="mapa"></div>
</body>
</html>