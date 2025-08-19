package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type WarehouseUseCase struct {
	repo repositories.WarehouseRepository
}

func NewWarehouseUseCase(repo repositories.WarehouseRepository) *WarehouseUseCase {
	return &WarehouseUseCase{repo: repo}
}

func (uc *WarehouseUseCase) GetWarehousesByStoreId(storeId uuid.UUID) ([]*entities.Warehouse, error) {
	return uc.repo.GetWarehousesByStoreId(storeId)
}

func (uc *WarehouseUseCase) CreateWarehouse(warehouse *entities.Warehouse) (*entities.Warehouse, error) {
	return uc.repo.Create(warehouse)
}
