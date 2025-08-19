package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StoreCommissionRepository interface {
	FindById(id uuid.UUID) (*entities.StoreCommission, error)
	Create(storeCommission *entities.StoreCommission) (*entities.StoreCommission, error)
	GetStoreCommissionsByStoreId(storeId uuid.UUID) ([]*entities.StoreCommission, error)
}
