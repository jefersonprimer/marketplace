
package entities

import (
    "time"

    "github.com/google/uuid"
)

type Warehouse struct {
    ID        uuid.UUID `json:"id"`
    StoreID   uuid.UUID `json:"store_id"`
    Name      string    `json:"name"`
    AddressID uuid.UUID `json:"address_id"`
    CreatedAt time.Time `json:"created_at"`
}
