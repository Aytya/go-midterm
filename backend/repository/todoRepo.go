package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"myproject/backend/domain"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (repo *TodoRepository) Create(ctx context.Context, todo *domain.Todo) (err error) {
	if err = repo.db.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepository) GetAll(ctx context.Context) (todos []*domain.Todo, err error) {
	if err = repo.db.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (repo *TodoRepository) GetByID(ctx context.Context, id string) (todo *domain.Todo, err error) {
	if err = repo.db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (repo *TodoRepository) Update(ctx context.Context, todo *domain.Todo) (err error) {
	if err = repo.db.Save(todo).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepository) CheckTodo(ctx context.Context, id string) (err error) {
	var todo domain.Todo

	if err := repo.db.First(&todo, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to find todo: %w", err)
	}

	newStatus := !todo.Status

	if err = repo.db.Model(&todo).Update("status", newStatus).Error; err != nil {
		return fmt.Errorf("failed to toggle complete status: %w", err)
	}

	return nil
}

func (repo *TodoRepository) Delete(ctx context.Context, id string) (err error) {
	if err = repo.db.Where("id = ?", id).Delete(&domain.Todo{}).Error; err != nil {
		return err
	}

	return nil
}
