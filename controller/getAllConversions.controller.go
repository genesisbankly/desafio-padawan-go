package controller

import (
	"challenge/service"
	"encoding/json"
	"net/http"
)

func GetAllConversions(w http.ResponseWriter, r *http.Request) {
	conversions, err := service.GetAllConversions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching conversions"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversions); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding conversions to JSON"))
	}
}
