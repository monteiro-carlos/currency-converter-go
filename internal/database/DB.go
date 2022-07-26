package database

import (
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DbConnection(uri string) error {
	DB, err := gorm.Open(postgres.Open(uri))
	if err != nil {
		log.Panic("database connection error")
		return err
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	err = DB.AutoMigrate(&repository.CurrencyRate{})
	if err != nil {
		log.Panic("can't execute migration")
		return err
	}
	return nil
}
