# PredadorPresaIA

Valores utilizados no programa:

- Presa:
    - Velocidade máxima: 2
    - Quantidade de iterações com velocidade máxima: 8
    - Quantidade de iterações em pânico (depois de fugir) para ficar "calma": 10
    - Campo de visão: 1 casa

- Predador:
    - Velocidade máxima: 4
    - Quantidade de iterações com velocidade máxima: 4
    - Multiplicador de intensidade das marcas: 5 (3 tipos de marcas: 11 a 15, 6 a 10, 1 a 5)
    - Campo de visão: 4 casas

- Ambiente:
    - Tamanho do mapa: 30x30
    - Limite de iterações: 5000


- Windows:

	Para rodar basta clicar no arquivo PredadorPresaIA.exe. O programa executará com a porta padrão 8000.

	Caso precise mudá-la, basta entrar pelo console (cmd) e rodar o executável com a opção Porta. Ex:

		C:\go\src\PredadorPresaIA\PredadorPresaIA.exe -Porta 8001

	Após, abrir no navegador o endereço "http://localhost:8000" (ou a porta que for utilizada).

	Para compilar, é necessário instalar o Go (https://golang.org/doc/install?download=go1.9.windows-amd64.msi).
	Depois, rodar (na pasta onde se encontra os arquivos):

		C:\go\src\PredadorPresaIA\go build
