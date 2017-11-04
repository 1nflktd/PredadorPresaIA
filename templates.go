package main

import (
	"html/template"
	"log"
)

var tmplIndex *template.Template
var tmplAmbiente *template.Template

func initTemplates() {
	var err error

	tmplIndex, err = template.New("index.html").ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Erro ao parsear template.", err)
	}

	tmplAmbiente, err = template.New("ambiente.html").ParseFiles("templates/ambiente.html")
	if err != nil {
		log.Fatal("Erro ao parsear template.", err)
	}
}