package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StorePolicyUseCase struct {
	repo repositories.StorePolicyRepository
}

func NewStorePolicyUseCase(repo repositories.StorePolicyRepository) *StorePolicyUseCase {
	return &StorePolicyUseCase{repo: repo}
}

func (uc *StorePolicyUseCase) GetStorePoliciesByStoreId(storeId uuid.UUID) ([]*entities.StorePolicy, error) {
	return uc.repo.GetStorePoliciesByStoreId(storeId)
}

func (uc *StorePolicyUseCase) CreateStorePolicy(policy *entities.StorePolicy) (*entities.StorePolicy, error) {
	return uc.repo.Create(policy)
}
