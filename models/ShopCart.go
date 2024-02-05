package models

type ShopCart struct {
	ID        int     `json: id`
	UserId    int     `json: user_id`
	CreatedAt string  `json: created_at`
	Subtotal  float64 `json: subtotal`
	Total     float64 `json: total`
	Products  []ShopCartProducts
}
