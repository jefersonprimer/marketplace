package entities

import "github.com/google/uuid"

type PromotionProduct struct {
	PromotionID uuid.UUID `json:"promotion_id"`
	ProductID   uuid.UUID `json:"product_id"`
}
