package data

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"test.codifin/models"
	"time"
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

func (u UserRepositoryImpl) Login(user models.User) (string, error) {
	var loged models.User
	result := u.Db.Where("name = ?", user.Name).Where("password = ?", user.Password).Find(&loged)
	if result.Error != nil {
		return "no existe el usuario", result.Error
	} else {
		token := jwt.New(jwt.SigningMethodEdDSA)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(60 * time.Minute)
		claims["authorized"] = true
		claims["user"] = user.Name
		claims["id"] = user.ID
		return token.Raw, nil
	}
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return UserRepositoryImpl{Db}
}
