
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ProductAttribute struct {
    ID        uuid.UUID `json:"id"`
    ProductID uuid.UUID `json:"product_id"`
    Name      string    `json:"name"`
    Value     string    `json:"value"`
    CreatedAt time.Time `json:"created_at"`
}
