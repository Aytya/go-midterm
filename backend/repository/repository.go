package repository

import (
	"context"
	"gorm.io/gorm"
	"myproject/backend/domain"
)

type Todo interface {
	Create(ctx context.Context, todo *domain.Todo) error
	GetByID(ctx context.Context, id string) (*domain.Todo, error)
	GetAll(ctx context.Context) ([]*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id string) error
	CheckTodo(ctx context.Context, id string) error
}

type Repository struct {
	Todo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Todo: NewTodoRepository(db),
	}
}
