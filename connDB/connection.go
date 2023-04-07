package connDB

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL's connection driver
)

// Connection opens the connection with the database
func Connection() (*sql.DB, error) {
	stringConexao := "root:123456@tcp(127.0.0.1:3008)/PadawanDB?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
