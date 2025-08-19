
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type DisputeRepository interface {
    Create(dispute *entities.Dispute) error
    GetByID(id uuid.UUID) (*entities.Dispute, error)
    GetByOrderID(orderID uuid.UUID) ([]*entities.Dispute, error)
    Update(dispute *entities.Dispute) error
    Delete(id uuid.UUID) error
}
