package model

import (
	"challenge/connDB"
)

type RegisterLog struct {
	ID            uint32 `gorm:"primaryKey"`
	Amount        float64
	From_currency string
	To_currency   string
	Rate          float64
}

func (RegisterLog) TableName() string {
	return "conversionsLogs"
}

func RegisterConversion(registerLog RegisterLog) (uint32, error) {
	db, err := connDB.Connection()
	if err != nil {
		return 0, err
	}

	err = db.Create(&registerLog).Error
	if err != nil {
		return 0, err
	}

	return registerLog.ID, nil
}
