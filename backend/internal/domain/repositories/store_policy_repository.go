package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StorePolicyRepository interface {
	FindById(id uuid.UUID) (*entities.StorePolicy, error)
	Create(storePolicy *entities.StorePolicy) (*entities.StorePolicy, error)
	GetStorePoliciesByStoreId(storeId uuid.UUID) ([]*entities.StorePolicy, error)
}
