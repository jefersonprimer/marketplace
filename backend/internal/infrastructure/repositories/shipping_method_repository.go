package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ShippingMethodRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ShippingMethodRepositoryImpl) FindById(id uuid.UUID) (*entities.ShippingMethod, error) {
	panic("implement me")
}

func (r *ShippingMethodRepositoryImpl) Create(shippingMethod *entities.ShippingMethod) (*entities.ShippingMethod, error) {
	panic("implement me")
}

func (r *ShippingMethodRepositoryImpl) GetAll() ([]*entities.ShippingMethod, error) {
	panic("implement me")
}
