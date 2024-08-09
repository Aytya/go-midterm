package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
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

func (a *App) CreateTodo(ctx context.Context, title, date, todoTime, priority string) (*domain.Todo, error) {
	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if date == "" || todoTime == "" || priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	id := uuid.New().String()
	activeAt := time.Now()

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		Date:     date,
		Time:     todoTime,
		ActiveAt: activeAt,
		Status:   false,
		Priority: priority,
	}

	err := a.todoRepository.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) UpdateTodo(ctx context.Context, id, title, date, todoTime, priority string) (*domain.Todo, error) {
	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if date == "" || todoTime == "" || priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		Date:     date,
		Time:     todoTime,
		Priority: priority,
	}

	err := a.todoRepository.Update(ctx, todo)
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

func (a *App) DeleteTodo(ctx context.Context, id string) error {
	return a.todoRepository.Delete(ctx, id)
}

func (a *App) CheckTodo(ctx context.Context, id string) error {
	return a.todoRepository.CheckTodo(ctx, id)
}
