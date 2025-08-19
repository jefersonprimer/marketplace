package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StoreTransactionRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StoreTransactionRepositoryImpl) FindById(id uuid.UUID) (*entities.StoreTransaction, error) {
	panic("implement me")
}

func (r *StoreTransactionRepositoryImpl) Create(storeTransaction *entities.StoreTransaction) (*entities.StoreTransaction, error) {
	panic("implement me")
}

func (r *StoreTransactionRepositoryImpl) GetStoreTransactionsByStoreId(storeId uuid.UUID) ([]*entities.StoreTransaction, error) {
	panic("implement me")
}
