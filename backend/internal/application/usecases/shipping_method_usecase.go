package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ShippingMethodUseCase struct {
	repo repositories.ShippingMethodRepository
}

func NewShippingMethodUseCase(repo repositories.ShippingMethodRepository) *ShippingMethodUseCase {
	return &ShippingMethodUseCase{repo: repo}
}

func (uc *ShippingMethodUseCase) GetAll() ([]*entities.ShippingMethod, error) {
	return uc.repo.GetAll()
}

func (uc *ShippingMethodUseCase) CreateShippingMethod(method *entities.ShippingMethod) (*entities.ShippingMethod, error) {
	return uc.repo.Create(method)
}
