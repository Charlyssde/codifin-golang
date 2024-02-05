package data

import (
	"gorm.io/gorm"
	"test.codifin/models"
)

type ShopCartRepositoryImpl struct {
	Db *gorm.DB
}

func NewShopCartRepositoryImpl(Db *gorm.DB) ShopCartRepository {
	return ShopCartRepositoryImpl{Db: Db}
}

func (s ShopCartRepositoryImpl) Create(shopCart models.ShopCart) (models.ShopCart, error) {
	result := s.Db.Create(&shopCart)
	if result.Error != nil {
		return models.ShopCart{}, result.Error
	} else {
		return shopCart, nil
	}
}

func (s ShopCartRepositoryImpl) GetByUser(id int) (models.ShopCart, error) {
	var sc models.ShopCart
	result := s.Db.Where("user_id = ?", id).Find(&sc)
	if result.Error != nil {
		return models.ShopCart{}, result.Error
	} else {
		return sc, nil
	}
}
