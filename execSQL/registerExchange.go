package execSQL

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

func RegisterConversion(registerLog RegisterLog) {

	db, erro := connDB.Connection()
	if erro != nil {
		fmt.Println("Erro ao conectar no banco de dados!")
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO PadawanDB.conversionsLogs (amount, from_currency, to_currency, rate) values (?, ?, ?, ?)")
	if erro != nil {
		fmt.Println("Erro ao criar o statement!")
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(registerLog.Amount, registerLog.From_currency, registerLog.To_currency, registerLog.Rate)
	if erro != nil {
		fmt.Println("Erro ao executar o statement!")
		return
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		fmt.Println("Erro ao obter o id inserido!")
		return
	}

	fmt.Println("Registro de conversão inserido com sucesso! Id: $d", idInsert)
}
