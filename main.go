package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
)

var tmpl *template.Template

func executeTemplate(w http.ResponseWriter, ambiente AmbienteTela) {
	if err := tmpl.ExecuteTemplate(w, "script", ambiente); err != nil {
		log.Println("unable to execute template.", err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if tmpl, err = template.New("main.tpl").ParseFiles("templates/main.tpl"); err != nil {
		log.Println("unable to parse main template.", err)
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Println("unable to execute template.", err)
		}
	}

	nPredadores := flag.Int("predadores", 10, "Numero de predadores")
	nPresas := flag.Int("presas", 10, "Numero de presas")
	flag.Parse()

	app := &App{}
	app.Run(w, *nPresas, *nPredadores)
}

func main() {
	var porta = flag.String("Porta", "8000", "Digite a porta do servidor")
	flag.Parse()

	http.HandleFunc("/", mainHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	log.Fatal(http.ListenAndServe(":" + *porta, nil))
}