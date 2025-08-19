package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	Create(notification *entities.Notification) error
	GetByID(id uuid.UUID) (*entities.Notification, error)
	GetByUserID(userID uuid.UUID) ([]*entities.Notification, error)
	Update(notification *entities.Notification) error
	Delete(id uuid.UUID) error
}
