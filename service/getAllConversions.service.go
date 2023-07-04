package service

import (
	"challenge/model"
)

func GetAllConversions() ([]model.RowExchangeResponse, error) {
	conversions, err := model.FetchAllConversions()
	if err != nil {
		return nil, err
	}

	return conversions, nil
}
