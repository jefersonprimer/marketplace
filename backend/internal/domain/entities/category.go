
package entities

import (
    "time"

    "github.com/google/uuid"
)

type Category struct {
    ID        uuid.UUID `json:"id"`
    Name      string    `json:"name"`
    Slug      string    `json:"slug"`
    ParentID  uuid.UUID `json:"parent_id"`
    CreatedAt time.Time `json:"created_at"`
}
