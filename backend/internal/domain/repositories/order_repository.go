package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order *entities.Order) error
	GetByID(id uuid.UUID) (*entities.Order, error)
	Update(order *entities.Order) error
	Delete(id uuid.UUID) error
}
