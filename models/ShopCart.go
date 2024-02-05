package models

import "gofr.dev/pkg/gofr/types"

type ShopCart struct {
	ID        int        `json: id`
	UserId    int        `json: user_id`
	CreatedAt types.Date `json: created_at`
	Subtotal  float64    `json: subtotal`
	Total     float64    `json: total`
}
