package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"myproject/backend/domain"
	"net/http"
)

func NewRouter(app *App) http.Handler {
	r := chi.NewRouter()

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received POST request to /todos")
		var todo domain.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdTodo, err := app.CreateTodo(r.Context(), todo.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(createdTodo)
	})

	r.Get("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := app.GetTodoByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(todo)
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos, err := app.GetAllTodos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	})

	r.Put("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var todo domain.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo.ID = id
		if err := app.UpdateTodo(r.Context(), &todo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	r.Delete("/todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := app.DeleteTodo(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return r
}
