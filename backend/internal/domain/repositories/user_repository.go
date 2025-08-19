package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID) error
}
