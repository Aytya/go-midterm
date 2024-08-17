package handler

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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CreateTodo(title, priority string, dateTime time.Time) (*domain.Todo, error) {
	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	id := uuid.New().String()
	activeAt := time.Now().UTC()

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		DateTime: dateTime,
		ActiveAt: activeAt,
		Status:   false,
		Priority: priority,
	}

	err := a.todoRepository.Create(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) UpdateTodo(id, title, priority string, dateTime time.Time) (*domain.Todo, error) {
	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		DateTime: dateTime,
		Priority: priority,
	}

	err := a.todoRepository.Update(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) GetTodoByID(id string) (*domain.Todo, error) {
	return a.todoRepository.GetByID(context.Background(), id)
}

func (a *App) GetAllTodos() ([]*domain.Todo, error) {
	ctx := context.Background()
	todos, err := a.todoRepository.GetAll(ctx)
	if err != nil {
		log.Printf("Error retrieving todos: %v", err)
		return nil, err
	}
	if todos == nil {
		log.Println("No todos found, returning empty slice.")
		return []*domain.Todo{}, nil
	}
	log.Println("Retrieved todos:", todos)
	return todos, nil
}

func (a *App) DeleteTodo(id string) error {
	return a.todoRepository.Delete(context.Background(), id)
}

func (a *App) CheckTodo(id string) error {
	return a.todoRepository.CheckTodo(context.Background(), id)
}
