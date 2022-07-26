package migration

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	if err := db.AutoMigrate(
		&repository.CurrencyRate{},
	); err != nil {
		return errors.Wrap(err, "can't execute migration")
	}

	return nil
}
