package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderItemRepository interface {
	Create(orderItem *entities.OrderItem) error
	GetByID(id uuid.UUID) (*entities.OrderItem, error)
	GetByOrderID(orderID uuid.UUID) ([]*entities.OrderItem, error)
	Update(orderItem *entities.OrderItem) error
	Delete(id uuid.UUID) error
}
