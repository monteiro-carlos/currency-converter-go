package currency

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/adapters/exchangeapi"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type ServiceI interface {
	AddNewCurrencyManually(currency *models.CurrencyPayload) error
	GetAllCurrencyRates() ([]models.CurrencyPayload, error)
	UpdateCurrenciesDatabase() ([]models.CurrencyPayload, error)
	ConvertValueToAllCurrencies(
		value *models.ConversionRequest,
	) (*[]models.ConversionResponse, error)
	GetCurrencyRatesByCode(code string) (*models.CurrencyPayload, error)
}

type Currency struct {
	repository repository.ServiceI
	Logger     zap.Logger
}

func NewCurrencyService(
	repository repository.ServiceI,
	logger zap.Logger,
) (*Currency, error) {
	if repository == nil {
		return nil, errors.New("repository can't be empty")
	}

	return &Currency{
		repository: repository,
		Logger:     logger,
	}, nil
}

func (c *Currency) AddNewCurrencyManually(currency *models.CurrencyPayload) error {
	input := &repository.CurrencyRate{
		Currency: repository.Currency{
			Code: currency.Currency.Code,
			Name: currency.Currency.Name,
		},
		Rate: currency.Rate,
	}
	if err := c.repository.CreateCurrencyRate(input); err != nil {
		c.Logger.Error("errorMsg", zap.Error(err))
		return err
	}
	c.Logger.Info("AddNewCurrencyManually",
		zap.Any("payload", input))
	return nil
}

func (c *Currency) GetAllCurrencyRates() ([]models.CurrencyPayload, error) {
	currenciesPayload := make([]models.CurrencyPayload, 0)
	cr, err := c.repository.GetAllLast()
	currencyRates := *cr

	for _, rate := range currencyRates {
		currencyModel := &models.CurrencyPayload{
			Currency: models.Currency{
				Code: rate.Currency.Code,
				Name: rate.Currency.Name,
			},
			Rate: rate.Rate,
		}
		currenciesPayload = append(currenciesPayload, *currencyModel)
	}
	if err != nil {
		c.Logger.Error("errorMsg", zap.Error(err))
		return nil, err
	}
	c.Logger.Info("GetAllCurrencyRates",
		zap.Any("payload", currenciesPayload))
	return currenciesPayload, nil
}

func (c *Currency) UpdateCurrenciesDatabase() ([]models.CurrencyPayload, error) {
	currenciesPayloadRep := make([]repository.CurrencyRate, 0)
	rates, err := exchangeapi.GetCurrencyRates()
	if err != nil {
		c.Logger.Error("errorMsg", zap.Error(err))
		return nil, err
	}
	for _, rate := range rates {
		decimalValue, err := decimal.NewFromString(rate.Ask)
		if err != nil {
			c.Logger.Error("errorMsg", zap.Error(err))
			return nil, err
		}
		currencyRep := &repository.CurrencyRate{
			Currency: repository.Currency{
				Code: rate.Code,
				Name: exchangeapi.GetCurrencyName(rate.Name),
			},
			Rate: decimalValue,
		}
		currenciesPayloadRep = append(currenciesPayloadRep, *currencyRep)
		if err := c.repository.CreateCurrencyRate(currencyRep); err != nil {
			c.Logger.Error("errorMsg", zap.Error(err))
			return nil, err
		}
	}
	currenciesPayloadMod := make([]models.CurrencyPayload, 0)
	for _, currencyPayload := range currenciesPayloadRep {
		currencyMod := &models.CurrencyPayload{
			Currency: models.Currency{
				Name: currencyPayload.Currency.Name,
				Code: currencyPayload.Currency.Code,
			},
			Rate: currencyPayload.Rate,
		}
		currenciesPayloadMod = append(currenciesPayloadMod, *currencyMod)
	}
	c.Logger.Info("UpdateCurrenciesDatabase",
		zap.Any("payload", currenciesPayloadMod))

	return currenciesPayloadMod, nil
}

func (c *Currency) ConvertValueToAllCurrencies(
	value *models.ConversionRequest,
) (*[]models.ConversionResponse, error) {
	conversions := make([]models.ConversionResponse, 0)
	cr, err := c.repository.GetAllLast()
	currencyRates := *cr
	if err != nil {
		c.Logger.Error("errorMsg", zap.Error(err))
		return nil, err
	}
	for _, rate := range currencyRates {
		convertedValue := rate.Rate.Mul(value.Value)
		conversion := &models.ConversionResponse{
			Currency: models.Currency{
				Name: rate.Currency.Name,
				Code: rate.Currency.Code,
			},
			Value: convertedValue,
		}
		conversions = append(conversions, *conversion)
	}
	c.Logger.Info("ConvertValueToAllCurrencies",
		zap.Any("payload", conversions))

	return &conversions, nil
}

func (c *Currency) GetCurrencyRatesByCode(code string) (*models.CurrencyPayload, error) {
	currencyRate, err := c.repository.GetLastByCode(code)
	if err != nil {
		c.Logger.Error("errorMsg", zap.Error(err))
		return nil, err
	}
	currencyPayload := &models.CurrencyPayload{
		Currency: models.Currency{
			Name: currencyRate.Currency.Name,
			Code: currencyRate.Currency.Code,
		},
		Rate: currencyRate.Rate,
	}
	c.Logger.Info("GetCurrencyRatesByCode",
		zap.Any("payload", currencyPayload))

	return currencyPayload, nil
}
