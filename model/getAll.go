package model

import (
	"challenge/connDB"
	"encoding/json"
	"net/http"
)

type RowExchange struct {
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
	ID            uint32
}

func GetAllConversions(w http.ResponseWriter, r *http.Request) {

	db, erro := connDB.Connection()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	rows, erro := db.Query("SELECT * FROM PadawanDB.conversionsLogs")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os registros das conversões"))
		return
	}
	defer rows.Close()

	var conversions []RowExchange
	for rows.Next() {
		var rowExchange RowExchange
		if erro := rows.Scan(&rowExchange.ID, &rowExchange.Amount, &rowExchange.From_currency, &rowExchange.To_currency, &rowExchange.Rate); erro != nil {
			w.Write([]byte("Erro na estrutura de alguma linha"))
			return
		}
		conversions = append(conversions, rowExchange)
	}

	if err := rows.Err(); err != nil {
		w.Write([]byte("Erro ao iterar pelos registros das conversões"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(conversions); erro != nil {
		w.Write([]byte("Erro ao converter os registros das conversões em JSON!"))
		return
	}
}
