package main

import (
	"embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"myproject/backend/repository"
	"net/http"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		log.Fatalf("Error creating tables: %s", err.Error())
	}

	repo := repository.NewRepository(db)

	// Create an instance of the app structure
	app := NewApp(repo)

	// Set up HTTP server
	httpRouter := NewRouter(app)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", httpRouter))
	}()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Todo-list",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	dsn := "host=" + viper.GetString("database.host") +
		" port=" + viper.GetString("database.port") +
		" user=" + viper.GetString("database.user") +
		" password=" + viper.GetString("database.password") +
		" dbname=" + viper.GetString("database.dbname") +
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
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(255) NOT NULL,
	    active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status BOOLEAN DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	return err
}
