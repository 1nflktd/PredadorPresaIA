<!DOCTYPE html>
<html>
<head>
	<title>Página de teste</title>
	<script type="text/javascript">
		function setMapa(mapa) {
			var table = "<table style=\"border: solid\">";
			for (i = 0; i < 15; i++) {
				table += "<tr>";
				for (j = 0; j < 15; j++) {
					var img = mapa[i][j] == 'a' ? "<img src=\"../images/predator.jpg\">" : mapa[i][j];
					//var img = mapa[i][j] % 2 == 0 ? "X" : " ";
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