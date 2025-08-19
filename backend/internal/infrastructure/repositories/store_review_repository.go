package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StoreReviewRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StoreReviewRepositoryImpl) FindById(id uuid.UUID) (*entities.StoreReview, error) {
	panic("implement me")
}

func (r *StoreReviewRepositoryImpl) Create(storeReview *entities.StoreReview) (*entities.StoreReview, error) {
	panic("implement me")
}

func (r *StoreReviewRepositoryImpl) GetStoreReviewsByStoreId(storeId uuid.UUID) ([]*entities.StoreReview, error) {
	panic("implement me")
}

func (r *StoreReviewRepositoryImpl) GetStoreReviewsByUserId(userId uuid.UUID) ([]*entities.StoreReview, error) {
	panic("implement me")
}
