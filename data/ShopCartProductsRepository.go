package data

import (
	"test.codifin/models"
)

type ShopCartProductsRepository interface {
	Create(pId int, scId int) (models.ShopCartProducts, error)
	Delete(id int) (string, error)
}
