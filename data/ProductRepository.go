package data

import (
	"test.codifin/models"
)

type ProductRepository interface {
	GetById(id int) (models.Product, error)
	GetAll(filter string, order string) ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id int) (string, error)
}
