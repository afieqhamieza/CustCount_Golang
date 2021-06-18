package main

import (
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/counter", counter)
	mux.HandleFunc("/settings", settings)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "counter", nil)
}

func settings(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "settings", nil)
}