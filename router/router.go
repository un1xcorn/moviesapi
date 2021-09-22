package router

import (
	"html/template"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func Init() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/docs", DocsHandler).Methods("GET")
	r.HandleFunc("/movie", MovieHandler).Methods("GET")
	r.HandleFunc("/movies", MoviesHandler).Methods("GET")

	return r

}

func GetTemplate(name string) *template.Template {

	tmpl, err := template.ParseFiles(name)
	if err != nil {
		log.Fatal("An error happened while parsing html template.", name)
	}
	return tmpl

}
