
package entities

import (
    "time"

    "github.com/google/uuid"
)

type UserActivityLog struct {
    ID        uuid.UUID `json:"id"`
    UserID    uuid.UUID `json:"user_id"`
    Action    string    `json:"action"`
    Metadata  string    `json:"metadata"`
    CreatedAt time.Time `json:"created_at"`
}
