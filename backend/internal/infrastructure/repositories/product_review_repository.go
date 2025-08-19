package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ProductReviewRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ProductReviewRepositoryImpl) FindById(id uuid.UUID) (*entities.ProductReview, error) {
	panic("implement me")
}

func (r *ProductReviewRepositoryImpl) Create(productReview *entities.ProductReview) (*entities.ProductReview, error) {
	panic("implement me")
}

func (r *ProductReviewRepositoryImpl) GetProductReviewsByProductId(productId uuid.UUID) ([]*entities.ProductReview, error) {
	panic("implement me")
}

func (r *ProductReviewRepositoryImpl) GetProductReviewsByUserId(userId uuid.UUID) ([]*entities.ProductReview, error) {
	panic("implement me")
}
