package models

type ShopCartProducts struct {
	ID         int     `json: id`
	ShopCartId int     `json: shop_cart_id`
	ProductId  int     `json: product_id`
	Quantity   int     `json: quantity`
	Amount     float64 `json: amount`
}
