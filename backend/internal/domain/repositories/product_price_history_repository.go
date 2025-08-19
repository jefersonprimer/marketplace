package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ProductPriceHistoryRepository interface {
	FindById(id uuid.UUID) (*entities.ProductPriceHistory, error)
	Create(productPriceHistory *entities.ProductPriceHistory) (*entities.ProductPriceHistory, error)
	GetProductPriceHistoryByProductId(productId uuid.UUID) ([]*entities.ProductPriceHistory, error)
}
