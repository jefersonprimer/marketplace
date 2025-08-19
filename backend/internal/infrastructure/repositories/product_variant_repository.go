package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ProductVariantRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ProductVariantRepositoryImpl) FindById(id uuid.UUID) (*entities.ProductVariant, error) {
	panic("implement me")
}

func (r *ProductVariantRepositoryImpl) Create(productVariant *entities.ProductVariant) (*entities.ProductVariant, error) {
	panic("implement me")
}

func (r *ProductVariantRepositoryImpl) GetProductVariantsByProductId(productId uuid.UUID) ([]*entities.ProductVariant, error) {
	panic("implement me")
}
