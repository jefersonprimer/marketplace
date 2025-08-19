
package entities

import (
    "time"

    "github.com/google/uuid"
)

type Wishlist struct {
    ID        uuid.UUID `json:"id"`
    UserID    uuid.UUID `json:"user_id"`
    ProductID uuid.UUID `json:"product_id"`
    CreatedAt time.Time `json:"created_at"`
}
