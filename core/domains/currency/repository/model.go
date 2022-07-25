package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Currency struct {
	Code string `gorm:"index"`
	Name string
}

type CurrencyRate struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Currency  Currency  `gorm:"embedded;embeddedPrefix:currency_"`
	Rate      decimal.Decimal
	CreatedAt time.Time
}
