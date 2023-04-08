package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ValidateParams(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		_, err := strconv.ParseFloat(params["amount"], 64)
		if err != nil {
			http.Error(w, "Invalid amount", http.StatusBadRequest)
			return
		}

		from := params["from"]
		to := params["to"]

		allowedCurrencies := map[string]bool{
			"BRL": true,
			"USD": true,
			"EUR": true,
		}

		if !allowedCurrencies[from] || !allowedCurrencies[to] {
			http.Error(w, "Invalid currency code. Allowed codes are BRL, USD, and EUR.", http.StatusBadRequest)
			return
		}

		_, err = strconv.ParseFloat(params["rate"], 64)
		if err != nil {
			http.Error(w, "Invalid rate", http.StatusBadRequest)
			return
		}

		next(w, r) // Call the next middleware or handler in the chain
	}
}
