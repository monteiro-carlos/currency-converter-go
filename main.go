package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	type Currency struct {
		Name string
		Code string
	}

	type CurrencyRate struct {
		Currency Currency
		Rate     decimal.Decimal
	}

	type CurrencyValue struct {
		Currency Currency
		Value    decimal.Decimal
	}
	var currencies []CurrencyRate
	currencies = append(currencies, CurrencyRate{
		Currency{"Estados Unidos", "USD"}, decimal.NewFromFloat(5.395398554413112),
	}, CurrencyRate{
		Currency{"Países da União Europeia", "EUR"}, decimal.NewFromFloat(6.365481623828969),
	}, CurrencyRate{
		Currency{"Índia", "INR"}, decimal.NewFromFloat(0.0724135905111813),
	})

	givenValue := decimal.NewFromFloat(529.99)

	for _, currency := range currencies {
		result := CurrencyValue{
			currency.Currency, givenValue.DivRound(currency.Rate, 2),
		}

		fmt.Printf("%v: %v (%v)\n", result.Currency.Code, result.Value, result.Currency.Name)
	}

}
