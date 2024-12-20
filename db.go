package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"myproject/backend/domain"
	"os"
)

func initDB() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.Todo{}, &domain.User{})
	return db, nil
}
