package controller

import (
	"challenge/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// This function receives the URL params and give a response in json with the respective answer or error
func CurrencyConverter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	conversion, err := service.ExchangeHandlerService(params)
	if err != nil {
		http.Error(w, "Failed to process conversion request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversion); err != nil {
		http.Error(w, "Failed to encode conversion response", http.StatusInternalServerError)
	}
}
