package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductAttributeRepository interface {
	Create(productAttribute *entities.ProductAttribute) error
	GetByID(id uuid.UUID) (*entities.ProductAttribute, error)
	GetByProductID(productID uuid.UUID) ([]*entities.ProductAttribute, error)
	Update(productAttribute *entities.ProductAttribute) error
	Delete(id uuid.UUID) error
}
