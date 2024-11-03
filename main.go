package main

import (
	"embed"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"myproject/backend/handler"
	"myproject/backend/repository"
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

	repo := repository.NewRepository(db)
	app := handler.NewApp(repo)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Todo-list",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
