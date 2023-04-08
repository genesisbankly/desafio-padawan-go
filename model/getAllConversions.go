package model

import (
	"challenge/connDB"
)

type RowExchange struct {
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
	ID            uint32
}

func FetchAllConversions() ([]RowExchange, error) {
	db, err := connDB.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM PadawanDB.conversionsLogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversions []RowExchange
	for rows.Next() {
		var rowExchange RowExchange
		if err := rows.Scan(&rowExchange.ID, &rowExchange.Amount, &rowExchange.From_currency, &rowExchange.To_currency, &rowExchange.Rate); err != nil {
			return nil, err
		}
		conversions = append(conversions, rowExchange)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return conversions, nil
}
