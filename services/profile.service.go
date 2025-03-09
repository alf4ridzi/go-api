package services

import (
	"api/models"
	"api/repositories"
	"context"
)

type ProfileService struct {
	repo *repositories.UserRepositories
}

func NewProfileService(repo *repositories.UserRepositories) *ProfileService {
	return &ProfileService{repo: repo}
}

func (r *ProfileService) GetUserProfile(ctx context.Context, username string) (*models.User, error) {
	var profile models.User
	err := r.repo.DB.WithContext(ctx).Where("username = ?", username).First(&profile).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
