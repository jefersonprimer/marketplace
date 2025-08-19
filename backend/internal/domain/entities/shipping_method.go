
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ShippingMethod struct {
    ID            uuid.UUID `json:"id"`
    Name          string    `json:"name"`
    Description   string    `json:"description"`
    Cost          float64   `json:"cost"`
    EstimatedDays int       `json:"estimated_days"`
    CreatedAt     time.Time `json:"created_at"`
}
