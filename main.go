package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/fetchemployee/{id}", fetchEmployeeDetails).Methods("GET")
	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
