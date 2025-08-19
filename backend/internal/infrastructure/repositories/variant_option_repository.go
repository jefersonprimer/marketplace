package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type VariantOptionRepositoryImpl struct {
	DB *gorm.DB
}

func (r *VariantOptionRepositoryImpl) FindById(id uuid.UUID) (*entities.VariantOption, error) {
	panic("implement me")
}

func (r *VariantOptionRepositoryImpl) Create(variantOption *entities.VariantOption) (*entities.VariantOption, error) {
	panic("implement me")
}

func (r *VariantOptionRepositoryImpl) GetVariantOptionsByVariantId(variantId uuid.UUID) ([]*entities.VariantOption, error) {
	panic("implement me")
}
