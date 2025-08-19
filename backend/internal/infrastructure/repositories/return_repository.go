package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ReturnRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ReturnRepositoryImpl) FindById(id uuid.UUID) (*entities.Return, error) {
	panic("implement me")
}

func (r *ReturnRepositoryImpl) Create(returnEntity *entities.Return) (*entities.Return, error) {
	panic("implement me")
}

func (r *ReturnRepositoryImpl) GetReturnsByOrderId(orderId uuid.UUID) ([]*entities.Return, error) {
	panic("implement me")
}
