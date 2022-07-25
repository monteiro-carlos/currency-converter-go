package models

import "github.com/shopspring/decimal"

type Currency struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ConversionResponse struct {
	Currency Currency        `json:"currency"`
	Value    decimal.Decimal `json:"value"`
}

type ConversionRequest struct {
	Value decimal.Decimal `json:"value"`
}
