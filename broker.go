package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"errors"
	"sync"

	"time"
	"encoding/json"
)

type Broker struct {
	clients map[chan []byte]bool
	newClients chan chan []byte
	defunctClients chan chan []byte
	messages chan []byte
	mutexClients *sync.Mutex
}

func NewBroker() (b *Broker) {
	b = &Broker{
		clients: make(map[chan []byte]bool),
		newClients: make(chan (chan []byte)),
		defunctClients: make(chan (chan []byte)),
		messages: make(chan []byte, 1), // VERIFICAR ESSE BUFFER
		mutexClients: &sync.Mutex{},
	}

	go b.listen()

	return
}

func (b *Broker) listen() {
	for {
		select {

		case s := <-b.newClients:
			b.mutexClients.Lock()
			b.clients[s] = true
			b.mutexClients.Unlock()

			log.Println("Adicionou um client")
		case s := <-b.defunctClients:
			b.mutexClients.Lock()
			delete(b.clients, s)
			b.mutexClients.Unlock()

			log.Println("Removeu um client")
		case msg := <-b.messages:
			b.mutexClients.Lock()
			for s, _ := range b.clients {
				s <- msg
			}
			//log.Printf("Mandando msg para %d clients", len(b.clients))
			b.mutexClients.Unlock()
		}
	}
}

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

	messageChan := make(chan []byte)

	b.newClients <- messageChan

	defer func() {
		b.defunctClients <- messageChan
		log.Println("HTTP conexao fechada.")
	}()

	notify := w.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notify:
			return
		default:
			msg, open := <-messageChan

			if !open {
				break
			}

			fmt.Fprintf(w, "data:%s\n\n", msg)

			f.Flush()
		}
	}

	log.Println("Terminou HTTP request 1 em ", r.URL.Path)
}

func obterValoresFormulario(r *http.Request) (error, int, int) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erro ao obter valores do formulario")
		return err, 0, 0
	}

	if r.PostForm["nPredadores"] == nil || r.PostForm["nPresas"] == nil {
		log.Println("Formulario vazio")
		return errors.New("Formulario vazio"), 0, 0
	}

	nPresas, errPresas := strconv.ParseInt(r.PostForm["nPresas"][0], 10, 32)
	if errPresas != nil {
		log.Println("Erro ao converter valores do formulario.")
		return errPresas, 0, 0
	}

	nPredadores, errPredadores := strconv.ParseInt(r.PostForm["nPredadores"][0], 10, 32)
	if errPredadores != nil {
		log.Println("Erro ao converter valores do formulario.")
		return errPredadores, 0, 0
	}

	return nil, int(nPresas), int(nPredadores)
}

func (b *Broker) MapaHandler(w http.ResponseWriter, r *http.Request) {
	errForm, nPresas, nPredadores := obterValoresFormulario(r)
	if errForm != nil {
		http.Redirect(w, r, "/", 200)
		return
	}

	if tmpl, err := template.New("mapa.html").ParseFiles("templates/mapa.html"); err != nil {
		log.Fatal("Erro ao parsear template.", err)
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Fatal("Nao foi possivel executar template.", err)
		}
	}

	log.Println("Terminou HTTP request 2 em ", r.URL.Path)

	app := &App{}
	app.Init(nPresas, nPredadores)

	go func() {
		app.Run()
	}()

	go func() {
		for {
			ambienteTela := app.ambiente.GetAmbienteTela()
			jsonAmbiente, err := json.Marshal(ambienteTela)
			if err == nil {
				b.messages <- jsonAmbiente
			}

			if ambienteTela.LimiteIteracoes == true || ambienteTela.PresasTotais == 0  {
				break
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()
}
