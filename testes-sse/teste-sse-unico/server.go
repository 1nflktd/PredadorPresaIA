// Golang HTML5 Server Side Events Example
//
// Run this code like:
//  > go run server.go
//
// Then open up your browser to http://localhost:8000
// Your browser must support HTML5 SSE, of course.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"strconv"
)

const TAMANHO_MAPA = 5

// A single Broker will be created in this program. It is responsible
// for keeping a list of which clients (browsers) are currently attached
// and broadcasting events (messages) to those clients.
//
type Broker struct {

	// Create a map of clients, the keys of the map are the channels
	// over which we can push messages to attached clients.  (The values
	// are just booleans and are meaningless.)
	//
	clients map[chan []byte]bool

	// Channel into which new clients can be pushed
	//
	newClients chan chan []byte

	// Channel into which disconnected clients should be pushed
	//
	defunctClients chan chan []byte

	// Channel into which messages are pushed to be broadcast out
	// to attahed clients.
	//
	messages chan []byte
}

// This Broker method starts a new goroutine.  It handles
// the addition & removal of clients, as well as the broadcasting
// of messages out to clients that are currently attached.
//
func (b *Broker) Start() {

	// Start a goroutine
	//
	go func() {

		// Loop endlessly
		//
		for {

			// Block until we receive from one of the
			// three following channels.
			select {

			case s := <-b.newClients:

				// There is a new client attached and we
				// want to start sending them messages.
				b.clients[s] = true
				log.Println("Added new client")

			case s := <-b.defunctClients:

				// A client has dettached and we want to
				// stop sending them messages.
				delete(b.clients, s)
				close(s)

				log.Println("Removed client")

			case msg := <-b.messages:

				// There is a new message to send.  For each
				// attached client, push the new message
				// into the client's message channel.
				for s, _ := range b.clients {
					s <- msg
				}
				log.Printf("Broadcast message to %d clients", len(b.clients))
			}
		}
	}()
}

// This Broker method handles and HTTP request at the "/events/" URL.
//
func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Make sure that the writer supports flushing.
	//
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Create a new channel, over which the broker can
	// send this client messages.
	messageChan := make(chan []byte)

	// Add this client to the map of those that should
	// receive updates
	b.newClients <- messageChan

	// Listen to the closing of the http connection via the CloseNotifier
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		// Remove this client from the map of attached clients
		// when `EventHandler` exits.
		b.defunctClients <- messageChan
		log.Println("HTTP connection just closed.")
	}()

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		// Read from our messageChan.
		msg, open := <-messageChan

		if !open {
			break
		}

		// Write to the ResponseWriter, `w`.
		fmt.Fprintf(w, "data:%s\n\n", msg)

		f.Flush()
	}

	// Done.
	log.Println("Finished HTTP request at ", r.URL.Path)
}

func (b *Broker) MapaHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erro ao obter valores do formulario")
		return
	}

	if r.PostForm["nPredadores"] == nil || r.PostForm["nPresas"] == nil {
		http.Redirect(w, r, "/", 200)
		return
	}

	nPredadores, errPredadores := strconv.ParseInt(r.PostForm["nPredadores"][0], 10, 32)
	if errPredadores != nil {
		log.Println("Erro ao converter valores do formulario.")
		http.Redirect(w, r, "/", 200)
		return
	}

	nPresas, errPresas := strconv.ParseInt(r.PostForm["nPresas"][0], 10, 32)
	if errPresas != nil {
		log.Println("Erro ao converter valores do formulario.")
		http.Redirect(w, r, "/", 200)
		return
	}

	log.Printf("Formulario\npredadores: %d\npresas: %d\n", nPredadores, nPresas)

	if tmpl, err := template.New("mapa.html").ParseFiles("templates/mapa.html"); err != nil {
		log.Fatal("WTF dude, error parsing your template.", err)
	} else {
		// Render the template, writing to `w`.
		if err = tmpl.Execute(w, "hnrq"); err != nil {
			log.Fatal("unable to execute template.", err)
		}
	}

	// Start processing events
	b.Start()

	ch := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			// Create a little message to send to clients,
			// including the current time.
			mapa := [TAMANHO_MAPA][TAMANHO_MAPA]int{}
			for im := 0; im < TAMANHO_MAPA; im++ {
				for j := 0; j < TAMANHO_MAPA; j++ {
					mapa[im][j] = im + j + time.Now().Second()
				}
			}

			jsonMapa, err := json.Marshal(mapa)
			if err == nil {
				b.messages <- jsonMapa
			}

			if len(b.clients) == 0 {
				ch <- true
			}

			// Print a nice log message and sleep for 5s.
			log.Printf("Sent message %d ", i)
			time.Sleep(1 * 1e9)
		}
	}()

	<-ch

	// Done.
	log.Println("Finished HTTP request at ", r.URL.Path)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.New("index.html").ParseFiles("templates/index.html"); err != nil {
		log.Println("unable to parse main template.", err)
	} else {
		if err = tmpl.Execute(w, nil); err != nil {
			log.Println("unable to execute template.", err)
		}
	}
}

func main() {

	// Make a new Broker instance
	b := &Broker{
		make(map[chan []byte]bool),
		make(chan (chan []byte)),
		make(chan (chan []byte)),
		make(chan []byte),
	}

	http.Handle("/events/", b)

	http.Handle("/mapa/", http.HandlerFunc(b.MapaHandler))

	http.Handle("/", http.HandlerFunc(MainHandler))

	// Start the server and listen forever on port 8000.
	http.ListenAndServe(":8000", nil)
}
