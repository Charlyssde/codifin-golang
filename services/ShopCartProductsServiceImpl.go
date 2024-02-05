package services

import (
	"test.codifin/data"
	"test.codifin/models"
)

type ShopCartProductsServiceImpl struct {
	repository data.ShopCartProductsRepository
}

func NewShopCartProductsServiceImpl(repository data.ShopCartProductsRepository) ShopCartProductsService {
	return &ShopCartProductsServiceImpl{repository: repository}
}

func (s ShopCartProductsServiceImpl) Create(scp models.ShopCartProducts) (models.ShopCartProducts, error) {
	res, err := s.repository.Create(scp)
	return res, err
}

func (s ShopCartProductsServiceImpl) Delete(id int) (string, error) {
	res, err := s.repository.Delete(id)
	return res, err
}
