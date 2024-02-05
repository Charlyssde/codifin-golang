package data

import (
	"gorm.io/gorm"
	"test.codifin/models"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) Create(user models.User) (models.User, error) {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	} else {
		return user, nil
	}
}

func (u UserRepositoryImpl) Login(user models.User) (int, error) {
	var logged models.User
	result := u.Db.Where("name = ?", user.Name).Where("password = ?", user.Password).Find(&logged)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return logged.ID, nil
	}
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return UserRepositoryImpl{Db}
}
