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

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) (*Service, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}

	return &Service{
		db: db,
	}, nil
}

func (s *Service) CreateCurrencyRate(currencyRate *CurrencyRate) error {
	return s.db.Debug().Create(currencyRate).Error
}

func (s *Service) GetLastByCode(currencyCode string) (*CurrencyRate, error) {
	currencyRate := &CurrencyRate{Currency: Currency{Code: currencyCode}}
	res := s.db.Order("created_at desc").Where(currencyRate).First(currencyRate)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRate, nil
}

func (s *Service) GetLastByName(currencyName string) (*CurrencyRate, error) {
	currencyRate := &CurrencyRate{Currency: Currency{Name: currencyName}}
	res := s.db.Order("created_at desc").Where(currencyRate).First(currencyRate)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRate, nil
}

func (s *Service) GetAllLast() (*[]CurrencyRate, error) {
	currencyRates := &[]CurrencyRate{}
	res := s.db.Select("DISTINCT ON (currency_code) currency_code", "currency_name", "rate", "created_at").
		Order("currency_code").Order("created_at desc").Find(&currencyRates)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, "can't execute find")
	}
	return currencyRates, nil
}

func (s *Service) UpdateCurrencyRatesOnline(currencyRate *CurrencyRate) error {
	return s.db.Create(currencyRate).Error
}
