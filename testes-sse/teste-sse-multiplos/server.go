package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Client struct {
	client chan string
}

type Broker struct {
	newClients chan Client
}

var b *Broker

var tmplEventos *template.Template

func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming nao suportado!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	client := Client{}
	client.client = make(chan string)

	b.newClients <- client

	defer func() {
		log.Println("HTTP conexao fechada.")
	}()

	notify := w.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notify:
			return
		default:
			msg, open := <-client.client

			if !open {
				break
			}

			fmt.Fprintf(w, "data:%s\n\n", msg)

			f.Flush()
		}
	}

	log.Println("Terminou HTTP request 1 em ", r.URL.Path)
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erro ao obter valores do formulario")
		return
	}

	if r.PostForm["nome"] == nil {
		log.Println("Formulario vazio")
		return
	}

	tmplEventos.Execute(w, r.PostForm["nome"])

	log.Println("Finished HTTP request 2 at ", r.URL.Path)

	log.Println("Main handler")

	go func() {
		log.Println("Esperando client..")
		client := <-b.newClients
		log.Println("Iniciou client..")
		log.Printf("%+v\n", client)
		for i := 0; ; i++ {
			client.client <- fmt.Sprintf("%d - the time is %v", i, time.Now())

			log.Printf("Sent message %d ", i)

			if i == 100 {
				log.Printf("entrou no break")
				break
			}

			time.Sleep(300 * time.Millisecond)
		}
	}()
}

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
	b = &Broker{
		make(chan Client),
	}

	var errTmp error
	tmplEventos, errTmp = template.ParseFiles("templates/eventos.html")
	if errTmp != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	http.Handle("/events/", b)

	http.HandleFunc("/", MainHandler)

	http.Handle("/eventos/", http.HandlerFunc(MainPageHandler))

	http.ListenAndServe(":8000", nil)
}
