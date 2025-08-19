
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StoreReview struct {
    ID        uuid.UUID `json:"id"`
    StoreID   uuid.UUID `json:"store_id"`
    UserID    uuid.UUID `json:"user_id"`
    Rating    int       `json:"rating"`
    Comment   string    `json:"comment"`
    CreatedAt time.Time `json:"created_at"`
}
