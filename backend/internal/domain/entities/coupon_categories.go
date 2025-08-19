package entities

import "github.com/google/uuid"

type CouponCategory struct {
	CouponID   uuid.UUID `json:"coupon_id"`
	CategoryID uuid.UUID `json:"category_id"`
}
