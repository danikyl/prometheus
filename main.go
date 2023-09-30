package main

import (
	"log"
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Hello, " + r.URL.Query().Get("name")))
}

func main() {
	http.HandleFunc("/hello", basicHandler)
	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
