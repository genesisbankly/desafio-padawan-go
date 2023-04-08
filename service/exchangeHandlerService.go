package service

import (
	"challenge/model"
	"strconv"
)

type Conversion struct {
	ConvertedAmount float64 `json:"valorConvertido"`
	CurrencySymbol  string  `json:"simboloMoeda"`
}

func ExchangeHandlerService(params map[string]string) (Conversion, error) {
	amount, _ := strconv.ParseFloat(params["amount"], 64)
	from := params["from"]
	to := params["to"]
	rate, _ := strconv.ParseFloat(params["rate"], 64)

	var symbol string
	switch to {
	case "USD":
		symbol = "$"
	case "BRL":
		symbol = "R$"
	case "EUR":
		symbol = "€"
	}

	convertedAmount := amount * rate
	conversion := Conversion{ConvertedAmount: convertedAmount, CurrencySymbol: symbol}
	registerLog := model.RegisterLog{Amount: amount, From_currency: from, To_currency: to, Rate: rate}

	err := model.RegisterConversion(registerLog)
	if err != nil {
		return Conversion{}, err
	}

	return conversion, nil
}
