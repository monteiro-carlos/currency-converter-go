package currency

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
)

type ServiceI interface {
	AddNewCurrencyManually(currency *models.CurrencyPayload) error
	GetAllCurrencyRates() ([]models.CurrencyPayload, error)
}

type Currency struct {
	Repository repository.ServiceI
}

func (c *Currency) AddNewCurrencyManually(currency *models.CurrencyPayload) error {
	input := &repository.CurrencyRate{
		Currency: repository.Currency{
			Code: currency.Currency.Code,
			Name: currency.Currency.Name,
		},
		Rate: currency.Rate,
	}
	if err := c.Repository.CreateCurrencyRate(input); err != nil {
		return err
	}
	return nil
}

func (c *Currency) GetAllCurrencyRates() ([]models.CurrencyPayload, error) {
	var currencyPayload []models.CurrencyPayload
	cr, err := c.Repository.GetAllLast()
	currencyRates := *cr

	for _, rate := range currencyRates {
		currency := &models.CurrencyPayload{
			Currency: models.Currency{
				Code: rate.Currency.Code,
				Name: rate.Currency.Name,
			},
			Rate: rate.Rate,
		}
		currencyPayload = append(currencyPayload, *currency)
	}
	if err != nil {
		return nil, err
	}
	return currencyPayload, nil
}
