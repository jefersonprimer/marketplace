package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductImageRepository interface {
	Create(productImage *entities.ProductImage) error
	GetByID(id uuid.UUID) (*entities.ProductImage, error)
	GetByProductID(productID uuid.UUID) ([]*entities.ProductImage, error)
	Update(productImage *entities.ProductImage) error
	Delete(id uuid.UUID) error
}
