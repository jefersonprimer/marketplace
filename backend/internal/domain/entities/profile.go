
package entities

import (
    "time"

    "github.com/google/uuid"
)

type Profile struct {
    ID        uuid.UUID `json:"id"`
    FullName  string    `json:"full_name"`
    Role      UserRole  `json:"role"`
    AvatarURL string    `json:"avatar_url"`
    CPF       string    `json:"cpf"`
    CNPJ      string    `json:"cnpj"`
    Phone     string    `json:"phone"`
    IsSeller  bool      `json:"is_seller"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
