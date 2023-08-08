package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tesohh/top100things/handler"
	"github.com/Tesohh/top100things/handler/form"
	"github.com/Tesohh/top100things/handler/htmx"
	"github.com/Tesohh/top100things/sb"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	sb.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/things/{userid}", handler.Things).Methods("GET")
	r.HandleFunc("/things/new/form", form.NewThing).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("GET")
	r.HandleFunc("/login/form", form.Login).Methods("POST")
	r.HandleFunc("/signup", handler.Signup).Methods("GET")
	r.HandleFunc("/signup/form", form.Signup).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")

	r.HandleFunc("/htmx/test", htmx.Test).Methods("GET")
	r.HandleFunc("/htmx/thingimage", htmx.ThingImage).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server error")
	}
}
