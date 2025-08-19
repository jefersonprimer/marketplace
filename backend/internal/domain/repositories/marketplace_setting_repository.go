package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type MarketplaceSettingRepository interface {
	Create(setting *entities.MarketplaceSetting) error
	GetByID(id uuid.UUID) (*entities.MarketplaceSetting, error)
	GetByKey(key string) (*entities.MarketplaceSetting, error)
	Update(setting *entities.MarketplaceSetting) error
	Delete(id uuid.UUID) error
}
