
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type StoreRepository interface {
    Create(store *entities.Store) error
    GetByID(id uuid.UUID) (*entities.Store, error)
    GetBySlug(slug string) (*entities.Store, error)
    GetByOwnerID(ownerID uuid.UUID) ([]entities.Store, error)
    Update(store *entities.Store) error
    Delete(id uuid.UUID) error
}
