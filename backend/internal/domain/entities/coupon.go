
package entities

import (
    "time"

    "github.com/google/uuid"
)

type CouponDiscountType string

const (
    CouponDiscountTypePercentage CouponDiscountType = "percentage"
    CouponDiscountTypeFixed      CouponDiscountType = "fixed"
)

type CouponAppliesTo string

const (
    CouponAppliesToAll        CouponAppliesTo = "all"
    CouponAppliesToProducts   CouponAppliesTo = "products"
    CouponAppliesToCategories CouponAppliesTo = "categories"
    CouponAppliesToShipping   CouponAppliesTo = "shipping"
)

type Coupon struct {
    ID             uuid.UUID          `json:"id"`
    Code           string             `json:"code"`
    DiscountType   CouponDiscountType `json:"discount_type"`
    DiscountValue  float64            `json:"discount_value"`
    MaxUses        int                `json:"max_uses"`
    UsedCount      int                `json:"used_count"`
    ValidFrom      time.Time          `json:"valid_from"`
    ValidUntil     time.Time          `json:"valid_until"`
    StoreID        uuid.UUID          `json:"store_id"`
    MinOrderValue  float64            `json:"min_order_value"`
    AppliesTo      CouponAppliesTo    `json:"applies_to"`
    CreatedAt      time.Time          `json:"created_at"`
}
