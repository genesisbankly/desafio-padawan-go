package model

import (
	"challenge/connDB"
)

type RegisterLog struct {
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

func RegisterConversion(registerLog RegisterLog) (int64, error) {
	db, err := connDB.Connection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO PadawanDB.conversionsLogs (amount, from_currency, to_currency, rate) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	insert, err := statement.Exec(registerLog.Amount, registerLog.From_currency, registerLog.To_currency, registerLog.Rate)
	if err != nil {
		return 0, err
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		return 0, err
	}

	return idInsert, nil
}
