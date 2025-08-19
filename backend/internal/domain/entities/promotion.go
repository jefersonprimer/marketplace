
package entities

import (
    "time"

    "github.com/google/uuid"
)

type PromotionDiscountType string

const (
    PromotionDiscountTypePercentage PromotionDiscountType = "percentage"
    PromotionDiscountTypeFixed      PromotionDiscountType = "fixed"
)

type PromotionAppliesTo string

const (
    PromotionAppliesToAll        PromotionAppliesTo = "all"
    PromotionAppliesToProducts   PromotionAppliesTo = "products"
    PromotionAppliesToCategories PromotionAppliesTo = "categories"
)

type Promotion struct {
    ID            uuid.UUID             `json:"id"`
    StoreID       uuid.UUID             `json:"store_id"`
    Name          string                `json:"name"`
    Description   string                `json:"description"`
    DiscountType  PromotionDiscountType `json:"discount_type"`
    DiscountValue float64               `json:"discount_value"`
    StartDate     time.Time             `json:"start_date"`
    EndDate       time.Time             `json:"end_date"`
    AppliesTo     PromotionAppliesTo    `json:"applies_to"`
    Active        bool                  `json:"active"`
    CreatedAt     time.Time             `json:"created_at"`
}
