package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderMessageRepository interface {
	Create(orderMessage *entities.OrderMessage) error
	GetByID(id uuid.UUID) (*entities.OrderMessage, error)
	GetByOrderID(orderID uuid.UUID) ([]*entities.OrderMessage, error)
	Update(orderMessage *entities.OrderMessage) error
	Delete(id uuid.UUID) error
}
