package services

import "test.codifin/models"

type ShopCartService interface {
	Create(cart models.ShopCart) (models.ShopCart, error)
	GetByUser(id int) (models.ShopCart, error)
}
