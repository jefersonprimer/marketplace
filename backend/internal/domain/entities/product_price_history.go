
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ProductPriceHistory struct {
    ID        uuid.UUID `json:"id"`
    ProductID uuid.UUID `json:"product_id"`
    OldPrice  float64   `json:"old_price"`
    NewPrice  float64   `json:"new_price"`
    ChangedAt time.Time `json:"changed_at"`
}
