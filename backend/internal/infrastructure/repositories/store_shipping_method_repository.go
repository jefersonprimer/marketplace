package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StoreShippingMethodRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StoreShippingMethodRepositoryImpl) FindById(id uuid.UUID) (*entities.StoreShippingMethod, error) {
	panic("implement me")
}

func (r *StoreShippingMethodRepositoryImpl) Create(storeShippingMethod *entities.StoreShippingMethod) (*entities.StoreShippingMethod, error) {
	panic("implement me")
}

func (r *StoreShippingMethodRepositoryImpl) GetStoreShippingMethodsByStoreId(storeId uuid.UUID) ([]*entities.StoreShippingMethod, error) {
	panic("implement me")
}
