package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StoreTransactionUseCase struct {
	repo repositories.StoreTransactionRepository
}

func NewStoreTransactionUseCase(repo repositories.StoreTransactionRepository) *StoreTransactionUseCase {
	return &StoreTransactionUseCase{repo: repo}
}

func (uc *StoreTransactionUseCase) GetStoreTransactionsByStoreId(storeId uuid.UUID) ([]*entities.StoreTransaction, error) {
	return uc.repo.GetStoreTransactionsByStoreId(storeId)
}

func (uc *StoreTransactionUseCase) CreateStoreTransaction(transaction *entities.StoreTransaction) (*entities.StoreTransaction, error) {
	return uc.repo.Create(transaction)
}
