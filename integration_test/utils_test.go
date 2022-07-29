package integration_test

import (
	"github.com/gin-gonic/gin"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/adapters/exchangeapi"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/models"
	"github.com/shopspring/decimal"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/database"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/log"
)

const (
	timeout = 5
)

func getCurrencyNameFromCode(code string) string {
	switch code {
	case "USD":
		return "Dólar Americano"
	case "EUR":
		return "Euro"
	case "INR":
		return "Rúpia Indiana"
	default:
		return "test not implemented yet for this currency"
	}
}

func getCurrencyCodesQuantity() int {
	currencies := strings.Split(os.Getenv("CURRENCY_CODES"), ",")
	return len(currencies)
}

func currencyPayload() models.CurrencyPayload {
	return models.CurrencyPayload{
		Currency: models.Currency{
			Code: "USD",
			Name: "Dólar Americano",
		},
		Rate: decimal.NewFromFloat(4.83),
	}
}

func conversionRequest() models.ConversionRequest {
	return models.ConversionRequest{
		Value: decimal.NewFromFloat(500.00),
	}
}

func currencyPayloadRespWithCode(code string) models.CurrencyPayload {
	currencyName := getCurrencyNameFromCode(code)
	return models.CurrencyPayload{
		Currency: models.Currency{
			Code: code,
			Name: currencyName,
		},
		Rate: decimal.NewFromFloat(4.83),
	}
}

type dependencies struct {
	log             *log.Logger
	currencyHandler currency.Handler
	cli             *http.Client
	router          *gin.Engine
}

func setupEnviroment() (*dependencies, error) {
	logger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	db, err := database.Open(os.Getenv("DSN_TEST"))
	if err != nil {
		return nil, err
	}

	routes := routesSetup()

	repository, err := currency.NewRepository(db)
	if err != nil {
		return nil, err
	}

	cmpClient := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	exchangeClient, err := exchangeapi.NewClient(logger)
	if err != nil {
		return nil, err
	}

	currencyService, err := currency.NewCurrencyService(
		repository,
		logger,
		exchangeClient,
	)
	if err != nil {
		return nil, err
	}

	currencyHandler := currency.NewHandler(currencyService, logger)

	return &dependencies{
		log:             logger,
		currencyHandler: *currencyHandler,
		cli:             cmpClient,
		router:          routes,
	}, nil
}

func routesSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	routes := gin.Default()

	return routes
}
