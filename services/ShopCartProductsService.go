package services

import "test.codifin/models"

type ShopCartProductsService interface {
	Create(scp models.ShopCartProducts) (models.ShopCartProducts, error)
	Delete(id int) (string, error)
}
