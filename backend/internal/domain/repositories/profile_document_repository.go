package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type ProfileDocumentRepository interface {
	FindById(id uuid.UUID) (*entities.ProfileDocument, error)
	Create(profileDocument *entities.ProfileDocument) (*entities.ProfileDocument, error)
	GetProfileDocumentsByProfileId(profileId uuid.UUID) ([]*entities.ProfileDocument, error)
}
