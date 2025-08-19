package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type PromotionRepositoryImpl struct {
	DB *gorm.DB
}

func (r *PromotionRepositoryImpl) FindById(id uuid.UUID) (*entities.Promotion, error) {
	panic("implement me")
}

func (r *PromotionRepositoryImpl) Create(promotion *entities.Promotion) (*entities.Promotion, error) {
	panic("implement me")
}

func (r *PromotionRepositoryImpl) GetPromotionsByStoreId(storeId uuid.UUID) ([]*entities.Promotion, error) {
	panic("implement me")
}
