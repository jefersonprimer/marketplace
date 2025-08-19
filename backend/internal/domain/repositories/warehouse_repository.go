package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type WarehouseRepository interface {
	FindById(id uuid.UUID) (*entities.Warehouse, error)
	Create(warehouse *entities.Warehouse) (*entities.Warehouse, error)
	GetWarehousesByStoreId(storeId uuid.UUID) ([]*entities.Warehouse, error)
}
