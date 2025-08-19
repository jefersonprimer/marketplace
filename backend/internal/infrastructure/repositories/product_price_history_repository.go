package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ProductPriceHistoryRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ProductPriceHistoryRepositoryImpl) FindById(id uuid.UUID) (*entities.ProductPriceHistory, error) {
	panic("implement me")
}

func (r *ProductPriceHistoryRepositoryImpl) Create(productPriceHistory *entities.ProductPriceHistory) (*entities.ProductPriceHistory, error) {
	panic("implement me")
}

func (r *ProductPriceHistoryRepositoryImpl) GetProductPriceHistoryByProductId(productId uuid.UUID) ([]*entities.ProductPriceHistory, error) {
	panic("implement me")
}
