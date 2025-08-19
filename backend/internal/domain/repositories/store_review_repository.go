package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StoreReviewRepository interface {
	FindById(id uuid.UUID) (*entities.StoreReview, error)
	Create(storeReview *entities.StoreReview) (*entities.StoreReview, error)
	GetStoreReviewsByStoreId(storeId uuid.UUID) ([]*entities.StoreReview, error)
	GetStoreReviewsByUserId(userId uuid.UUID) ([]*entities.StoreReview, error)
}
