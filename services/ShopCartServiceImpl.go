package services

import (
	"test.codifin/data"
	"test.codifin/models"
)

type ShopCartServiceImpl struct {
	repository data.ShopCartRepository
}

func NewShopCartServiceImpl(repository data.ShopCartRepository) ShopCartService {
	return &ShopCartServiceImpl{
		repository: repository,
	}
}

func (s ShopCartServiceImpl) Create(cart models.ShopCart) (models.ShopCart, error) {
	res, err := s.repository.Create(cart)
	return res, err
}

func (s ShopCartServiceImpl) GetByUser(id int) (models.ShopCart, error) {
	res, err := s.repository.GetByUser(id)
	return res, err
}
