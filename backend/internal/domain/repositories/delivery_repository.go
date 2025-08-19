
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type DeliveryRepository interface {
    Create(delivery *entities.Delivery) error
    GetByID(id uuid.UUID) (*entities.Delivery, error)
    GetByOrderID(orderID uuid.UUID) (*entities.Delivery, error)
    Update(delivery *entities.Delivery) error
    Delete(id uuid.UUID) error
}
