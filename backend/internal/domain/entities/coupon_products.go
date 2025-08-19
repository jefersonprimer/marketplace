package entities

import "github.com/google/uuid"

type CouponProduct struct {
	CouponID  uuid.UUID `json:"coupon_id"`
	ProductID uuid.UUID `json:"product_id"`
}
