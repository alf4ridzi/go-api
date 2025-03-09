package repositories

import (
	"api/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepositories struct {
	DB *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *UserRepositories {
	return &UserRepositories{DB: db}
}

func (r *UserRepositories) CreateUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(&user).Error
}

func (r *UserRepositories) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepositories) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepositories) IsEmailExist(ctx context.Context, email string) (bool, error) {
	user, err := r.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (r *UserRepositories) IsUsernameExist(ctx context.Context, username string) (bool, error) {
	user, err := r.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}
