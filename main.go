package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tesohh/top100things/handler"
	"github.com/Tesohh/top100things/sb"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	sb.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server error")
	}
}
