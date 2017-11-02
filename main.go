package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.New("index.html").ParseFiles("templates/index.html"); err != nil {
		log.Fatal("Erro ao parsear template.", err)
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Fatal("Nao foi possivel executar template.", err)
		}
	}
}

func main() {
	var porta = flag.String("Porta", "8000", "Digite a porta do servidor")
	flag.Parse()

	b := NewBroker()

	http.HandleFunc("/", MainHandler)

	http.Handle("/ambiente/", http.HandlerFunc(b.AmbienteHandler))

	http.Handle("/events/", b)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	log.Fatal(http.ListenAndServe(":" + *porta, nil))
}