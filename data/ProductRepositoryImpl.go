package data

import (
	"errors"
	"gorm.io/gorm"
	"test.codifin/models"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p ProductRepositoryImpl) GetById(id int) (models.Product, error) {
	var product models.Product
	result := p.Db.Find(&product, id)
	if result != nil {
		return product, nil
	} else {
		return models.Product{}, errors.New("no existe un producto con ese ID")
	}
}

func (p ProductRepositoryImpl) GetAll(filter string, order string) ([]models.Product, error) {
	var product []models.Product
	result := p.Db.Where("name=?", filter).Find(&product).Order("name " + order)

	if result.Error != nil {
		return []models.Product{}, result.Error
	}
	return product, nil
}

func (p ProductRepositoryImpl) Create(product models.Product) (models.Product, error) {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return product, result.Error
	} else {
		return product, nil
	}
}

func (p ProductRepositoryImpl) Update(product models.Product) (models.Product, error) {
	result := p.Db.Updates(&product)
	if result.Error != nil {
		return product, result.Error
	} else {
		return product, nil
	}
}

func (p ProductRepositoryImpl) Delete(id int) (string, error) {
	var product models.Product
	result := p.Db.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return "", result.Error
	} else {
		return "eliminado correctamente", nil
	}
}
