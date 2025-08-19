package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StoreTransactionRepository interface {
	FindById(id uuid.UUID) (*entities.StoreTransaction, error)
	Create(storeTransaction *entities.StoreTransaction) (*entities.StoreTransaction, error)
	GetStoreTransactionsByStoreId(storeId uuid.UUID) ([]*entities.StoreTransaction, error)
}
