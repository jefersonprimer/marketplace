package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ProductReviewRepository interface {
	FindById(id uuid.UUID) (*entities.ProductReview, error)
	Create(productReview *entities.ProductReview) (*entities.ProductReview, error)
	GetProductReviewsByProductId(productId uuid.UUID) ([]*entities.ProductReview, error)
	GetProductReviewsByUserId(userId uuid.UUID) ([]*entities.ProductReview, error)
}
