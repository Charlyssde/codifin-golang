package data

import (
	"test.codifin/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	Login(user models.User) (int, error)
}
