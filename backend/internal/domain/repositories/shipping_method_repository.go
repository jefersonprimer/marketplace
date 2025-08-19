package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ShippingMethodRepository interface {
	FindById(id uuid.UUID) (*entities.ShippingMethod, error)
	Create(shippingMethod *entities.ShippingMethod) (*entities.ShippingMethod, error)
	GetAll() ([]*entities.ShippingMethod, error)
}
