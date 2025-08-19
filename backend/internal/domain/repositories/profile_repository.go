package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type ProfileRepository interface {
	Create(profile *entities.Profile) error
	GetByID(id uuid.UUID) (*entities.Profile, error)
	Update(profile *entities.Profile) error
	Delete(id uuid.UUID) error
}
