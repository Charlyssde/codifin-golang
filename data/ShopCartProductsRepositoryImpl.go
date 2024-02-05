package data

import (
	"gorm.io/gorm"
	"test.codifin/models"
)

type ShopCartProductsRepositoryImpl struct {
	Db *gorm.DB
}

func NewShopCartProductsRepositoryImpl(Db *gorm.DB) ShopCartProductsRepository {
	return &ShopCartProductsRepositoryImpl{Db: Db}
}

func (s ShopCartProductsRepositoryImpl) Create(pId int, scId int) (models.ShopCartProducts, error) {
	var scp models.ShopCartProducts
	res := s.Db.
		Where("shop_cart_id = ?", scId).
		Where("product_id = ?", pId).Find(&scp)
	if res != nil {
		result := s.Db.Create(&scp)
		if result.Error != nil {
			return models.ShopCartProducts{}, result.Error
		} else {
			return scp, nil
		}
	} else {
		scp.Quantity += 1
		result := s.Db.Updates(scp)
		if result.Error != nil {
			return models.ShopCartProducts{}, result.Error
		} else {
			return scp, nil
		}
	}
}

func (s ShopCartProductsRepositoryImpl) Delete(id int) (string, error) {
	var scp models.ShopCartProducts
	result := s.Db.Where("id = ?", id).Delete(&scp)
	if result.Error != nil {
		return "", result.Error
	} else {
		return "eliminado correctamente", nil
	}
}
