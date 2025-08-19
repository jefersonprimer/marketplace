package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StockMovementRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StockMovementRepositoryImpl) FindById(id uuid.UUID) (*entities.StockMovement, error) {
	panic("implement me")
}

func (r *StockMovementRepositoryImpl) Create(stockMovement *entities.StockMovement) (*entities.StockMovement, error) {
	panic("implement me")
}

func (r *StockMovementRepositoryImpl) GetStockMovementsByProductId(productId uuid.UUID) ([]*entities.StockMovement, error) {
	panic("implement me")
}
