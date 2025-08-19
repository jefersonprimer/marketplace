package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type UserActivityLogRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserActivityLogRepositoryImpl) FindById(id uuid.UUID) (*entities.UserActivityLog, error) {
	panic("implement me")
}

func (r *UserActivityLogRepositoryImpl) Create(userActivityLog *entities.UserActivityLog) (*entities.UserActivityLog, error) {
	panic("implement me")
}

func (r *UserActivityLogRepositoryImpl) GetUserActivityLogsByUserId(userId uuid.UUID) ([]*entities.UserActivityLog, error) {
	panic("implement me")
}
