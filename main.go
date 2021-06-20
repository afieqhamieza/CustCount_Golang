package main

import (
	"log"
	"net/http"
	// "fmt"
	// "html/template"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render


type Info struct{
	storeName string;
	count int;
	capacity int;
}

var info Info

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}

	rnd = renderer.New(opts)
	info.storeName = "My groceries store"
	info.count = 0
	info.capacity = 200
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/counter", counter)
	mux.HandleFunc("/settings", settings)
	port := ":8080"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
        "storeName": info.storeName,
		"count": info.count,
		"capacity": info.capacity,
    }
	rnd.HTML(w, http.StatusOK, "counter", data)
	
}

func settings(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "settings", nil)
}