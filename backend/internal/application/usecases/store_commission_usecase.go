package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StoreCommissionUseCase struct {
	repo repositories.StoreCommissionRepository
}

func NewStoreCommissionUseCase(repo repositories.StoreCommissionRepository) *StoreCommissionUseCase {
	return &StoreCommissionUseCase{repo: repo}
}

func (uc *StoreCommissionUseCase) GetStoreCommissionsByStoreId(storeId uuid.UUID) ([]*entities.StoreCommission, error) {
	return uc.repo.GetStoreCommissionsByStoreId(storeId)
}

func (uc *StoreCommissionUseCase) CreateStoreCommission(commission *entities.StoreCommission) (*entities.StoreCommission, error) {
	return uc.repo.Create(commission)
}
