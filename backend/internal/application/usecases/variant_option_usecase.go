package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type VariantOptionUseCase struct {
	repo repositories.VariantOptionRepository
}

func NewVariantOptionUseCase(repo repositories.VariantOptionRepository) *VariantOptionUseCase {
	return &VariantOptionUseCase{repo: repo}
}

func (uc *VariantOptionUseCase) GetVariantOptionsByVariantId(variantId uuid.UUID) ([]*entities.VariantOption, error) {
	return uc.repo.GetVariantOptionsByVariantId(variantId)
}

func (uc *VariantOptionUseCase) CreateVariantOption(option *entities.VariantOption) (*entities.VariantOption, error) {
	return uc.repo.Create(option)
}
