package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type PromotionRepository interface {
	FindById(id uuid.UUID) (*entities.Promotion, error)
	Create(promotion *entities.Promotion) (*entities.Promotion, error)
	GetPromotionsByStoreId(storeId uuid.UUID) ([]*entities.Promotion, error)
}
