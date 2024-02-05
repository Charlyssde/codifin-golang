package services

import "test.codifin/models"

type ProductsService interface {
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id int) (string, error)
	GetById(id int) (models.Product, error)
	GetAll(filter string, order string) ([]models.Product, error)
}
