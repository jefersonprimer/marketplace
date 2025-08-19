package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type WarehouseRepositoryImpl struct {
	DB *gorm.DB
}

func (r *WarehouseRepositoryImpl) FindById(id uuid.UUID) (*entities.Warehouse, error) {
	panic("implement me")
}

func (r *WarehouseRepositoryImpl) Create(warehouse *entities.Warehouse) (*entities.Warehouse, error) {
	panic("implement me")
}

func (r *WarehouseRepositoryImpl) GetWarehousesByStoreId(storeId uuid.UUID) ([]*entities.Warehouse, error) {
	panic("implement me")
}
