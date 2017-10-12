package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"time"
	"encoding/json"
)

type Broker struct {
	clients map[chan []byte]bool
	newClients chan chan []byte
	defunctClients chan chan []byte
	messages chan []byte
}

func (b *Broker) Start() {
	go func() {
		for {

			select {

			case s := <-b.newClients:
				b.clients[s] = true
				log.Println("Added new client")

			case s := <-b.defunctClients:
				delete(b.clients, s)
				close(s)

				log.Println("Removed client")

			case msg := <-b.messages:
				for s, _ := range b.clients {
					s <- msg
				}
				log.Printf("Broadcast message to %d clients", len(b.clients))
			}
		}
	}()
}

func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte)

	b.newClients <- messageChan

	notify := w.(http.CloseNotifier).CloseNotify()

	go func() {
		<-notify

		b.defunctClients <- messageChan
		log.Println("HTTP connection just closed.")
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		msg, open := <-messageChan

		if !open {
			break
		}

		fmt.Fprintf(w, "data:%s\n\n", msg)

		f.Flush()
	}

	log.Println("Finished HTTP request at ", r.URL.Path)
}

func (b *Broker) MapaHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erro ao obter valores do formulario")
		http.Redirect(w, r, "/", 200)
		return
	}

	if r.PostForm["nPredadores"] == nil || r.PostForm["nPresas"] == nil {
		log.Println("Formulario vazio")
		http.Redirect(w, r, "/", 200)
		return
	}

	nPresas, errPresas := strconv.ParseInt(r.PostForm["nPresas"][0], 10, 32)
	if errPresas != nil {
		log.Println("Erro ao converter valores do formulario.")
		http.Redirect(w, r, "/", 200)
		return
	}

	nPredadores, errPredadores := strconv.ParseInt(r.PostForm["nPredadores"][0], 10, 32)
	if errPredadores != nil {
		log.Println("Erro ao converter valores do formulario.")
		http.Redirect(w, r, "/", 200)
		return
	}

	if tmpl, err := template.New("mapa.html").ParseFiles("templates/mapa.html"); err != nil {
		log.Fatal("Error parsing your template.", err)
	} else {
		if err = tmpl.Execute(w, "hnrq"); err != nil {
			log.Fatal("unable to execute template.", err)
		}
	}

	b.Start()

	ch := make(chan bool)

	app := &App{}

	go func() {
		app.Run(int(nPresas), int(nPredadores))
	}()

	go func() {
		for {
			jsonAmbiente, err := json.Marshal(app.ambiente.getMapa())
			if err == nil {
				b.messages <- jsonAmbiente
			}

			if len(b.clients) == 0 {
				ch <- true
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	<-ch

	log.Println("Finished HTTP request at ", r.URL.Path)
}
