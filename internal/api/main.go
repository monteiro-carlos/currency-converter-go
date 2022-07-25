package main

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/database"
)

func main() {
	if err := database.DbConnection(); err != nil {
		return
	}

	//var currencies []models.Currency
	//currencies = append(currencies, models.CurrencyRate{
	//	Currency: models.Currency{"Estados Unidos", "USD"},
	//	Rate:     decimal.NewFromFloat(5.395398554413112),
	//}, models.CurrencyRate{
	//	Currency: models.Currency{Name: "Países da União Europeia", Code: "EUR"},
	//	Rate:     decimal.NewFromFloat(6.365481623828969),
	//}, models.CurrencyRate{
	//	Currency: models.Currency{Name: "Índia", Code: "INR"},
	//	Rate:     decimal.NewFromFloat(0.0724135905111813),
	//})
	//
	//givenValue := decimal.NewFromFloat(529.99)
	//
	//for _, currency := range currencies {
	//	result := models.ConversionResponse{
	//		Currency: currency.Currency, Value: givenValue.DivRound(currency.Rate, 2),
	//	}
	//
	//	fmt.Printf("%v: %v (%v)\n", result.Currency.Code, result.Value, result.Currency.Name)
	//}

}
