
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StockMovement struct {
    ID        uuid.UUID `json:"id"`
    ProductID uuid.UUID `json:"product_id"`
    Change    int       `json:"change"`
    Reason    string    `json:"reason"`
    CreatedAt time.Time `json:"created_at"`
}
