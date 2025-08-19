
package entities

import (
    "github.com/google/uuid"
)

type StoreCategory struct {
    StoreID    uuid.UUID `json:"store_id"`
    CategoryID uuid.UUID `json:"category_id"`
}
