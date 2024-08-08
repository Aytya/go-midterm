package main

import (
	"context"
	"embed"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"myproject/backend/domain"
	"myproject/backend/repository"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/viper"
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
	app := NewApp(repo)

	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
	})

	r.Use(c.Handler)

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := app.GetTodoByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(todo)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		todos, err := app.GetAllTodos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var todo domain.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdTodo, err := app.CreateTodo(r.Context(), todo.Title, todo.Date, todo.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(createdTodo)
	})

	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var todo domain.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedTodo, err := app.UpdateTodo(r.Context(), id, todo.Title, todo.Date, todo.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedTodo)
	})

	r.Put("/check/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := app.CheckTodo(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := app.DeleteTodo(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

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

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

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
        active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        status BOOLEAN DEFAULT FALSE
    );`
	_, err := db.Exec(query)
	return err
}
