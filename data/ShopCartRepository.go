package data

import (
	"test.codifin/models"
)

type ShopCartRepository interface {
	Create(shopCart models.ShopCart) (models.ShopCart, error)
	AddProduct(id int, product int) (models.ShopCart, error)
	RemoveProduct(id int, product int) (models.ShopCart, error)
	GetByUser(id int) (models.ShopCart, error)
}
