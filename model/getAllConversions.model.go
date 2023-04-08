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

func (RowExchange) TableName() string {
	return "conversionsLogs"
}

func FetchAllConversions() ([]RowExchange, error) {
	db, err := connDB.Connection()
	if err != nil {
		return nil, err
	}

	var conversions []RowExchange
	err = db.Find(&conversions).Error
	if err != nil {
		return nil, err
	}

	return conversions, nil
}
