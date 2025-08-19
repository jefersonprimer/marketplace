
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StoreShippingMethod struct {
    ID               uuid.UUID `json:"id"`
    StoreID          uuid.UUID `json:"store_id"`
    ShippingMethodID uuid.UUID `json:"shipping_method_id"`
    Active           bool      `json:"active"`
    CreatedAt        time.Time `json:"created_at"`
}
