package main

import (
	"log"
	"net/http"
	// "fmt"
	// "html/template"
	"strconv"
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
	mux.HandleFunc("/addCount", addCount)
	mux.HandleFunc("/minusCount", minusCount)
	mux.HandleFunc("/startApp", startApp)
	port := ":8080"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

func startApp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		info.storeName = r.FormValue("storeName")
		k, _ := strconv.Atoi(r.FormValue("capacity"))
		info.capacity = k
		// fmt.Println("Receive ajax post data string ", ajax_post_data)
		// w.Write([]byte("<h2>after<h2>"))
	}
}

func addCount(w http.ResponseWriter, r *http.Request) {
	info.count++
}

func minusCount(w http.ResponseWriter, r *http.Request) {
	info.count--
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