package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderStatusHistoryRepository interface {
	Create(orderStatusHistory *entities.OrderStatusHistory) error
	GetByID(id uuid.UUID) (*entities.OrderStatusHistory, error)
	GetByOrderID(orderID uuid.UUID) ([]*entities.OrderStatusHistory, error)
	Update(orderStatusHistory *entities.OrderStatusHistory) error
	Delete(id uuid.UUID) error
}
