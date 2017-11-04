package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"errors"

	"time"
	"encoding/json"
)

type Client struct {
	dados chan []byte
}

type Broker struct {
	newClients chan Client
}

func NewBroker() (b *Broker) {
	b = &Broker{
		newClients: make(chan Client),
	}
	return
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

	client := Client{}
	client.dados = make(chan []byte)

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
			msg, open := <-client.dados

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

func (b *Broker) AmbienteHandler(w http.ResponseWriter, r *http.Request) {
	errForm, nPresas, nPredadores := obterValoresFormulario(r)
	if errForm != nil {
		http.Redirect(w, r, "/", 200)
		return
	}

	if err := tmplAmbiente.Execute(w, nil); err != nil {
		log.Fatal("Nao foi possivel executar template.", err)
	}

	log.Println("Terminou HTTP request 2 em ", r.URL.Path)

	go func() {
		log.Println("Esperando client..")

		client := <-b.newClients

		log.Println("Iniciou client..")

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
					client.dados <- jsonAmbiente
				}

				if ambienteTela.LimiteIteracoes == true || ambienteTela.PresasTotais == 0  {
					break
				}

				time.Sleep(100 * time.Millisecond)
			}
		}()
	}()
}
