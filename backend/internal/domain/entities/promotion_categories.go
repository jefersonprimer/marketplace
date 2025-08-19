package entities

import "github.com/google/uuid"

type PromotionCategory struct {
	PromotionID uuid.UUID `json:"promotion_id"`
	CategoryID  uuid.UUID `json:"category_id"`
}
