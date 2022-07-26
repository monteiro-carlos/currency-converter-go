package container

import (
	"os"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/service"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/database"
)

//type components struct {
//	Log log.Logger
//}

type Services struct {
	Currency service.ServiceI
}

type Dependency struct {
	//Components components
	Services Services
}

func New() (*Dependency, error) {
	db, err := database.Open(os.Getenv("DSN"))
	if err != nil {
		return nil, err
	}

	repository, err := currency.NewRepository(db)
	if err != nil {
		return nil, err
	}

	currencyService, err := service.NewCurrencyService(
		repository,
	)
	if err != nil {
		return nil, err
	}

	srv := Services{
		Currency: currencyService,
	}

	dep := Dependency{
		Services: srv,
	}

	return &dep, err
}
