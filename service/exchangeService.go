package service

import (
	"challenge/model"
)

func GetAllConversions() ([]model.RowExchange, error) {
	conversions, err := model.FetchAllConversions()
	if err != nil {
		return nil, err
	}

	return conversions, nil
}
