package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ProductReviewUseCase struct {
	repo repositories.ProductReviewRepository
}

func NewProductReviewUseCase(repo repositories.ProductReviewRepository) *ProductReviewUseCase {
	return &ProductReviewUseCase{repo: repo}
}

func (uc *ProductReviewUseCase) GetProductReviewsByProductId(productId uuid.UUID) ([]*entities.ProductReview, error) {
	return uc.repo.GetProductReviewsByProductId(productId)
}

func (uc *ProductReviewUseCase) GetProductReviewsByUserId(userId uuid.UUID) ([]*entities.ProductReview, error) {
	return uc.repo.GetProductReviewsByUserId(userId)
}

func (uc *ProductReviewUseCase) CreateProductReview(review *entities.ProductReview) (*entities.ProductReview, error) {
	return uc.repo.Create(review)
}
