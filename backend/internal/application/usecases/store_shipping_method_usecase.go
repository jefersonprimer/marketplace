package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StoreShippingMethodUseCase struct {
	repo repositories.StoreShippingMethodRepository
}

func NewStoreShippingMethodUseCase(repo repositories.StoreShippingMethodRepository) *StoreShippingMethodUseCase {
	return &StoreShippingMethodUseCase{repo: repo}
}

func (uc *StoreShippingMethodUseCase) GetStoreShippingMethodsByStoreId(storeId uuid.UUID) ([]*entities.StoreShippingMethod, error) {
	return uc.repo.GetStoreShippingMethodsByStoreId(storeId)
}

func (uc *StoreShippingMethodUseCase) CreateStoreShippingMethod(method *entities.StoreShippingMethod) (*entities.StoreShippingMethod, error) {
	return uc.repo.Create(method)
}
