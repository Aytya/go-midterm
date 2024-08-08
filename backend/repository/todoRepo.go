package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"myproject/backend/domain"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (repo *TodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	query := `INSERT INTO todos (id, title, date, time, active_at, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	log.Printf("Executing query: %s with values id: %s, title: %s, date: %s, time: %s, active_at: %s, status: %v", query, todo.ID, todo.Title, todo.Date, todo.Time, todo.ActiveAt, todo.Status)
	var id string
	err := repo.db.QueryRowContext(ctx, query, todo.ID, todo.Title, todo.Date, todo.Time, todo.ActiveAt, todo.Status).Scan(&id)
	if err != nil {
		log.Printf("Error executing query: %s, error: %v", query, err)
		return fmt.Errorf("failed to create todo: %w", err)
	}
	todo.ID = id
	log.Printf("Created todo with ID: %s", id)
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
	query := `UPDATE todos SET title = $2, date = $3, time = $4 WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Date, todo.Time)
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

func (repo TodoRepository) CheckedTodo(id string) error {
	//TODO implement me
	panic("implement me")
}
