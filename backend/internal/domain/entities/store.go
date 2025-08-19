
package entities

import (
    "time"

    "github.com/google/uuid"
)

type Store struct {
    ID          uuid.UUID `json:"id"`
    OwnerID     uuid.UUID `json:"owner_id"`
    Name        string    `json:"name"`
    Slug        string    `json:"slug"`
    Description string    `json:"description"`
    LogoURL     string    `json:"logo_url"`
    BannerURL   string    `json:"banner_url"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
