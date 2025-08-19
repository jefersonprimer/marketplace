package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ProductPriceHistoryUseCase struct {
	repo repositories.ProductPriceHistoryRepository
}

func NewProductPriceHistoryUseCase(repo repositories.ProductPriceHistoryRepository) *ProductPriceHistoryUseCase {
	return &ProductPriceHistoryUseCase{repo: repo}
}

func (uc *ProductPriceHistoryUseCase) GetProductPriceHistoryByProductId(productId uuid.UUID) ([]*entities.ProductPriceHistory, error) {
	return uc.repo.GetProductPriceHistoryByProductId(productId)
}
