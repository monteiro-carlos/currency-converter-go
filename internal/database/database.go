package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type ServiceI interface {
	Open(uri string) (db *gorm.DB, err error)
}

func Open(uri string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open(uri))
	if err != nil {
		log.Panic("database connection error")
		return nil, err
	}
	return DB, nil
}
