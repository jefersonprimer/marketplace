package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StoreCommissionRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StoreCommissionRepositoryImpl) FindById(id uuid.UUID) (*entities.StoreCommission, error) {
	panic("implement me")
}

func (r *StoreCommissionRepositoryImpl) Create(storeCommission *entities.StoreCommission) (*entities.StoreCommission, error) {
	panic("implement me")
}

func (r *StoreCommissionRepositoryImpl) GetStoreCommissionsByStoreId(storeId uuid.UUID) ([]*entities.StoreCommission, error) {
	panic("implement me")
}
