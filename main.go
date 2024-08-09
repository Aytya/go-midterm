package main

import (
	"context"
	"embed"
	"github.com/joho/godotenv"
	"log"
	"myproject/backend/repository"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		log.Fatalf("Error creating tables: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	app := NewApp(repo)

	r := setupRouter(app)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err.Error())
		}
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shut down: %s", err.Error())
	}
}
