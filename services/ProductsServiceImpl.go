package services

import (
	"test.codifin/data"
	"test.codifin/models"
)

type ProductsServiceImpl struct {
	repository data.ProductRepository
}

func NewProductServiceImpl(repository data.ProductRepository) ProductsService {
	return &ProductsServiceImpl{repository: repository}
}

func (p ProductsServiceImpl) Create(product models.Product) (models.Product, error) {
	res, err := p.repository.Create(product)
	if err != nil {
		return models.Product{}, err
	} else {
		return res, nil
	}
}

func (p ProductsServiceImpl) Update(product models.Product) (models.Product, error) {
	res, err := p.repository.Update(product)
	if err != nil {
		return models.Product{}, err
	} else {
		return res, nil
	}
}

func (p ProductsServiceImpl) Delete(id int) (string, error) {
	res, err := p.repository.Delete(id)
	if err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (p ProductsServiceImpl) GetById(id int) (models.Product, error) {
	res, err := p.repository.GetById(id)
	if err != nil {
		return models.Product{}, err
	} else {
		return res, nil
	}
}

func (p ProductsServiceImpl) GetAll(filter string, order string) ([]models.Product, error) {
	res, err := p.repository.GetAll(filter, order)
	if err != nil {
		return []models.Product{}, err
	} else {
		return res, nil
	}
}
