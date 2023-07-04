package model

import (
	"challenge/connDB"
	"time"
)

type RowExchange struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

type RowExchangeResponse struct {
	ID            uint    `json:"id"`
	Amount        float64 `json:"amount"`
	From_currency string  `json:"from_currency"`
	To_currency   string  `json:"to_currency"`
	Rate          float64 `json:"rate"`
}

func (RowExchange) TableName() string {
	return "conversionsLogs"
}

func FetchAllConversions() ([]RowExchangeResponse, error) {
	db, err := connDB.Connection()
	if err != nil {
		return nil, err
	}

	var conversions []RowExchange
	err = db.Find(&conversions).Error
	if err != nil {
		return nil, err
	}

	var responseConversions []RowExchangeResponse
	for _, c := range conversions {
		responseConversions = append(responseConversions, RowExchangeResponse{
			ID:            c.ID,
			Amount:        c.Amount,
			From_currency: c.From_currency,
			To_currency:   c.To_currency,
			Rate:          c.Rate,
		})
	}

	return responseConversions, nil
}
