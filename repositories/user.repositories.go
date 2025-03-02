package repositories

import (
	"api/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepositories struct {
	DB *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *UserRepositories {
	return &UserRepositories{DB: db}
}

func (r *UserRepositories) CreateUser(user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepositories) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepositories) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepositories) IsEmailExist(email string) (bool, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (r *UserRepositories) IsUsernameExist(username string) (bool, error) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}
