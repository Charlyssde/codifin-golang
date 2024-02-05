package data

import (
	"test.codifin/models"
)

type ShopCartRepository interface {
	Create(shopCart models.ShopCart) (models.ShopCart, error)
	GetByUser(id int) (models.ShopCart, error)
}
