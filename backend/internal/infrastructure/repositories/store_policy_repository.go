package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StorePolicyRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StorePolicyRepositoryImpl) FindById(id uuid.UUID) (*entities.StorePolicy, error) {
	panic("implement me")
}

func (r *StorePolicyRepositoryImpl) Create(storePolicy *entities.StorePolicy) (*entities.StorePolicy, error) {
	panic("implement me")
}

func (r *StorePolicyRepositoryImpl) GetStorePoliciesByStoreId(storeId uuid.UUID) ([]*entities.StorePolicy, error) {
	panic("implement me")
}
