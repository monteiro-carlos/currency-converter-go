package container

import (
	"os"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/database"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type components struct {
	Log zap.Logger
}

type Services struct {
	Currency currency.ServiceI
	Health   health.ServiceI
}

type Dependency struct {
	Components components
	Services   Services
}

func New() (*Dependency, error) {
	cmp, err := setupComponents()
	if err != nil {
		return nil, err
	}

	db, err := database.Open(os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	repository, err := currency.NewRepository(db)
	if err != nil {
		return nil, err
	}

	currencyService, err := currency.NewCurrencyService(
		repository,
		cmp.Log,
	)
	if err != nil {
		return nil, err
	}

	healthService, err := health.NewService(
		repository,
		cmp.Log,
	)
	if err != nil {
		return nil, err
	}

	srv := Services{
		Currency: currencyService,
		Health:   healthService,
	}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return &dep, err
}

func setupComponents() (*components, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &components{
		Log: *logger,
	}, nil
}
