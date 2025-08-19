package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type PromotionUseCase struct {
	repo repositories.PromotionRepository
}

func NewPromotionUseCase(repo repositories.PromotionRepository) *PromotionUseCase {
	return &PromotionUseCase{repo: repo}
}

func (uc *PromotionUseCase) GetPromotionsByStoreId(storeId uuid.UUID) ([]*entities.Promotion, error) {
	return uc.repo.GetPromotionsByStoreId(storeId)
}

func (uc *PromotionUseCase) CreatePromotion(promotion *entities.Promotion) (*entities.Promotion, error) {
	return uc.repo.Create(promotion)
}
