package model

import (
	"challenge/connDB"
	"fmt"
)

type RegisterLog struct {
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

func RegisterConversion(registerLog RegisterLog) error {

	db, erro := connDB.Connection()
	if erro != nil {
		fmt.Println("Error connecting to the database!")
		return erro
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO PadawanDB.conversionsLogs (amount, from_currency, to_currency, rate) values (?, ?, ?, ?)")
	if erro != nil {
		fmt.Println("Error creating the statement!")
		return erro
	}
	defer statement.Close()

	insert, erro := statement.Exec(registerLog.Amount, registerLog.From_currency, registerLog.To_currency, registerLog.Rate)
	if erro != nil {
		fmt.Println("Error executing the statement!")
		return erro
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		fmt.Println("Error obtaining the inserted id!")
		return erro
	}

	fmt.Printf("Conversion record inserted successfully! Id: %d\n", idInsert)
	return nil
}
