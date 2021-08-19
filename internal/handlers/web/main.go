package web

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type MainHandler struct {}

func NewMainHandler() MainHandler {
	return MainHandler{}
}

func (hh *MainHandler) homePage(writer http.ResponseWriter, request *http.Request) {
	absTemplatePath, err := filepath.Abs("./web/templates/main.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles(absTemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (hh *MainHandler) RegisterRoutes (router *mux.Router) {
	router.HandleFunc("/", hh.homePage).Methods("GET")
}