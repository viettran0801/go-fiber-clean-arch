package infrastructure

import (
	"gofiber/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	dns := "host=db user=postgres password=password dbname=testdb port=5432"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Product{})
	return db, nil
}
