
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type AddressRepository interface {
    Create(address *entities.Address) error
    GetByID(id uuid.UUID) (*entities.Address, error)
    GetByUserID(userID uuid.UUID) ([]*entities.Address, error)
    Update(address *entities.Address) error
    Delete(id uuid.UUID) error
}
