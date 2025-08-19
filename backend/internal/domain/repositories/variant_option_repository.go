package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type VariantOptionRepository interface {
	FindById(id uuid.UUID) (*entities.VariantOption, error)
	Create(variantOption *entities.VariantOption) (*entities.VariantOption, error)
	GetVariantOptionsByVariantId(variantId uuid.UUID) ([]*entities.VariantOption, error)
}
