package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ReturnRepository interface {
	FindById(id uuid.UUID) (*entities.Return, error)
	Create(returnEntity *entities.Return) (*entities.Return, error)
	GetReturnsByOrderId(orderId uuid.UUID) ([]*entities.Return, error)
}
