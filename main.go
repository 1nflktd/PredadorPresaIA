package main

import (
	"flag"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmplIndex.Execute(w, nil); err != nil {
		log.Fatal("Nao foi possivel executar template.", err)
	}
}

func main() {
	var porta = flag.String("Porta", "8000", "Digite a porta do servidor")
	flag.Parse()

	initTemplates()

	b := NewBroker()

	http.HandleFunc("/", MainHandler)

	http.HandleFunc("/ambiente/", b.AmbienteHandler)

	http.Handle("/events/", b)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	log.Println("Escutando na porta: ", *porta)

	log.Fatal(http.ListenAndServe(":" + *porta, nil))
}