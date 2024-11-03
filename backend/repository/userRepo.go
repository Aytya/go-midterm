package repository

import (
	"context"
	"gorm.io/gorm"
	"myproject/backend/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) Create(ctx context.Context, user *domain.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByID(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}
	if err := u.db.Where("id = ?", userId).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Update(ctx context.Context, user *domain.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(ctx context.Context, userId string) error {
	if err := u.db.Delete(&domain.User{}, userId).Error; err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
