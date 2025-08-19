package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderGroupRepository interface {
	Create(orderGroup *entities.OrderGroup) error
	GetByID(id uuid.UUID) (*entities.OrderGroup, error)
	Update(orderGroup *entities.OrderGroup) error
	Delete(id uuid.UUID) error
}
