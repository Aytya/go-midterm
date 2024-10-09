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

	//db, err := sqlx.Connect("postgres", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.Todo{})
	return db, nil
}

//func createTables(db *sqlx.DB) error {
//	query := `
//    CREATE TABLE IF NOT EXISTS todos (
//        id UUID PRIMARY KEY,
//        title VARCHAR(255) NOT NULL,
//        datetime TIMESTAMPTZ,
//        priority VARCHAR(255) NOT NULL,
//        active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//        status BOOLEAN DEFAULT FALSE
//    );`
//	_, err := db.Exec(query)
//	return err
//}
