package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StockMovementRepository interface {
	FindById(id uuid.UUID) (*entities.StockMovement, error)
	Create(stockMovement *entities.StockMovement) (*entities.StockMovement, error)
	GetStockMovementsByProductId(productId uuid.UUID) ([]*entities.StockMovement, error)
}
