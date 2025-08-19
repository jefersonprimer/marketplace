
package entities

import (
    "github.com/google/uuid"
)

type ProductVariant struct {
    ID        uuid.UUID `json:"id"`
    ProductID uuid.UUID `json:"product_id"`
    Name      string    `json:"name"`
}
