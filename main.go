package main

import (
	"log"
	"net/http"

	"github.com/Tesohh/top100things/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server error")
	}
}
