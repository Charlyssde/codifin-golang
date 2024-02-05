package services

import "test.codifin/models"

type UserService interface {
	Create(user models.User) (models.User, error)
	Login(user models.User) (string, error)
}
