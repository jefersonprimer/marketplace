
package entities

import (
    "time"

    "github.com/google/uuid"
)

type VariantOption struct {
    ID              uuid.UUID `json:"id"`
    VariantID       uuid.UUID `json:"variant_id"`
    Value           string    `json:"value"`
    PriceAdjustment float64   `json:"price_adjustment"`
    Stock           int       `json:"stock"`
    CreatedAt       time.Time `json:"created_at"`
}
