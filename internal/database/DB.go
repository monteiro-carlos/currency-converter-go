package database

import (
	"log"
	"os"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection() error {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DSN")))
	if err != nil {
		log.Panic("database connection error")
		return err
	}
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	err := DB.AutoMigrate(&repository.CurrencyRate{})
	if err != nil {
		log.Panic("can't execute migration")
		return err
	}
	return nil
}
