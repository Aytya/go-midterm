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

type User interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, userId string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, userId string) error
}

type Repository struct {
	Todo
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Todo: NewTodoRepository(db),
		User: NewUserRepository(db),
	}
}
