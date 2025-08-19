package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StoreReviewUseCase struct {
	repo repositories.StoreReviewRepository
}

func NewStoreReviewUseCase(repo repositories.StoreReviewRepository) *StoreReviewUseCase {
	return &StoreReviewUseCase{repo: repo}
}

func (uc *StoreReviewUseCase) GetStoreReviewsByStoreId(storeId uuid.UUID) ([]*entities.StoreReview, error) {
	return uc.repo.GetStoreReviewsByStoreId(storeId)
}

func (uc *StoreReviewUseCase) GetStoreReviewsByUserId(userId uuid.UUID) ([]*entities.StoreReview, error) {
	return uc.repo.GetStoreReviewsByUserId(userId)
}

func (uc *StoreReviewUseCase) CreateStoreReview(review *entities.StoreReview) (*entities.StoreReview, error) {
	return uc.repo.Create(review)
}
