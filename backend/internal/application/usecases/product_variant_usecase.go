package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ProductVariantUseCase struct {
	repo repositories.ProductVariantRepository
}

func NewProductVariantUseCase(repo repositories.ProductVariantRepository) *ProductVariantUseCase {
	return &ProductVariantUseCase{repo: repo}
}

func (uc *ProductVariantUseCase) GetProductVariantsByProductId(productId uuid.UUID) ([]*entities.ProductVariant, error) {
	return uc.repo.GetProductVariantsByProductId(productId)
}

func (uc *ProductVariantUseCase) CreateProductVariant(variant *entities.ProductVariant) (*entities.ProductVariant, error) {
	return uc.repo.Create(variant)
}
