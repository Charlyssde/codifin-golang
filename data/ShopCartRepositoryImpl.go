package data

import (
	"gorm.io/gorm"
	"test.codifin/models"
	"time"
)

type ShopCartRepositoryImpl struct {
	Db *gorm.DB
}

func NewShopCartRepositoryImpl(Db *gorm.DB) ShopCartRepository {
	return ShopCartRepositoryImpl{Db: Db}
}

func (s ShopCartRepositoryImpl) AddProduct(id int, productId int) (models.ShopCart, error) {
	var shopCart models.ShopCart
	var scp *models.ShopCartProducts
	var product models.Product

	s.Db.Where("id = ?", id).Find(&shopCart)
	s.Db.Where("id = ?", productId).Find(&product)
	s.Db.Where("shop_cart_id = ?", id).Where("product_id = ?", productId).Find(&scp)

	if scp != nil {
		s.Db.Create(&scp)
		shopCart.Subtotal += product.Price
		shopCart.Total += product.Price
		s.Db.Updates(&shopCart)
	} else {
		scp.Quantity += 1
		scp.Amount += product.Price
		shopCart.Subtotal += product.Price
		shopCart.Total += product.Price
		s.Db.Updates(&shopCart)
		s.Db.Updates(&scp)
	}
	return shopCart, nil
}

func (s ShopCartRepositoryImpl) RemoveProduct(id int, productId int) (models.ShopCart, error) {
	var shopCart models.ShopCart
	var scp *models.ShopCartProducts
	var product models.Product
	var products []models.ShopCartProducts

	s.Db.Where("id = ?", id).Find(&shopCart)
	s.Db.Where("id = ?", productId).Find(&product)
	s.Db.Where("shop_cart_id = ?", id).Where("product_id = ?", productId).Find(&scp)

	if scp != nil {
		s.Db.Delete(&scp)
		shopCart.Subtotal -= product.Price
		shopCart.Total -= product.Price
		s.Db.Updates(&shopCart)
		s.Db.Where("shop_cart_id = ? and quantity > 0", id).Find(&products)
		shopCart.Products = products
	} else {
		scp.Quantity -= 1
		scp.Amount -= product.Price
		shopCart.Subtotal -= product.Price
		shopCart.Total -= product.Price
		s.Db.Updates(&shopCart)
		s.Db.Updates(&scp)
		s.Db.Where("shop_cart_id = ? and quantity > 0", id).Find(&products)
		shopCart.Products = products
	}
	return shopCart, nil
}

func (s ShopCartRepositoryImpl) Create(shopCart models.ShopCart) (models.ShopCart, error) {
	shopCart.CreatedAt = time.Now().String()
	result := s.Db.Create(&shopCart)
	if result.Error != nil {
		return models.ShopCart{}, result.Error
	} else {
		return shopCart, nil
	}
}

func (s ShopCartRepositoryImpl) GetByUser(id int) (models.ShopCart, error) {
	var sc models.ShopCart
	result := s.Db.Where("user_id = ?", id).Find(&sc)
	if result.Error != nil {
		return models.ShopCart{}, result.Error
	} else {
		var products []models.ShopCartProducts
		s.Db.Where("shop_cart_id = ?", sc.ID).Find(&products)
		sc.Products = products
		return sc, nil
	}
}
