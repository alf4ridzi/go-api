package services

import (
	"api/crypto"
	"api/models"
	"api/repositories"
	"errors"
)

type UserService struct {
	repo *repositories.UserRepositories
}

func NewUserService(repo *repositories.UserRepositories) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	// check username
	isUsername, err := s.repo.IsUsernameExist(user.Username)
	if err != nil {
		return err
	}

	if isUsername {
		return errors.New("username already exist")
	}

	// check email
	isEmail, err := s.repo.IsEmailExist(*user.Email)
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
	return s.repo.CreateUser(user)
}

func (s UserService) VerifyLogin(user *models.Login) error {
	userDetail, err := s.repo.GetUserByEmail(user.Email)
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
