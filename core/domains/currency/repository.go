package currency

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/internal/migration"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) (*repository.Service, error) {
	if err := migration.Migrate(db); err != nil {
		return nil, err
	}

	rep, err := repository.NewService(db)
	if err != nil {
		return nil, err
	}

	return rep, nil
}
