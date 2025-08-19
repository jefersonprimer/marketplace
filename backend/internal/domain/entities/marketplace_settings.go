
package entities

import (
    "time"

    "github.com/google/uuid"
)

type MarketplaceSetting struct {
    ID        uuid.UUID `json:"id"`
    Key       string    `json:"key"`
    Value     string    `json:"value"`
    UpdatedAt time.Time `json:"updated_at"`
}
