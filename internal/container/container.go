package container

import (
	"os"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/adapters/exchangeapi"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/database"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/log"
)

type components struct {
	Log *log.Logger
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

	exchangeClient, err := exchangeapi.NewClient(cmp.Log)
	if err != nil {
		return nil, err
	}

	currencyService, err := currency.NewCurrencyService(
		repository,
		cmp.Log,
		exchangeClient,
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
	logger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	return &components{
		Log: logger,
	}, nil
}
