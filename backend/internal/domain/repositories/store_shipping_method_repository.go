package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StoreShippingMethodRepository interface {
	FindById(id uuid.UUID) (*entities.StoreShippingMethod, error)
	Create(storeShippingMethod *entities.StoreShippingMethod) (*entities.StoreShippingMethod, error)
	GetStoreShippingMethodsByStoreId(storeId uuid.UUID) ([]*entities.StoreShippingMethod, error)
}
