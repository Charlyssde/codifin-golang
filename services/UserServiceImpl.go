package services

import (
	"test.codifin/data"
	"test.codifin/models"
)

type UserServiceImpl struct {
	repository data.UserRepository
}

func NewUserServiceImpl(repository data.UserRepository) UserService {
	return &UserServiceImpl{repository: repository}
}

func (u UserServiceImpl) Create(user models.User) (models.User, error) {
	res, err := u.repository.Create(user)
	return res, err
}

func (u UserServiceImpl) Login(user models.User) (string, error) {
	res, err := u.repository.Login(user)
	return res, err
}
