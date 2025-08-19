
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ProductImage struct {
    ID        uuid.UUID `json:"id"`
    ProductID uuid.UUID `json:"product_id"`
    ImageURL  string    `json:"image_url"`
    IsMain    bool      `json:"is_main"`
    CreatedAt time.Time `json:"created_at"`
}
