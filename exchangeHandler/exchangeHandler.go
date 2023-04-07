package exchangeHandler

import (
	"challenge/execSQL"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Conversion struct {
	ConvertedAmount float64 `json:"valorConvertido"`
	CurrencySymbol  string  `json:"simboloMoeda"`
}

type registerLog struct {
	amount        float64
	from_currency string
	to_currency   string
	rate          float64
}

// This function receives the URL params and give a response in json with the respective answer or error
func CurrencyConverter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	amount, err := strconv.ParseFloat(params["amount"], 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	from := params["from"]
	to := params["to"]

	rate, err := strconv.ParseFloat(params["rate"], 64)
	if err != nil {
		http.Error(w, "Invalid rate", http.StatusBadRequest)
		return
	}

	var symbol string
	switch to {
	case "USD":
		symbol = "$"
	case "BRL":
		symbol = "R$"
	case "EUR":
		symbol = "€"
	default:
		http.Error(w, "Invalid currency", http.StatusBadRequest)
		return
	}

	convertedAmount := amount * rate
	conversion := Conversion{ConvertedAmount: convertedAmount, CurrencySymbol: symbol}
	registerLog := execSQL.RegisterLog{Amount: amount, From_currency: from, To_currency: to, Rate: rate}

	execSQL.RegisterConversion(registerLog)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversion)
}
