package main

import (
	"context"
	"github.com/google/uuid"
	"myproject/backend/domain"
	"myproject/backend/repository"
	"time"
)

// App struct
type App struct {
	ctx            context.Context
	todoRepository repository.Todo
}

// NewApp creates a new App application struct
func NewApp(todoRepo repository.Todo) *App {
	return &App{
		todoRepository: todoRepo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CreateTodo(ctx context.Context, title string) (*domain.Todo, error) {
	id := uuid.New().String()

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		ActiveAt: time.Now(),
		Status:   false,
	}

	err := a.todoRepository.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) GetTodoByID(ctx context.Context, id string) (*domain.Todo, error) {
	return a.todoRepository.GetByID(ctx, id)
}

func (a *App) GetAllTodos(ctx context.Context) ([]*domain.Todo, error) {
	return a.todoRepository.GetAll(ctx)
}

func (a *App) UpdateTodo(ctx context.Context, todo *domain.Todo) error {
	return a.todoRepository.Update(ctx, todo)
}

func (a *App) DeleteTodo(ctx context.Context, id string) error {
	return a.todoRepository.Delete(ctx, id)
}
