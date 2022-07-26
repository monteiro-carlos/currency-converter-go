package repository

import (
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type ServiceI interface {
	CreateCurrencyRate(currencyRate *CurrencyRate) error
	GetLastByCode(currencyCode string) (*CurrencyRate, error)
	GetLastByName(currencyName string) (*CurrencyRate, error)
	GetAllLast() (*[]CurrencyRate, error)
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateCurrencyRate(currencyRate *CurrencyRate) error {
	return r.db.Create(currencyRate).Error
}

func (r *Repository) GetLastByCode(currencyCode string) (*CurrencyRate, error) {
	currencyRate := &CurrencyRate{Currency: Currency{Code: currencyCode}}
	res := r.db.Order("created_at").Where(currencyRate).First(currencyRate)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRate, nil
}

func (r *Repository) GetLastByName(currencyName string) (*CurrencyRate, error) {
	currencyRate := &CurrencyRate{Currency: Currency{Name: currencyName}}
	res := r.db.Order("created_at").Where(currencyRate).First(currencyRate)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRate, nil
}

func (r *Repository) GetAllLast() (*[]CurrencyRate, error) {
	currencyRates := &[]CurrencyRate{}
	res := r.db.Order("created_at").Distinct("currency_code").Find(currencyRates)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRates, nil
}
