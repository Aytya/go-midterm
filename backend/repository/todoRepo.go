package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"myproject/backend/domain"
	"time"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db}
}

const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func (repo *TodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	query := `INSERT INTO todos (id, title, datetime, active_at, status, priority) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	var id string
	err := repo.db.QueryRowContext(ctx, query, todo.ID, todo.Title, todo.DateTime.Format(RFC3339), todo.ActiveAt, todo.Status, todo.Priority).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}
	todo.ID = id
	return nil
}

func (repo *TodoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	query := `SELECT * FROM todos`
	err := repo.db.SelectContext(ctx, &todos, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}
	return todos, nil
}

func (repo TodoRepository) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	var todo domain.Todo
	query := `SELECT * FROM todos WHERE id = $1`
	err := repo.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo by id: %w", err)
	}
	return &todo, nil
}

func (repo *TodoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	query := `UPDATE todos SET title = $2, datetime = $3, priority = $4 WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.DateTime.Format(time.RFC3339), todo.Priority)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	return nil
}

func (repo TodoRepository) CheckTodo(ctx context.Context, id string) error {
	query := `
		UPDATE todos
		SET status = NOT status
		WHERE id = $1
	`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to check todo: %w", err)
	}
	return nil
}

func (repo TodoRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}
