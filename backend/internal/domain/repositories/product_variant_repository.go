package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ProductVariantRepository interface {
	FindById(id uuid.UUID) (*entities.ProductVariant, error)
	Create(productVariant *entities.ProductVariant) (*entities.ProductVariant, error)
	GetProductVariantsByProductId(productId uuid.UUID) ([]*entities.ProductVariant, error)
}
