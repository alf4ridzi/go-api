package services

import (
	"api/crypto"
	"api/models"
	"api/repositories"
	"context"
	"errors"
)

type AuthService struct {
	repo *repositories.UserRepositories
}

func NewAuthService(repo *repositories.UserRepositories) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *models.User) error {
	// check username
	isUsername, err := s.repo.IsUsernameExist(ctx, user.Username)
	if err != nil {
		return err
	}

	if isUsername {
		return errors.New("username already exist")
	}

	// check email
	isEmail, err := s.repo.IsEmailExist(ctx, *user.Email)
	if err != nil {
		return err
	}

	if isEmail {
		return errors.New("email already exist")
	}

	hashPw, err := crypto.HashPassword(user.Password)
	if err != nil {
		return errors.New(err.Error())
	}

	user.Password = hashPw
	return s.repo.CreateUser(ctx, user)
}

func (s AuthService) VerifyLogin(ctx context.Context, user *models.Login) error {
	userDetail, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if userDetail == nil {
		return errors.New("invalid email or password")
	}

	if !crypto.CheckPasswordHash(user.Password, userDetail.Password) {
		return errors.New("invalid email or password")
	}

	return nil
}
