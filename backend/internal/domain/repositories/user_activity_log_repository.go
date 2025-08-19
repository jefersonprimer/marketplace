package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type UserActivityLogRepository interface {
	FindById(id uuid.UUID) (*entities.UserActivityLog, error)
	Create(userActivityLog *entities.UserActivityLog) (*entities.UserActivityLog, error)
	GetUserActivityLogsByUserId(userId uuid.UUID) ([]*entities.UserActivityLog, error)
}
