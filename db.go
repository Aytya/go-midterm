package main

import (
	"github.com/jmoiron/sqlx"
	"os"
)

func initDB() (*sqlx.DB, error) {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sqlx.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id UUID PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        date DATE,
        time TIME,
        priority VARCHAR(255) NOT NULL,
        active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        status BOOLEAN DEFAULT FALSE
    );`
	_, err := db.Exec(query)
	return err
}
