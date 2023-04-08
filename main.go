package main

import (
	"challenge/controller"
	"challenge/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("public"))

	router.HandleFunc("/exchange/{amount}/{from}/{to}/{rate}", controller.CurrencyConverter).Methods(http.MethodGet)
	router.HandleFunc("/exchange", model.GetAllConversions).Methods(http.MethodGet)
	router.Handle("/", fs).Methods(http.MethodGet)

	fmt.Println("Executando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
